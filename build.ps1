npm ci
npm run build

$build_dir="build"
if(Test-Path $build_dir){
    Remove-Item -Force -Recurse $build_dir
}
New-Item -Name $build_dir -ItemType "directory"

$importpath="github.com/scutrobotlab/asuwave/helper"
$gittag=$(git describe --tags --abbrev=0)
$build_prefix="asuwave_"
$os_list="linux", "darwin", "windows"
$arch_list="amd64", "arm64"
$flags="-w -s -X '${importpath}.GitTag=${gittag}' -X '${importpath}.GitHash=$(git describe --tags --long)' -X '${importpath}.BuildTime=$(Get-Date -Format 'yyyy-MM-dd HH:mm')' -X '${importpath}.GoVersion=$(go version)'"

foreach ($os in $os_list) {
    foreach ($arch in $arch_list) {
        $suffix=""
        if ($os -eq "windows") {
            if($arch -eq "arm64"){
                continue
            }else{
                $suffix=".exe"
            }
        }
        $file="${build_dir}\${build_prefix}${gittag}_${os}_${arch}"
        $out="${build_dir}\${build_prefix}${os}_${arch}${suffix}"
        $Env:CGO_ENABLED=0
        $Env:GOOS=${os}
        $Env:GOARCH=${arch}
        go build -tags release -ldflags="$flags" -o $out
        Compress-Archive -CompressionLevel "Optimal" -Path $out -DestinationPath "${file}.zip"
    }
}

