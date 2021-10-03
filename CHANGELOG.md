# 1.0.0 (2021-10-03)

### Bug Fixes

- **api:** updates to handle session timeout, remove LastAccess code ([fcbc491](https://github.com/Seagate/seagate-exos-x-api-go/commit/fcbc491075e2ec86d1390fb85b57e6f3cdcdc8b1))
- **ci:** fix entrypoint in image ([9e7cc90](https://github.com/Seagate/seagate-exos-x-api-go/commit/9e7cc9084c6c497d78369e66fb871ed3b647b3ec))
- **ci:** use an image with docker compose ([2993fcf](https://github.com/Seagate/seagate-exos-x-api-go/commit/2993fcf3e69ba1a6e063e11d9c2360586a43aa89))
- **client:** check if address is configured before executing a request ([ba83121](https://github.com/Seagate/seagate-exos-x-api-go/commit/ba83121ef8b26dbd1c4ea871e191a221b4a51c9a))
- **client:** check if address is configured before executing a request ([de16a8b](https://github.com/Seagate/seagate-exos-x-api-go/commit/de16a8b6f6b68a2eeb387226fb455979d629903a))
- **client:** delete dangerous print ([a1fcc06](https://github.com/Seagate/seagate-exos-x-api-go/commit/a1fcc067a3119c8b3cf79276cf4438ce816bbe1e))
- **client:** don't abort on return code !=, but on response type numeric instead ([671bd4a](https://github.com/Seagate/seagate-exos-x-api-go/commit/671bd4a007ef1e2fafed3f1813ec82d1a7f76cc8))
- **client:** refresh login when session key expires ([248496f](https://github.com/Seagate/seagate-exos-x-api-go/commit/248496f009c5c0c4940f80c21767f5483c3ab100))
- **client:** remove useless code and comments ([80378db](https://github.com/Seagate/seagate-exos-x-api-go/commit/80378db38edae68d8d8bb4f2feef9ab23359c8da))
- **client:** return an error status instead of nil ([47be182](https://github.com/Seagate/seagate-exos-x-api-go/commit/47be182b36951f4289106b4000fd6ffcf790472f)), closes [#5](https://github.com/Seagate/seagate-exos-x-api-go/issues/5)
- **client:** reuse the same connection for all requests (keep-alive) ([2b8d60c](https://github.com/Seagate/seagate-exos-x-api-go/commit/2b8d60c4cc4ee4c29b775ae342aefad0923a3e20))
- **client:** set timeout to 15s ([3015fdb](https://github.com/Seagate/seagate-exos-x-api-go/commit/3015fdb7698f2c022fc009b1c0bece426fcdd9c6))
- Discover iSCSI IQN and Portals from storage appliance, removes StorageClass requirement ([3e8bcc4](https://github.com/Seagate/seagate-exos-x-api-go/commit/3e8bcc4755fc411100511f596237680193d1fa34))
- **endpoints:** add accessMode to map volume call ([c87f948](https://github.com/Seagate/seagate-exos-x-api-go/commit/c87f94894741773871b540e9aee76ce49ec2363f))
- **endpoints:** allow for empty host field in unmap volume call ([b90384a](https://github.com/Seagate/seagate-exos-x-api-go/commit/b90384a7c31da5fc4431de792bc1388372231d26))
- **endpoints:** don't crash on non-existent host for show host maps call ([96bca84](https://github.com/Seagate/seagate-exos-x-api-go/commit/96bca84aedc43e786ce861bf65266f9fd42a165f))
- **endpoints:** fix ShowHostMaps("") behavior ([655c8d9](https://github.com/Seagate/seagate-exos-x-api-go/commit/655c8d9fca8d4ff98eed909d4d4d78ccba022647))
- handle session timeout error code (2) ([dad3e24](https://github.com/Seagate/seagate-exos-x-api-go/commit/dad3e240b25060ccda74b9b36f01fd759d0346ed))
- **tests:** change returnCode check to responseTypeNumeric ([0010a2f](https://github.com/Seagate/seagate-exos-x-api-go/commit/0010a2fde44714641c9db9f97b42389997104773))
- **tests:** fix error handling tests crashing ([26a0ec8](https://github.com/Seagate/seagate-exos-x-api-go/commit/26a0ec8688732834258fa654d10786658cb2b706))
- **tests:** update tests go image to enix' one because of a breaking update from the previous image ([182fefe](https://github.com/Seagate/seagate-exos-x-api-go/commit/182fefedec22fd58002ad1e2761eceeeb022d9a0))

### Chores

- [skip release] add checkout step to workflow ([fed2c28](https://github.com/Seagate/seagate-exos-x-api-go/commit/fed2c285217da0988bfae517e1f4c8e83a30e7a9))
- [skip release] added main as release branch ([3f3217f](https://github.com/Seagate/seagate-exos-x-api-go/commit/3f3217f8c6a0f42863276ec9456e15daabb58fbd))
- [skip release] adding npm package files for semantic-release ([94e06c6](https://github.com/Seagate/seagate-exos-x-api-go/commit/94e06c6434202bc37a3ced83dd165c2a8fd93b43))
- [skip release] Corrected release workflow ([43ffd22](https://github.com/Seagate/seagate-exos-x-api-go/commit/43ffd22f7007fea4334c7e5ef697c99ba68b33ac))
- [skip release] Corrected release workflow ([5b639bf](https://github.com/Seagate/seagate-exos-x-api-go/commit/5b639bf7f253508c1b417efba9e05a15eedf9cdb))
- [skip release] github release workflow ([5f114e7](https://github.com/Seagate/seagate-exos-x-api-go/commit/5f114e7fdebbb14998b74d4f48c89d1808fa8b6b))
- [skip release] github workflow update ([b6e76ba](https://github.com/Seagate/seagate-exos-x-api-go/commit/b6e76ba6f35c85f234c9b158617f18faf855a3b8))
- [skip release] release and changelog workflow updates ([9e98925](https://github.com/Seagate/seagate-exos-x-api-go/commit/9e989254816c86436a254edf69f6e3ab8c453cbb))
- [skip release] remove changelog workflow ([233d066](https://github.com/Seagate/seagate-exos-x-api-go/commit/233d0668a68b47665dd82ae609d7f61c92f854c4))
- [skip release] Use GitHub Actions Release me! ([0d23a3b](https://github.com/Seagate/seagate-exos-x-api-go/commit/0d23a3b12067c67d34d3b3b96588d0c4a2ad39b1))
- add badges in readme ([22b1b04](https://github.com/Seagate/seagate-exos-x-api-go/commit/22b1b043ecbb198af9b0ac94ca15d39b3922fd6c))
- add pipeline status badge in readme ([94e02c7](https://github.com/Seagate/seagate-exos-x-api-go/commit/94e02c75102f5bed73d5947e52a864f87da6d792))
- change license to MIT ([6398825](https://github.com/Seagate/seagate-exos-x-api-go/commit/6398825541c4fc590d9dae4f7090d6cd4b1d8cd1))
- **ci:** fix releaserc ([a9b8ffd](https://github.com/Seagate/seagate-exos-x-api-go/commit/a9b8ffd7c31711e7e4c29c2cadff164705f917ca))
- **release:** v1.0.1 ([5a31127](https://github.com/Seagate/seagate-exos-x-api-go/commit/5a3112725cc7af5b18407d42c9dd3415de22847c))
- **tests:** add instructions on how to run tests ([d20f2f4](https://github.com/Seagate/seagate-exos-x-api-go/commit/d20f2f430dc1d4f5aec61ee56b7611973da13d18))

### Code Refactoring

- **client:** remove useless NewClient call ([4a653f4](https://github.com/Seagate/seagate-exos-x-api-go/commit/4a653f4a387190763dfeafbf1a7105df073e75e3))
- **client:** remove useless Options struct ([edc4b8b](https://github.com/Seagate/seagate-exos-x-api-go/commit/edc4b8b9b39b2101b9cd53c68b8af0970b0baeed))
- **client:** rename client.go to dothill.go for clarity reasons ([8dd8fe4](https://github.com/Seagate/seagate-exos-x-api-go/commit/8dd8fe4cfeeb6de4b0e0c6202ed31bb131ef0988))
- **client:** simplify public request() call ([5d19acb](https://github.com/Seagate/seagate-exos-x-api-go/commit/5d19acb1122a96e9bb2819e4cec03b9d48e2ddd8))
- remove an ineffectual assignment ([91544be](https://github.com/Seagate/seagate-exos-x-api-go/commit/91544bee1cf55587fa702255b776db3e7b309cc0))

### Documentation

- **license:** add LICENSE (resolves #2) ([484c36f](https://github.com/Seagate/seagate-exos-x-api-go/commit/484c36f27c3cb3d6fc21284f17fc41a534d9526a)), closes [#2](https://github.com/Seagate/seagate-exos-x-api-go/issues/2)
- **readme:** add a link the the API reference pdf and to the lib's godoc page ([403375d](https://github.com/Seagate/seagate-exos-x-api-go/commit/403375d22acfcc97e570749df1b0400f10af5fd4))

### Features

- add metrics ([5a8dcc3](https://github.com/Seagate/seagate-exos-x-api-go/commit/5a8dcc3a2dd20ef0eaee628706a8c2d87b066c85))
- **apitest2:** Added test cases for all API calls used by the driver ([94dd991](https://github.com/Seagate/seagate-exos-x-api-go/commit/94dd9917e2923b5b72e473bd577cb135f13fab65))
- **apitest2:** v0.5.1 new api test and system info retrieval ([3f1553a](https://github.com/Seagate/seagate-exos-x-api-go/commit/3f1553a21462209c22b04f85b27413c0dc92dcad))
- **apiupdates:** Storage API changes to account for newer systems, draft 1 ([f5e013f](https://github.com/Seagate/seagate-exos-x-api-go/commit/f5e013faee9395325a5ac1851ca7c6ad63755009))
- **apiupdates:** Storage API changes to account for newer systems, draft 2 ([8412fd3](https://github.com/Seagate/seagate-exos-x-api-go/commit/8412fd3eacd84bbab594fbcd39a67a3f22198f17))
- **ci:** add .gitlab-ci.yml ([a76cd56](https://github.com/Seagate/seagate-exos-x-api-go/commit/a76cd56ef4e590a28d3943fdf8e7f86a619a25ef))
- **ci:** add .releaserc.yml for semantic release ([96493b4](https://github.com/Seagate/seagate-exos-x-api-go/commit/96493b44c24c17ef085d5f19a1e5a612176f0c93))
- **ci:** add Dockerfile to run tests ([4887b15](https://github.com/Seagate/seagate-exos-x-api-go/commit/4887b15e5482a00f2fb2b2b5f44450b95706dc00))
- **ci:** configure semantic release so it releases on github too ([76e54b0](https://github.com/Seagate/seagate-exos-x-api-go/commit/76e54b0bea91682f059fddfb118bce21774616a8))
- **ci:** update .gitlab.yml to add semantic release job ([781e4a8](https://github.com/Seagate/seagate-exos-x-api-go/commit/781e4a87096803d48c5fe8f9a0376961f5045ed2))
- **client:** add logging with klog ([a3ea1bf](https://github.com/Seagate/seagate-exos-x-api-go/commit/a3ea1bfa97d60a9f703eedbed5d274266b21442d))
- **client:** add more genericity to create models/endpoints more easialy ([8e55ab8](https://github.com/Seagate/seagate-exos-x-api-go/commit/8e55ab8f9e2bdaf9f7384074eb0269d200d2d952))
- **client:** fire a real request to the given API server ([02e1797](https://github.com/Seagate/seagate-exos-x-api-go/commit/02e17971b193cde21b82214437164a76fb262549))
- **client:** generic codebase to parse responses to usable go structs ([924c34c](https://github.com/Seagate/seagate-exos-x-api-go/commit/924c34cc2823e7e77d656dc5d5ebb7f9fc44b2c8))
- **client:** ignore TLS cert errors ([1553650](https://github.com/Seagate/seagate-exos-x-api-go/commit/1553650a2dd3c355e2e4191bb35e047233cca84d))
- **client:** implement login and automatic request authentication + add test ([7fe218b](https://github.com/Seagate/seagate-exos-x-api-go/commit/7fe218b3659c02969f1b1f86d6ab6d5dbda2b5ef))
- **client:** log requests ([8287add](https://github.com/Seagate/seagate-exos-x-api-go/commit/8287addbf01aa1f5dcba4a4d2399f0f189b1d1ac))
- **client:** recursively parse nested objects ([14dc922](https://github.com/Seagate/seagate-exos-x-api-go/commit/14dc9221d13dab16824ffaf2d4c2c9e388152f77))
- **endpoint:** implement delete host route ([5366976](https://github.com/Seagate/seagate-exos-x-api-go/commit/53669760ce1db70856fe95079e0f1493b783933b))
- **endpoints:** add expand volume route ([b19576c](https://github.com/Seagate/seagate-exos-x-api-go/commit/b19576cb6d7cebb420808679488d29760192fa0c))
- **endpoints:** add show volumes call ([e4774e9](https://github.com/Seagate/seagate-exos-x-api-go/commit/e4774e9bd88ad07accb47731718dd8889dcb7426))
- **endpoints:** add snapshots endpoints ([2d3d925](https://github.com/Seagate/seagate-exos-x-api-go/commit/2d3d9256db426813d4d20e61cc8aa937c3213010))
- **endpoints:** allow listing of all lun mappings from show host mappings call ([f705a63](https://github.com/Seagate/seagate-exos-x-api-go/commit/f705a6327c08bae37d241b1507339b1d6536d18a))
- **endpoints:** implement create host call ([0ff1e0f](https://github.com/Seagate/seagate-exos-x-api-go/commit/0ff1e0f4618a31ec119cc26f1897603e8010b2c6))
- **endpoints:** implement create volume and map volume routes ([8b20879](https://github.com/Seagate/seagate-exos-x-api-go/commit/8b20879b2327a2691f194034f9ed9d40a73e8c65))
- **endpoints:** implement delete volume call ([56f6625](https://github.com/Seagate/seagate-exos-x-api-go/commit/56f6625722ec6841654cf2fc4d64484cbe48d0e2))
- **endpoints:** implement list host-view call ([ad4b543](https://github.com/Seagate/seagate-exos-x-api-go/commit/ad4b543acd2f51ac9cdc12e96d448a126a9a3ad9))
- **endpoints:** implement unmap volume call ([98eccbd](https://github.com/Seagate/seagate-exos-x-api-go/commit/98eccbdb652dc961cc533d1e0c6468cdbc6d6285))
- **endpoints:** volume copy ([e9e1556](https://github.com/Seagate/seagate-exos-x-api-go/commit/e9e155614655a47300e8bc7d8d965d1befa196e7))
- **response:** add GetProperties helper function in Object ([1202ad9](https://github.com/Seagate/seagate-exos-x-api-go/commit/1202ad9d42fe4f58adbf90b295ebedf6ae20770f))
- **tests:** add bad URL test ([1fb3e5f](https://github.com/Seagate/seagate-exos-x-api-go/commit/1fb3e5fa859c2e792d756a7408542ef89fcef2b2))
- **tests:** add error handling tests ([b03c9a7](https://github.com/Seagate/seagate-exos-x-api-go/commit/b03c9a75a458a8b78aec109e9679ea5c45a560e9))
- **tests:** add login/auth capabilities to mock server ([d1636b9](https://github.com/Seagate/seagate-exos-x-api-go/commit/d1636b9c4aab21badc5f57e2c624b3e5321c07f2))
- **tests:** create mock server ([733db5b](https://github.com/Seagate/seagate-exos-x-api-go/commit/733db5bd7f1ea754a3eb61c8ee74763e5363a3af))
- **tests:** run tests suite in docker compose with mock server ([c6c4904](https://github.com/Seagate/seagate-exos-x-api-go/commit/c6c4904f840130e41a80d765b4754d6a221bb547))
- updates to API login and main tests ([281948d](https://github.com/Seagate/seagate-exos-x-api-go/commit/281948d4c574fc430f0718f8176287a8fdd8b43f))

### Other

- Merge pull request #10 from Seagate/bug/vol-exist-error ([295e91f](https://github.com/Seagate/seagate-exos-x-api-go/commit/295e91f9c0770daf0b90299dcc848ed608e485fa)), closes [#10](https://github.com/Seagate/seagate-exos-x-api-go/issues/10)
- v0.5.7 Do not display error when checking if volume exists ([2a600d5](https://github.com/Seagate/seagate-exos-x-api-go/commit/2a600d5304b550192a024c50db0ddfe71515f5ee))
- Merge pull request #9 from Seagate/bug/snapshots ([c003d68](https://github.com/Seagate/seagate-exos-x-api-go/commit/c003d68172785775d2a623c93eb86953fcf7dd1a)), closes [#9](https://github.com/Seagate/seagate-exos-x-api-go/issues/9)
- bug/snapshots - v0.5.3 system information spelling correction ([f6ff166](https://github.com/Seagate/seagate-exos-x-api-go/commit/f6ff16671607598790e554e7685b8329f66cf8bd))
- bug/snapshots - v0.5.3 show snapshots working with all storage api versions ([7e0fc28](https://github.com/Seagate/seagate-exos-x-api-go/commit/7e0fc28cf235778cdbbd8fa183a3f6082f36c913))
- Merge pull request #8 from Seagate/feat/apitest2 ([98b20f9](https://github.com/Seagate/seagate-exos-x-api-go/commit/98b20f9e873b16c62a03c1af91765585c3f378af)), closes [#8](https://github.com/Seagate/seagate-exos-x-api-go/issues/8)
- Merge pull request #7 from Seagate/feat/apiupdates ([2e89919](https://github.com/Seagate/seagate-exos-x-api-go/commit/2e8991914930abb2ed7982d5f1eb9f7f30854015)), closes [#7](https://github.com/Seagate/seagate-exos-x-api-go/issues/7)
- Merge pull request #5 from Seagate/bug/fmw-47105 ([33673fd](https://github.com/Seagate/seagate-exos-x-api-go/commit/33673fd47b6a7ec8c7666f8cbcdf13d433b8ac8b)), closes [#5](https://github.com/Seagate/seagate-exos-x-api-go/issues/5)
- Merge pull request #4 from Seagate/feat/rebrand ([846ea75](https://github.com/Seagate/seagate-exos-x-api-go/commit/846ea75e2586fb7e7abd1960efbea7bf9a3ca644)), closes [#4](https://github.com/Seagate/seagate-exos-x-api-go/issues/4)
- Merge pull request #3 from Seagate/feat/rebrand ([57bb56e](https://github.com/Seagate/seagate-exos-x-api-go/commit/57bb56e53f15bb701c0cd9e4732ccb0aa05fc504)), closes [#3](https://github.com/Seagate/seagate-exos-x-api-go/issues/3)
- updated main test to use env variables over .env file, removed use of sha256 ([d09e637](https://github.com/Seagate/seagate-exos-x-api-go/commit/d09e637702272743ba3f830284f74597dd3dcb58))
- rebranding the api to use exosx ([0fbeaeb](https://github.com/Seagate/seagate-exos-x-api-go/commit/0fbeaeb94119372e07889992c30956be2e064907))
- Merge pull request #2 from Seagate/feat/relocate ([de2729f](https://github.com/Seagate/seagate-exos-x-api-go/commit/de2729f4545cd70bf9da517c11cd58f11d1e7997)), closes [#2](https://github.com/Seagate/seagate-exos-x-api-go/issues/2)
- Merge branch 'feat/snapshots' ([03fba9c](https://github.com/Seagate/seagate-exos-x-api-go/commit/03fba9c0c97aa332225fb51ff12bdc3a64562f0d))
- Merge branch 'fix/status' ([2660161](https://github.com/Seagate/seagate-exos-x-api-go/commit/2660161c617d1a41b93b59d85fea379a55ef77bd))
- Merge branch 'develop' into 'master' ([cc80886](https://github.com/Seagate/seagate-exos-x-api-go/commit/cc808866457138b9fa90bea572d2f2486630e971))
- Merge branch 'develop' into 'master' ([3b65b1a](https://github.com/Seagate/seagate-exos-x-api-go/commit/3b65b1a69e6b80a46d1f45e59d3f8145fd659857))
- Merge branch 'develop' into 'master' ([2355982](https://github.com/Seagate/seagate-exos-x-api-go/commit/2355982a416e4a09dcbe788dfa824f343a0f1bb1))
- Merge branch 'develop' into 'master' ([357b677](https://github.com/Seagate/seagate-exos-x-api-go/commit/357b677c079d31393873a27b97ac3d8998e4530b))
- Merge branch 'develop' into 'master' ([deb7d96](https://github.com/Seagate/seagate-exos-x-api-go/commit/deb7d96896a85440f27f189d3971922efb6c3623))
- Merge branch 'develop' into 'master' ([8e26c8c](https://github.com/Seagate/seagate-exos-x-api-go/commit/8e26c8c4822cecc7d7d16c3872ff8ed4787033fc)), closes [#2](https://github.com/Seagate/seagate-exos-x-api-go/issues/2)
- Merge branch 'develop' into 'master' ([8e25a51](https://github.com/Seagate/seagate-exos-x-api-go/commit/8e25a51186ecf7f54dfb13b091043c7d24354dfe))
- Merge branch 'develop' into 'master' ([3fbbde6](https://github.com/Seagate/seagate-exos-x-api-go/commit/3fbbde6631f14634ed9717678f5dd47f39f3c429))
- Merge branch 'develop' into 'master' ([248eb24](https://github.com/Seagate/seagate-exos-x-api-go/commit/248eb242ec80f0d3e8f948050c14ebb5d267c5fb))
- Merge branch 'develop' into 'master' ([a0bccca](https://github.com/Seagate/seagate-exos-x-api-go/commit/a0bccca0c657280e2ca89f66c5254592f1322c16))
- Initial commit ([e29ddef](https://github.com/Seagate/seagate-exos-x-api-go/commit/e29ddef4942ba82c32ed35aafa0185dc6942c1ac))
- **build:** add go module file ([12f7066](https://github.com/Seagate/seagate-exos-x-api-go/commit/12f70663e9e50289e91cf86c07a328ee061c4d8a))
- **build:** remove dep call from docker build ([73d7977](https://github.com/Seagate/seagate-exos-x-api-go/commit/73d797761eb64aa2c0b7cc0fbf22b0385f9b8aa2))
- **client:** automatically login if not already ([f12f970](https://github.com/Seagate/seagate-exos-x-api-go/commit/f12f97092a7681bf242d9f552c1c9bce72b8d54d))
- **client:** fill models from objects, not full response ([69fa0bc](https://github.com/Seagate/seagate-exos-x-api-go/commit/69fa0bc63c85d0ed052f2a5fdae87e4634c03461))
- **tests:** fix failing tests after go mod migration ([daf9d6c](https://github.com/Seagate/seagate-exos-x-api-go/commit/daf9d6c31182742b6ba9abf3a8236ff8a4f23676))

### Tests

- add tests for error status and refactor them using gomega ([f38efe2](https://github.com/Seagate/seagate-exos-x-api-go/commit/f38efe2d5fe882e568f3954af9b60caad1770be3))
- upgrade docker-compose version ([8716754](https://github.com/Seagate/seagate-exos-x-api-go/commit/8716754619e1616f1abd18881d180d82720da7ee))

## [1.0.1](https://github.com/Seagate/seagate-exos-x-api-go/compare/v1.0.0...v1.0.1) (2021-10-02)

### Bug Fixes

- handle session timeout error code (2) ([5a09c9d](https://github.com/Seagate/seagate-exos-x-api-go/commit/5a09c9da207cc66aad664f01d162c50b5eaf9227))

### Chores

- remove changelog workflow ([0d89814](https://github.com/Seagate/seagate-exos-x-api-go/commit/0d898149b7280a46bb5e924143fd84601bf5a830))

### Other

- Merge pull request #12 from Seagate/fix/session-timeout ([70cdf3f](https://github.com/Seagate/seagate-exos-x-api-go/commit/70cdf3fd02616d12e31498e28c4630e23f9c44ee)), closes [#12](https://github.com/Seagate/seagate-exos-x-api-go/issues/12)
- added main as release branch ([6bb7775](https://github.com/Seagate/seagate-exos-x-api-go/commit/6bb7775ea6d951b9a98fa8db3d8483324c893b77))
- add checkout step to workflow ([4a703a7](https://github.com/Seagate/seagate-exos-x-api-go/commit/4a703a746b13d3e11e96c5536d12f8c7464b2a88))
- Corrected release workflow ([fcc773f](https://github.com/Seagate/seagate-exos-x-api-go/commit/fcc773fe0c4b62ad9104f5b8fa957caf43e23677))
- Corrected release workflow ([346d4c9](https://github.com/Seagate/seagate-exos-x-api-go/commit/346d4c95b298f1d0f85314bd0af5b29b5c7e83da))
- Use GitHub Actions Release me! ([b8a2cb4](https://github.com/Seagate/seagate-exos-x-api-go/commit/b8a2cb4f40322ca3892053e27b80e63956df41e5))
- github workflow update ([1dff4e3](https://github.com/Seagate/seagate-exos-x-api-go/commit/1dff4e33024cd17f578e3be1803b0f982245a93a))
- adding npm package files for semantic-release ([4ae2307](https://github.com/Seagate/seagate-exos-x-api-go/commit/4ae230785ad67c45e1e224a01895c9ad85820778))
- release and changelog workflow updates ([777a1eb](https://github.com/Seagate/seagate-exos-x-api-go/commit/777a1eb63a850cc734102d16bcade91e0da4e458))
- github release workflow ([f456d8f](https://github.com/Seagate/seagate-exos-x-api-go/commit/f456d8fa3f7e57733f48d9356b1227646155dcb5))
