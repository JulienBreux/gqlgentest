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

## Error trace

```text
2018/08/16 02:24:22 http: panic serving 127.0.0.1:53242: reflect: call of reflect.Value.Type on zero Value
goroutine 4 [running]:
net/http.(*conn).serve.func1(0xc420198000)
	/usr/local/Cellar/go/1.10.3/libexec/src/net/http/server.go:1726 +0xd0
panic(0x137e4c0, 0xc42000c740)
	/usr/local/Cellar/go/1.10.3/libexec/src/runtime/panic.go:502 +0x229
reflect.Value.Type(0x0, 0x0, 0x0, 0x6, 0xc420190088)
	/usr/local/Cellar/go/1.10.3/libexec/src/reflect/value.go:1713 +0x16d
github.com/vektah/gqlparser/validator.(*varValidator).validateVarType(0xc420179bf8, 0xc420157110, 0x0, 0x0, 0x0, 0xc42000c701)
	/Users/x/go/src/github.com/vektah/gqlparser/validator/vars.go:112 +0x97a
github.com/vektah/gqlparser/validator.VariableValues(0xc420156b70, 0xc420185180, 0xc4201570e0, 0xc420018268, 0x3)
	/Users/x/go/src/github.com/vektah/gqlparser/validator/vars.go:54 +0x2e4
github.com/99designs/gqlgen/handler.GraphQL.func1(0x143cec0, 0xc4201880e0, 0xc42016c200)
	/Users/x/go/src/github.com/99designs/gqlgen/handler/graphql.go:172 +0x320
net/http.HandlerFunc.ServeHTTP(0xc42000c260, 0x143cec0, 0xc4201880e0, 0xc42016c200)
	/usr/local/Cellar/go/1.10.3/libexec/src/net/http/server.go:1947 +0x44
net/http.(*ServeMux).ServeHTTP(0x160f960, 0x143cec0, 0xc4201880e0, 0xc42016c200)
	/usr/local/Cellar/go/1.10.3/libexec/src/net/http/server.go:2337 +0x130
net/http.serverHandler.ServeHTTP(0xc4201705b0, 0x143cec0, 0xc4201880e0, 0xc42016c200)
	/usr/local/Cellar/go/1.10.3/libexec/src/net/http/server.go:2694 +0xbc
net/http.(*conn).serve(0xc420198000, 0x143d180, 0xc420187340)
	/usr/local/Cellar/go/1.10.3/libexec/src/net/http/server.go:1830 +0x651
created by net/http.(*Server).Serve
	/usr/local/Cellar/go/1.10.3/libexec/src/net/http/server.go:2795 +0x27b
```