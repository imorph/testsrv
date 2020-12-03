# testsrv service

Test http service.
Will listen on :8080

## Build

```shell
go build
```

## Run

```shell
./testsrv
```

## Paths available

* `/`
* `/articles`
* `/metrics`

### /

will print "Welcome!"

### /articles

google results for search: "ansible+articles"

### /metrics

Prometheus formatted metrics for application

