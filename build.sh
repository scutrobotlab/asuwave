#!/bin/bash

npm ci
npm run build

cp ./server/server.go ./server/server.go.org
sed -i 's/http.Dir(\".\/dist\/\")/assetFS()/g' ./server/server.go

go get github.com/go-bindata/go-bindata/...
go get github.com/elazarl/go-bindata-assetfs/...
go-bindata-assetfs -pkg server -o ./server/bindata.go ./dist/...
go mod tidy

build_dir="build"

if [[ ! -d $build_dir ]]; then
    mkdir $build_dir
else
    rm -rf $build_dir
fi

build_prefix="asuwave_"
os_list=("linux" "darwin" "windows")
build_suffix=("" "" ".exe")
flags="-w -s -X 'main.githash=$(git describe --tags --long)' -X 'main.buildtime=$(date)' -X 'main.goversion=$(go version)'"

for ((i = 0 ; i < 3 ; i++)); do
    file=$build_dir/$build_prefix${os_list[$i]}
    out=$file${build_suffix[$i]}
    CGO_ENABLED=0 GOOS=${os_list[$i]} GOARCH=amd64 go build -ldflags="$flags" -o $out
    upx -q -9 $out
    zip -j -9 $file.zip $out
done

rm -f ./server/server.go
mv ./server/server.go.org ./server/server.go
rm -f ./server/bindata.go
go mod tidy
