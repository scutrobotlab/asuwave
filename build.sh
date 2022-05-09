#!/bin/bash

importpath="github.com/scutrobotlab/asuwave/internal/helper"
build_prefix="asuwave_"
os_list=("linux" "darwin" "windows")
arch_list=("amd64" "arm64")
if [ -v $1 ];
then
    echo "No"
    gittag=`git describe --tags --abbrev=0`
else
    echo "Yes"
    gittag=$1
fi

echo ${gittag}
sed -i "s/VUE_APP_GITTAG=.*/VUE_APP_GITTAG=${gittag}/g" .env
cat .env

npm ci
npm run build

build_dir="build"

if [[ ! -d $build_dir ]]; then
    mkdir $build_dir
else
    rm -rf $build_dir
fi

flags="-w -s -X '${importpath}.GitTag=${gittag}' -X '${importpath}.GitHash=$(git describe --tags --long)' -X '${importpath}.BuildTime=$(date +'%Y-%m-%d %H:%M')' -X '${importpath}.GoVersion=$(go version)'"

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
        file=$build_dir/$build_prefix${gittag}_${os}_${arch}
        out=$build_dir/$build_prefix${os}_${arch}$suffix
        CGO_ENABLED=0 GOOS=${os} GOARCH=${arch} go build -tags release -ldflags="$flags" -o $out
        zip -j -9 $file.zip $out
    done
done
