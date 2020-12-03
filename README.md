# testsrv service

Test http service.
Will listen on :8080

## Build

```
go build
```

## Run

./testsrv

## Paths available

* /
* /articles
* /metrics

### /

will print "Welcome!"

### /articles

google search for "ansible+articles"

### /metrics

Prometheus formatted metrics for application

