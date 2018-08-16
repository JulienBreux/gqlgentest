# GQLGEN Test

```bash
brew install dep
brew upgrade dep
dep ensure
go run main.go
```

Go to [test](http://0.0.0.0:8081).

## To reproduce

```bash
curl 'http://0.0.0.0:8081/query' --data-binary $'{"operationName":"Get","variables":{"name":null},"query":"query Get($name: String) { Get(name: $name) { ...ObjectReturned __typename }} fragment ObjectReturned on ObjectReturned { name }"}' --compressed
```