## [1.0.12](https://github.com/Seagate/seagate-exos-x-api-go/compare/v1.0.11...v1.0.12) (2022-12-22)

### Bug Fixes

- avoid recursive call to Login() (#27) ([3030dd2](https://github.com/Seagate/seagate-exos-x-api-go/commit/3030dd2fdea1423a997120fb8653a92b89f0f32b)), closes [#27](https://github.com/Seagate/seagate-exos-x-api-go/issues/27)

### Other

- Merge pull request #31 from Seagate/fix_recursive_login ([f617f52](https://github.com/Seagate/seagate-exos-x-api-go/commit/f617f5221b1d436ffa37e9f365d84c83fde04886)), closes [#31](https://github.com/Seagate/seagate-exos-x-api-go/issues/31) [#27](https://github.com/Seagate/seagate-exos-x-api-go/issues/27)

## [1.0.11](https://github.com/Seagate/seagate-exos-x-api-go/compare/v1.0.10...v1.0.11) (2022-12-12)

### Chores

- Update logging to klog2 for compatibility with csi driver ([815552d](https://github.com/Seagate/seagate-exos-x-api-go/commit/815552d29c5d1b445b9295165fb7c6d9b7c210ec))

## [1.0.10](https://github.com/Seagate/seagate-exos-x-api-go/compare/v1.0.9...v1.0.10) (2022-11-23)

### Bug Fixes

- **ShowHostMaps:** Return host maps when initiator has a nickname or is in a group ([9b72e0f](https://github.com/Seagate/seagate-exos-x-api-go/commit/9b72e0ff3de8c30b15c39cd6d46b294dfabd2b3e))

### Chores

- add workflow_dispatch to enable manual workflow run ([35cf7d9](https://github.com/Seagate/seagate-exos-x-api-go/commit/35cf7d9492fbf3bc60d9122773a4f9637066619b))
- remove release.yml changes ([d7686d1](https://github.com/Seagate/seagate-exos-x-api-go/commit/d7686d161cee7ab5ed2d998f02dfefd4b2f79f6f))
- update workflow_dispatch ([b3f01be](https://github.com/Seagate/seagate-exos-x-api-go/commit/b3f01be4189dd22dccdedbd9fde0070d235802fd))

### Other

- Merge pull request #26 from Seagate/chore/manualrelease2 ([c266540](https://github.com/Seagate/seagate-exos-x-api-go/commit/c266540f044528c470a60413ef263f3e21a67f7a)), closes [#26](https://github.com/Seagate/seagate-exos-x-api-go/issues/26)
- Merge pull request #25 from Seagate/chore/manualrelease ([b6d4a6e](https://github.com/Seagate/seagate-exos-x-api-go/commit/b6d4a6e5a31adb828470a8fb0b0dbe65a3f59a67)), closes [#25](https://github.com/Seagate/seagate-exos-x-api-go/issues/25)

## [1.0.8](https://github.com/Seagate/seagate-exos-x-api-go/compare/v1.0.7...v1.0.8) (2022-05-31)

### Bug Fixes

- ChooseLUN uses next highest lun before wrapping ([80883af](https://github.com/Seagate/seagate-exos-x-api-go/commit/80883af1a22a1043f83fe461343eaed57564711f))
- Correctly handle a volumes slice of zero length ([242b3fa](https://github.com/Seagate/seagate-exos-x-api-go/commit/242b3fa540319d8ec1d21200d0a390a5fed275ca))

### Other

- Merge pull request #19 from Seagate/fix/chooselun ([3d1a38b](https://github.com/Seagate/seagate-exos-x-api-go/commit/3d1a38b18ac6440960897e2ae08f06a62a52d86f)), closes [#19](https://github.com/Seagate/seagate-exos-x-api-go/issues/19)

## [1.0.7](https://github.com/Seagate/seagate-exos-x-api-go/compare/v1.0.6...v1.0.7) (2022-05-26)

### Bug Fixes

- CheckVolumeExists() optimization and new GetVolumeWwn() ([c73083a](https://github.com/Seagate/seagate-exos-x-api-go/commit/c73083a51c34135523b83c97b10c79da810fa950))

### Code Style

- GetVolumeWwn() improvement ([5e2424f](https://github.com/Seagate/seagate-exos-x-api-go/commit/5e2424f40cf0fb6561a620c602e490df436f3b94))

### Other

- Merge pull request #18 from Seagate/fix/wwn ([4bbb191](https://github.com/Seagate/seagate-exos-x-api-go/commit/4bbb191320d7abe0f53fc3c9e73eab711788c774)), closes [#18](https://github.com/Seagate/seagate-exos-x-api-go/issues/18)

## [1.0.6](https://github.com/Seagate/seagate-exos-x-api-go/compare/v1.0.5...v1.0.6) (2022-05-26)

### Chores

- **deps:** bump minimist from 1.2.5 to 1.2.6 ([a58ea6f](https://github.com/Seagate/seagate-exos-x-api-go/commit/a58ea6fc323c0eb20349658486b63389a3042460))
- **deps:** bump node-fetch from 2.6.5 to 2.6.7 ([a7bfda6](https://github.com/Seagate/seagate-exos-x-api-go/commit/a7bfda62b438d89ac2c88e378722faa71585b019))

### Other

- Merge pull request #17 from Seagate/dependabot/npm_and_yarn/node-fetch-2.6.7 ([b38d199](https://github.com/Seagate/seagate-exos-x-api-go/commit/b38d199d806d0e2cb257145ec6b8731de816cb08)), closes [#17](https://github.com/Seagate/seagate-exos-x-api-go/issues/17)
- Merge pull request #16 from Seagate/dependabot/npm_and_yarn/minimist-1.2.6 ([a8414ff](https://github.com/Seagate/seagate-exos-x-api-go/commit/a8414ff7b5e04d1065c7cc4d51e992472624e1c6)), closes [#16](https://github.com/Seagate/seagate-exos-x-api-go/issues/16)

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
