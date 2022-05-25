# OPA Gin Example

This repository shows how to integrate a http service written in [Gin](https://github.com/gin-gonic/gin) with the OPA to perform API authorization.

## Build
```
make
```

## Run
```
make run
```
or `go run main.go`

test opa:
1. None resource request is allowed, `curl http://localhost:8080/`
2. Unauthenticated user is not allowed for api resource, `curl http://localhost:8080/api/users`
3. Authenticated user get resource is allowed, `curl http://localhost:8080/api/users?user=bob`

## Authz with OPA
Authentication rego config in (authz.rego)[./authz/authz.rego]

Gin middleware code in [server.go](./server/opa.go#L14)
