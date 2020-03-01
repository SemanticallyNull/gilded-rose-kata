# GO Starter

- Run :

```shell
go run gilded-rose.go
```

- Run tests :

```shell
go test ./...
```
**OR**
```shell
go get github.com/onsi/ginkgo/ginkgo
ginkgo -r
```

- Run tests and coverage :

```shell
go test ./... -coverprofile=coverage.out

go tool cover -html=coverage.out
```
