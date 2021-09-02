//
// Copyright (c) 2021 Seagate Technology LLC and/or its Affiliates
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// For any questions about this software or licensing,
// please email opensource@seagate.com or cortx-questions@seagate.com.

package exosx

import (
	"fmt"
	"strconv"
	"testing"

	. "github.com/onsi/gomega"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"k8s.io/klog"
)

//
// An API Test Suite
//
// Goal: Execute every possible Storage Array API call used by the CSI Driver so
//       those calls are validated against all supported storage array API versions.
//
// API Calls Tested:
//     /login
//     /show/versions/detail
//     /show/controllers
//     /show/pools
//     /create/volume
//     /show/volumes
//     /show/initiators/
//     /show/maps
//     /show/maps/initiators/
//     /map/volume
//     /expand/volume
//     /copy/volume
//     /show/snapshots
//     /create/snapshots
//     /delete/snapshot
//     /unmap/volume
//     /delete/volumes
//     /set/initiator - called if an initiator nickname is needed
//

var size = "1GiB"
var volname1 = "apitest_1"
var volname2 = "apitest_2"
var expandSize = "1GiB"
var snap1 = "snap1"
var snap2 = "snap2"
var loginFail = false

func ShowVolume(t *testing.T, volumeName string) {
	g := NewWithT(t)

	response, status, err := client.ShowVolumes(volumeName)
	g.Expect(err).To(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(0))

	volume := response.ObjectsMap["volume"]
	bint, _ := strconv.ParseInt(volume.PropertiesMap["blocks"].Data, 10, 64)
	bsint, _ := strconv.ParseInt(volume.PropertiesMap["blocksize"].Data, 10, 64)

	p := message.NewPrinter(language.English)
	p.Printf("\n")
	p.Printf("    volume-name       = %v\n", volume.PropertiesMap["volume-name"].Data)
	p.Printf("    storage-pool-name = %v\n", volume.PropertiesMap["storage-pool-name"].Data)
	p.Printf("    blocksize         = %v\n", bsint)
	p.Printf("    blocks            = %d\n", bint)
	p.Printf("    current size      = %d\n", bsint*bint)
	p.Printf("    tier-affinity     = %v\n", volume.PropertiesMap["tier-affinity"].Data)
	p.Printf("\n")
}

func ShowVolumes(t *testing.T) {
	g := NewWithT(t)

	response, status, err := client.ShowVolumes()
	g.Expect(err).To(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(0))

	if err == nil {
		fmt.Printf("\n")
		fmt.Printf("Volumes:\n")
		for _, object := range response.Objects {
			if object.Name == "volume" {
				blocks, _ := strconv.ParseInt(object.PropertiesMap["blocks"].Data, 10, 64)
				blocksize, _ := strconv.ParseInt(object.PropertiesMap["blocksize"].Data, 10, 64)

				fmt.Printf("%8v, %32v, %10v, %8v, %10v, %v, %8v, %v\n",
					object.PropertiesMap["storage-pool-name"].Data,
					object.PropertiesMap["volume-name"].Data,
					object.PropertiesMap["total-size"].Data,
					blocks,
					blocksize,
					object.PropertiesMap["storage-type"].Data,
					object.PropertiesMap["volume-type"].Data,
					object.PropertiesMap["health"].Data,
				)
			}
		}
	}

	fmt.Printf("\n")
}

func ShowSnapshots(t *testing.T, names string) {
	g := NewWithT(t)

	var err error
	var status *ResponseStatus
	var response *Response
	response, status, err = client.ShowSnapshots(names)
	g.Expect(err).To(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(0))

	if err == nil {
		for _, obj := range response.Objects {
			fmt.Printf("++ obj.Name = %v\n", obj.Name)
		}
	}

	fmt.Printf("\n")
}

func ConditionalSkip(t *testing.T) {
	if loginFail {
		t.Skip()
	}
}
func TestAPILogin(t *testing.T) {
	g := NewWithT(t)
	err := client.Login()
	if err != nil {
		loginFail = true
	}
	g.Expect(err).To(BeNil())
}

func TestAPIVersions(t *testing.T) {
	ConditionalSkip(t)
	g := NewWithT(t)

	response, status, err := client.FormattedRequest("/show/versions/detail")
	g.Expect(err).To(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(0))

	fmt.Printf("\n")
	fmt.Printf("                  Controller A Versions        Controller B Versions\n")
	fmt.Printf("                  ---------------------        ---------------------\n")

	var mccodea, mcbasea, mccodeb, mcbaseb string

	if err == nil {
		for _, obj := range response.Objects {
			if obj.Name == "controller-a-versions" {
				mccodea = obj.PropertiesMap["mc-fw"].Data
				mcbasea = obj.PropertiesMap["mc-base-fw"].Data
			}
			if obj.Name == "controller-b-versions" {
				mccodeb = obj.PropertiesMap["mc-fw"].Data
				mcbaseb = obj.PropertiesMap["mc-base-fw"].Data
			}
		}
	}

	fmt.Printf("MC Code Version:  %-28v %v\n", mccodea, mccodeb)
	fmt.Printf("MC Base Version:  %-28v %v\n", mcbasea, mcbaseb)
	fmt.Printf("\n")
}

