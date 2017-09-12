#!/bin/bash

cd ui && npm install && ng build --prod  && cd ..

go-bindata -pkg ui -o bindata/ui/bindata.go assets/dist/...

go run server.go
