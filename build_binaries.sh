#!/bin/bash

binary_name=sau
OSes=("linux" "darwin" "windows")

archs=("amd64" "386")

build() {
  local os=$1
  local arch=$2
  echo "Building for $os/$arch..."
  GOOS=$os GOARCH=$arch go build -o "$binary_name-$os-$arch"
  echo "Done."
}
for os in "${OSes[@]}"; do
  for arch in "${archs[@]}"; do
    build $os $arch
  done
done

echo "All builds complete."
