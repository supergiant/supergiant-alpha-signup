#!/bin/bash

cd ui && npm install && ng build --env=prod  && cd ..

go-bindata -pkg ui -o bindata/ui/bindata.go ui/dist/...

go run server.go
