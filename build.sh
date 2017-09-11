#!/bin/bash

cd ui && npm install && ng build --aot --prod && cd ../

go-bindata -pkg ui -o bindata/ui/bindata.go ui/dist/...

go run server.go
