### Requirements:

- Golang 18 +
- Docker

### For Unit Tests Install:
- [Ginkgo](https://onsi.github.io/ginkgo/#installing-ginkgo)
- [Mockgen](https://github.com/golang/mock#go-116)
- run `go generate ./...` to generate mock files
- run `ginkgo ./...` to run all tests
### Steps:
- run `docker-compose up` to build up Kafka Broker
- run `go run main.go` to run the application
- open on your browser `localhost:3001`