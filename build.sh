#!/bin/sh

npm ci
npm run build

cp ./server/server.go ./server/server.go.org
sed -i 's/http.Dir(\".\/dist\/\")/assetFS()/g' ./server/server.go

go get github.com/go-bindata/go-bindata/...
go get github.com/elazarl/go-bindata-assetfs/...
go-bindata-assetfs -pkg server -o ./server/bindata.go ./dist/...
go mod tidy

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o asuwave_linux
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-w -s" -o asuwave_mac
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-w -s" -o asuwave_windows.exe
upx -q -9 asuwave_linux asuwave_mac asuwave_windows.exe
zip asuwave_linux.zip asuwave_linux
zip asuwave_mac.zip asuwave_mac
zip asuwave_windows.zip asuwave_windows.exe

rm -f ./server/server.go
mv ./server/server.go.org ./server/server.go
rm -f ./server/bindata.go
go mod tidy
