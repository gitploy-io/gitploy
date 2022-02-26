#!/bin/sh
set -e
set -x

# disable CGO for cross-compiling
export CGO_ENABLED=0

# compile for all architectures
GOOS=linux   GOARCH=amd64   go build -ldflags "-X main.Version=${GITHUB_REF_NAME##v}" -o release/linux/amd64/gitploy       ./cmd/cli
GOOS=linux   GOARCH=arm64   go build -ldflags "-X main.Version=${GITHUB_REF_NAME##v}" -o release/linux/arm64/gitploy       ./cmd/cli
GOOS=linux   GOARCH=ppc64le go build -ldflags "-X main.Version=${GITHUB_REF_NAME##v}" -o release/linux/ppc64le/gitploy     ./cmd/cli
GOOS=linux   GOARCH=arm     go build -ldflags "-X main.Version=${GITHUB_REF_NAME##v}" -o release/linux/arm/gitploy         ./cmd/cli
GOOS=windows GOARCH=amd64   go build -ldflags "-X main.Version=${GITHUB_REF_NAME##v}" -o release/windows/amd64/gitploy.exe ./cmd/cli
GOOS=darwin  GOARCH=amd64   go build -ldflags "-X main.Version=${GITHUB_REF_NAME##v}" -o release/darwin/amd64/gitploy      ./cmd/cli
GOOS=darwin  GOARCH=arm64   go build -ldflags "-X main.Version=${GITHUB_REF_NAME##v}" -o release/darwin/arm64/gitploy      ./cmd/cli

# tar binary files prior to upload
tar -cvzf release/gitploy_linux_amd64.tar.gz   -C release/linux/amd64   gitploy
tar -cvzf release/gitploy_linux_arm64.tar.gz   -C release/linux/arm64   gitploy
tar -cvzf release/gitploy_linux_ppc64le.tar.gz -C release/linux/ppc64le gitploy
tar -cvzf release/gitploy_linux_arm.tar.gz     -C release/linux/arm     gitploy
tar -cvzf release/gitploy_windows_amd64.tar.gz -C release/windows/amd64 gitploy.exe
tar -cvzf release/gitploy_darwin_amd64.tar.gz  -C release/darwin/amd64  gitploy
tar -cvzf release/gitploy_darwin_arm64.tar.gz  -C release/darwin/arm64  gitploy

# generate shas for tar files
sha256sum release/*.tar.gz > release/gitploy_checksums.txt