func TestAPIControllers(t *testing.T) {
	ConditionalSkip(t)
	g := NewWithT(t)

	respone, status, err := client.FormattedRequest("/show/controllers")
	g.Expect(err).To(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(0))

	if err == nil {
		for _, obj1 := range respone.Objects {
			if obj1.Name == "controller" || obj1.Name == "controllers" {
				fmt.Printf("\n")
				fmt.Printf("Controller     = %v\n", obj1.PropertiesMap["controller-id"].Data)
				fmt.Printf("Disks          = %v\n", obj1.PropertiesMap["disks"].Data)
				fmt.Printf("IP Address     = %v\n", obj1.PropertiesMap["ip-address"].Data)
				fmt.Printf("Model          = %v\n", obj1.PropertiesMap["model"].Data)
				fmt.Printf("Platform       = %v\n", obj1.PropertiesMap["platform-type"].Data)
				fmt.Printf("Serial Number  = %v\n", obj1.PropertiesMap["serial-number"].Data)
				fmt.Printf("Status         = %v\n", obj1.PropertiesMap["status"].Data)
				fmt.Printf("Storage Pools  = %v\n", obj1.PropertiesMap["number-of-storage-pools"].Data)
				fmt.Printf("Disk Groups    = %v\n", obj1.PropertiesMap["virtual-disks"].Data)

				for _, obj2 := range obj1.Objects {
					if obj2.Name == "ports" {
						fmt.Printf("-- Port        = %v, %v, %v",
							obj2.PropertiesMap["port"].Data,
							obj2.PropertiesMap["port-type"].Data,
							obj2.PropertiesMap["target-id"].Data,
						)
						if obj2.PropertiesMap["port-type"].Data == "iSCSI" {
							for _, obj3 := range obj2.Objects {
								// fmt.Printf("obj3=%+v\n", obj3)
								// fmt.Printf("obj3.Name=%v\n", obj3.Name)
								if obj3.Name == "port-details" {
									fmt.Printf(", %15v, %12v, %v",
										obj3.PropertiesMap["ip-address"].Data,
										obj3.PropertiesMap["sfp-present"].Data,
										obj3.PropertiesMap["sfp-ethernet-compliance"].Data,
									)
								}
							}
						}
						fmt.Printf("\n")
					}
				}
			}
		}
	}

	fmt.Printf("\n")
}

func TestAPIInitSystem(t *testing.T) {
	ConditionalSkip(t)
	g := NewWithT(t)
	err := client.InitSystem(client.PoolName)
	g.Expect(err).To(BeNil())
}

func TestAPICreateVolume(t *testing.T) {
	ConditionalSkip(t)
	g := NewWithT(t)

	var err error
	var status *ResponseStatus
	_, status, err = client.CreateVolume(volname1, size, client.PoolName, client.PoolData.Type)
	g.Expect(err).To(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(0))
	ShowVolume(t, volname1)
	ShowVolumes(t)
}

func TestAPIShowInitiators(t *testing.T) {
	ConditionalSkip(t)
	g := NewWithT(t)

	respone, status, err := client.FormattedRequest("/show/initiators/")
	g.Expect(err).To(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(0))

	fmt.Printf("\n")
	fmt.Printf("Nickname        Discovered Mapped Profile  Host Type  ID\n")

	if err == nil {
		for _, obj := range respone.Objects {
			if obj.Name == "initiator" {
				fmt.Printf("%-16s", obj.PropertiesMap["nickname"].Data)
				fmt.Printf("%-11s", obj.PropertiesMap["discovered"].Data)
				fmt.Printf("%-7s", obj.PropertiesMap["mapped"].Data)
				fmt.Printf("%-9s", obj.PropertiesMap["profile"].Data)
				fmt.Printf("%-11s", obj.PropertiesMap["host-bus-type"].Data)
				fmt.Printf("%s", obj.PropertiesMap["id"].Data)
				fmt.Printf("\n")
			}
		}
	}

	fmt.Printf("\n")
}

func TestAPIGetMaps(t *testing.T) {
	ConditionalSkip(t)
	g := NewWithT(t)

	respone, status, err := client.FormattedRequest("/show/maps/%s", volname1)
	g.Expect(err).To(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(0))

	hostNames := []string{}

	if err == nil {
		for _, rootObj := range respone.Objects {
			if rootObj.Name != "volume-view" {
				continue
			}

			for _, object := range rootObj.Objects {
				hostName := object.PropertiesMap["identifier"].Data
				if object.Name == "host-view" && hostName != "all other initiators" {
					hostNames = append(hostNames, hostName)
				}
			}
		}
	}

	fmt.Printf("volume %q host names:\n", volname1)
	for i, h := range hostNames {
		fmt.Printf("    [%d] %s\n", i, h)
	}
}

