Revision=$(git rev-parse HEAD)
Branch=$(git rev-parse --abbrev-ref HEAD)
BuildDate=$(date +'%Y-%m-%d %H:%M:%d')


echo $Revision
echo $Branch
echo $BuildDate

# 将以上变量序列化至 LDFlags 变量中
LDFlags=" \
    -X 'kqdk/app.Revision=${Revision}' \
    -X 'kqdk/app.Branch=${Branch}' \
    -X 'kqdk/app.BuildDate=${BuildDate}' \
"

"C:\Program Files\Go\bin\go.exe" build -ldflags "-s -w ${LDFlags}" -o "C:\Users\hy\GolandProjects\kqdk\dist\release_win.exe"

