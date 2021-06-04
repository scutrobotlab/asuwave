#!/bin/bash

npm ci
npm run build

build_dir="build"

if [[ ! -d $build_dir ]]; then
    mkdir $build_dir
else
    rm -rf $build_dir
fi

build_prefix="asuwave_"
os_list=("linux" "darwin" "windows")
arch_list=("amd64" "arm64")
flags="-w -s -X 'main.githash=$(git describe --tags --long)' -X 'main.buildtime=$(date)' -X 'main.goversion=$(go version)'"

for os in ${os_list[@]}; do
    for arch in ${arch_list[@]}; do
        suffix=""
        if [ "$os" == "windows" ] ; then
            if [ "$arch" == "arm64" ]; then
                continue
            else
                suffix=".exe"
            fi
        fi
        file=$build_dir/$build_prefix${os}_${arch}
        out=$file$suffix
        CGO_ENABLED=0 GOOS=${os} GOARCH=${arch} go build -tags release -ldflags="$flags" -o $out
        upx -q -9 $out > /dev/null
        zip -j -9 $file.zip $out
    done
done
