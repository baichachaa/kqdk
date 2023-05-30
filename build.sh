Revision=$(git rev-parse HEAD)
Branch=$(git rev-parse --abbrev-ref HEAD)
BuildDate=$(date +'%Y-%m-%d %H:%M:%d')

Version=1.0.1

echo $Revision
echo $Branch
echo $BuildDate
echo $Version

# 将以上变量序列化至 LDFlags 变量中
LDFlags=" \
    -X 'kqdk/app.Revision=${Revision}' \
    -X 'kqdk/app.Branch=${Branch}' \
    -X 'kqdk/app.BuildDate=${BuildDate}' \
    -X 'kqdk/app.Version=${Version}' \
"

go.exe build -ldflags "-s -w ${LDFlags}" -o ".\dist\release_amd64_win.exe"

go env -w GOARCH=arm64
go env -w GOOS=linux
go.exe build -ldflags "-s -w ${LDFlags}" -o ".\dist\release_arm64_linux"
go env -u GOARCH
go env -u GOOS
