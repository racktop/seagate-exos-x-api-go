## [1.0.5](https://github.com/Seagate/seagate-exos-x-api-go/compare/v1.0.4...v1.0.5) (2021-11-24)

### Bug Fixes

- **GetTargetId:** Retrieve id by type, such as iscsi, fc, etc ([5c8772b](https://github.com/Seagate/seagate-exos-x-api-go/commit/5c8772b9e7eb57b91976c329d037225c45e444ec))

### Other

- Merge pull request #15 from Seagate/fix/get-targetid ([a4f9042](https://github.com/Seagate/seagate-exos-x-api-go/commit/a4f9042a8fb9dea9ef08f903f9af5653fb7786a9)), closes [#15](https://github.com/Seagate/seagate-exos-x-api-go/issues/15)

## [1.0.4](https://github.com/Seagate/seagate-exos-x-api-go/compare/v1.0.3...v1.0.4) (2021-11-22)

### Bug Fixes

- ShowSnapshots() now takes two parameters, snapshot id and source volume id ([539fafe](https://github.com/Seagate/seagate-exos-x-api-go/commit/539fafe16eb06ca33391170986b7d9ec0c20dde0))

### Other

- Merge pull request #14 from Seagate/test/csi-sanity-other ([150a719](https://github.com/Seagate/seagate-exos-x-api-go/commit/150a719963981c13ac19ac2cea37028e558356e3)), closes [#14](https://github.com/Seagate/seagate-exos-x-api-go/issues/14)
- Merge pull request #13 from Seagate/test/csi-sanity-volumes ([42401a3](https://github.com/Seagate/seagate-exos-x-api-go/commit/42401a3c2c03e7b45642ca73ca0e0dce7daf6025)), closes [#13](https://github.com/Seagate/seagate-exos-x-api-go/issues/13)

### Tests

- correct ControllerPublishVolume csi-sanity issues ([7e4e2e5](https://github.com/Seagate/seagate-exos-x-api-go/commit/7e4e2e5dd1f5f23ab2b3b3c7e78b295d0aab2016))

## [1.0.3](https://github.com/Seagate/seagate-exos-x-api-go/compare/v1.0.2...v1.0.3) (2021-10-04)

### Bug Fixes

- correct go.mod path ([b470e32](https://github.com/Seagate/seagate-exos-x-api-go/commit/b470e328841368ab49213bf9159f043b5c14cb09))

### Other

- Merge branch 'main' of github.com:Seagate/seagate-exos-x-api-go into main ([b539b0d](https://github.com/Seagate/seagate-exos-x-api-go/commit/b539b0d76e5358f8b689559b0decd8ecfc3615c8))

## [1.0.2](https://github.com/Seagate/seagate-exos-x-api-go/compare/v1.0.1...v1.0.2) (2021-10-03)

### Chores

- trim changelog ([6fdf232](https://github.com/Seagate/seagate-exos-x-api-go/commit/6fdf232ba2ad7bbaa1012be5e9926c0ad7491aa0))

## [1.0.1](https://github.com/Seagate/seagate-exos-x-api-go/compare/v1.0.0...v1.0.1) (2021-10-03)

### Chores

- correct license in package.json ([5d6b648](https://github.com/Seagate/seagate-exos-x-api-go/commit/5d6b648cd8f63675fbc6dcd63c1a6727ca8b180b))

# 1.0.0 (2021-10-03)

### Bug Fixes

- **fix:** Discover iSCSI IQN and Portals from storage appliance, removes StorageClass requirement ([3e8bcc4](https://github.com/Seagate/seagate-exos-x-api-go/commit/3e8bcc4755fc411100511f596237680193d1fa34))
- **fix:** handle session timeout error code (2) ([dad3e24](https://github.com/Seagate/seagate-exos-x-api-go/commit/dad3e240b25060ccda74b9b36f01fd759d0346ed))