func TestAPIMapVolume(t *testing.T) {
	ConditionalSkip(t)
	g := NewWithT(t)
	lun, err := client.PublishVolume(volname1, client.Initiator)
	g.Expect(err).To(BeNil())
	g.Expect(lun).ToNot(Equal(0))
}

func TestAPIExpandVolume(t *testing.T) {
	ConditionalSkip(t)
	g := NewWithT(t)

	klog.Infof("expand volume (%s) from original size (%s) to new size (%s)", volname1, size, expandSize)
	_, status, err := client.ExpandVolume(volname1, expandSize)
	if err != nil {
		if status != nil && status.ReturnCode != 0 {
			fmt.Printf("expand volume failed, status.ReturnCode=%v", status.ReturnCode)
		}
	}
	g.Expect(err).To(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(0))
	klog.Infof("successfully expanded volume (%s)", volname1)
	ShowVolume(t, volname1)
}

func TestAPICreateSnapshots(t *testing.T) {
	ConditionalSkip(t)

	if client.PoolData.Type == "Linear" {
		fmt.Printf("Linear snapshots are not supported\n")
		return
	}

	g := NewWithT(t)

	klog.Infof("snapshot volume (%s) using name (%s)", volname1, snap1)
	_, status, err := client.CreateSnapshot(volname1, snap1)
	if err != nil {
		if status != nil && status.ReturnCode != 0 {
			fmt.Printf("snapshot volume failed, status.ReturnCode=%v", status.ReturnCode)
		}
	}
	g.Expect(err).To(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(0))
	klog.Infof("successfully snapped volume (%s)", snap1)
	ShowVolumes(t)

	klog.Infof("snapshot volume (%s) using name (%s)", volname1, snap2)
	_, status, err = client.CreateSnapshot(volname1, snap2)
	if err != nil {
		if status != nil && status.ReturnCode != 0 {
			fmt.Printf("snapshot volume failed, status.ReturnCode=%v", status.ReturnCode)
		}
	}
	g.Expect(err).To(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(0))
	klog.Infof("successfully snapped volume (%s)", snap2)
	ShowVolumes(t)
}

func TestAPIDeleteSnapshots(t *testing.T) {
	ConditionalSkip(t)

	if client.PoolData.Type == "Linear" {
		fmt.Printf("Linear snapshots are not supported\n")
		return
	}

	g := NewWithT(t)

	klog.Infof("delete snapshot (%s)", snap1)
	_, status, err := client.DeleteSnapshot(snap1)
	if err != nil {
		if status != nil && status.ReturnCode != 0 {
			fmt.Printf("delete snapshot failed, status.ReturnCode=%v", status.ReturnCode)
		}
	}
	g.Expect(err).To(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(0))
	klog.Infof("successfully deleted snapshot (%s)", snap1)

	klog.Infof("delete snapshot (%s)", snap2)
	_, status, err = client.DeleteSnapshot(snap2)
	if err != nil {
		if status != nil && status.ReturnCode != 0 {
			fmt.Printf("delete snapshot failed, status.ReturnCode=%v", status.ReturnCode)
		}
	}
	g.Expect(err).To(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(0))
	klog.Infof("successfully deleted snapshot (%s)", snap2)
	ShowVolumes(t)
}

func TestAPIUnmapVolume(t *testing.T) {
	ConditionalSkip(t)
	g := NewWithT(t)

	klog.Infof("unmapping volume %s from initiator %s", volname1, client.Initiator)
	_, status, err := client.UnmapVolume(volname1, client.Initiator)
	if err != nil {
		if status != nil && status.ReturnCode == unmapFailedErrorCode {
			fmt.Printf("unmap failed, assuming volume is already unmapped")
		}
	}
	g.Expect(err).To(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(0))
	klog.Infof("successfully unmapped volume %s from initiator %s", volname1, client.Initiator)
}

func TestAPICopyVolume(t *testing.T) {
	ConditionalSkip(t)

	if client.PoolData.Type == "Linear" {
		fmt.Printf("Linear snapshots are not supported\n")
		return
	}

	g := NewWithT(t)

	klog.Infof("copy volume (%s) to (%s) using pool (%s)", volname1, volname1, client.PoolName)

	_, status, err := client.CopyVolume(volname1, volname2, client.PoolName)
	if err != nil {
		if status != nil && status.ReturnCode != 0 {
			fmt.Printf("copy volume failed, status.ReturnCode=%v", status.ReturnCode)
		}
	}
	g.Expect(err).To(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(0))
	klog.Infof("successfully copied volume to (%s)", volname2)
	ShowVolumes(t)
}

func TestAPIDeleteVolumes(t *testing.T) {
	ConditionalSkip(t)
	g := NewWithT(t)

	_, status, err := client.DeleteVolume(volname1)
	g.Expect(err).To(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(0))

	if client.PoolData.Type != "Linear" {
		_, status, err := client.DeleteVolume(volname2)
		g.Expect(err).To(BeNil())
		g.Expect(status.ResponseTypeNumeric).To(Equal(0))
	}

	ShowVolumes(t)
}
