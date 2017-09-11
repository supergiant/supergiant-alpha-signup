#!/usr/bin/env bash

# Travis deployment script. After test success actions go here.

TAG=${TRAVIS_BRANCH:-unstable}
BUILD_TARGETS=darwin/amd64,linux/amd64
LDFLAGS="-X main.version=$(git describe --tags)"

cd ui/assets && npm install && ng build --env=prod  && cd ../..
go-bindata -pkg ui -o bindata/ui/bindata.go ui/assets/dist/...

xgo -ldflags "$LDFLAGS" -go 1.7.1 --targets=$BUILD_TARGETS -out dist/alpha-ui ./ui
xgo -ldflags "$LDFLAGS" -go 1.7.1 --targets=$BUILD_TARGETS -out dist/alpha-api ./api



echo "Tag Name: ${TAG}"
if [[ "$TAG" =~ ^v[0-100]. ]]; then
  echo "global deploy"
  # ./packer build build/build_release.json
else
  echo "private unstable"
  docker login -u $DOCKER_USER -p $DOCKER_PASS

  ## UI Docker Build
  REPO=supergiant/supergiant-alpha-signup
  cp dist/alpha-ui-linux-amd64 build/ui/
  docker build -t $REPO:$TAG-linux-x64 build/ui/
  docker push $REPO

  REPO=supergiant/supergiant-alpha-api
  cp dist/alpha-api-linux-amd64 build/api/
  docker build -t $REPO:$TAG-linux-x64 build/api/
  docker push $REPO
fi
