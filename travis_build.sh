#!/usr/bin/env bash

# Travis deployment script. After test success actions go here.

TAG=${TRAVIS_BRANCH:-unstable}
BUILD_TARGETS=darwin/amd64,linux/amd64
LDFLAGS="-X main.version=$(git describe --tags)"

echo "Tag Name: ${TAG}"
if [[ "$TAG" =~ ^v[0-100]. ]]; then
  echo "global deploy"
  ./packer build build/build_release.json
else
  xgo -ldflags "$LDFLAGS" -go 1.7.1 --targets=$BUILD_TARGETS -out dist/supergiant-ui ./cmd/ui
  echo "private unstable"
  docker login -u $DOCKER_USER -p $DOCKER_PASS

  ## UI Docker Build
  REPO=supergiant/supergiant-alpha-signup
  cp dist/supergiant-ui-linux-amd64 build/docker/ui/linux-amd64/
  docker build -t $REPO:$TAG-linux-x64 build/docker/ui/linux-amd64/
  docker push $REPO
fi
