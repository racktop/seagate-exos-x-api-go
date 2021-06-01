# seagate-exos-x-api-go

[![Go Report Card](https://goreportcard.com/badge/github.com/Seagate/seagate-exos-x-api-go)](https://goreportcard.com/report/github.com/Seagate/seagate-exos-x-api-go)
[![Go Reference](https://pkg.go.dev/badge/github.com/Seagate/seagate-exos-x-api-go.svg)](https://pkg.go.dev/github.com/Seagate/seagate-exos-x-api-go)

A Go implementation of the [Seagate EXOS X API](https://www.seagate.com/files/www-content/support-content/raid-systems/_shared/documentation/83-00007047-13-01_G265_SMG.pdf).

## Test Using A Live System

This option runs the Go language test cases against a live storage system. Two steps are required:
- Update .env with the correct system IP Address and credentials
- Run `go test -v`

Another option is to define environment variables, which take precedence over .env values
- export TEST_STORAGEIP=http:/<ipaddress>
- export TEST_USERNAME=<username>
- export TEST_USERNAME=<password>
- Run `go test -v`
- unset TEST_STORAGEIP TEST_PASSWORD TEST_USERNAME


## Test Using a Mock Server

### Using node.js

You can run tests with docker-compose:

```sh
docker-compose up --build --abort-on-container-exit --exit-code-from tests
```

### Using node.js

In order to run tests against a mock server, you will need to install node.js and npm to run the mock server. When it's done, go to the `mock` directory, install dependencies and start the mock server.

```sh
cd ./mock
npm install
npm run start
```

- Update .env with an IP Address of `localhost:8080` and correct credentials
- You're now ready to go, just run `go test -v` to run the tests suite.

