
Some useful links:
 - https://golang.org/doc/code
 - Create module outside GOPATH
 - GOPATH is the cache of modules that are used
 - Configure GOBIN for bin output
 - go mod init github.com/rjtokenring/goms (better to make it already ready to be pushed on git)
 - go install github.com/rjtokenring/goms
 - package main app should be in the root of the project
 - then add other packages to the module and use them with the complete path (eg import github.com/rjtokenring/goms/stringstxt)

###### **Open Api**

https://github.com/deepmap/oapi-codegen

Steps:
 - go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen
 - oapi-codegen will be generated into $GOBIN dir
 - ./bin/oapi-codegen -generate types,server Openapi/openapi.yaml > Openapi/oapi.gen.go (just generate server stubs)
 - or: ./bin/oapi-codegen --config openapi/codegen-config.yaml openapi/openapi.yaml
