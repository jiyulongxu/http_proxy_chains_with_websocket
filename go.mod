module github.com/n454149301/http_proxy

go 1.16

require (
	golang.org/x/net v0.0.0
)

replace (
	golang.org/x/net => github.com/golang/net v0.0.0-20200222033325-078779b8f2d8
)
