#!/bin/bash
echo "start build for linux/windows/mac!"

echo "build server"
cd src/server
# 添加图标.教程: https://blog.csdn.net/u014633966/article/details/82984037
rsrc -ico icon.ico -o main.syso

# 编译各版本服务器端
echo "build server to linux"
export  GOOS=linux
export  GOARCH=amd64
go build -ldflags "-s -w" -o ../../bin/server_linux
echo "build server to windows"
export  GOOS=windows
export  GOARCH=amd64
go build -ldflags "-s -w -H windowsgui" -o ../../bin/server_windows.exe
echo "build server to mac"
export  GOOS=darwin
export  GOARCH=amd64
go build -ldflags "-s -w"  -o ../../bin/server_mac
echo "build server finished!"

echo "build client"
cd ../client
# 添加图标.教程: https://blog.csdn.net/u014633966/article/details/82984037
rsrc -ico icon.ico -o main.syso

# 编译各版本客户端
echo "build client to linux"
export  GOOS=linux
export  GOARCH=amd64
go build -ldflags "-s -w" -o ../../bin/client_linux
echo "build client to windows"
export  GOOS=windows
export  GOARCH=amd64
go build -ldflags "-s -w -H windowsgui" -o ../../bin/client_windows.exe
echo "build client to mac"
export  GOOS=darwin
export  GOARCH=amd64
go build -ldflags "-s -w" -o ../../bin/client_mac
echo "build client finished!"

# 加壳
echo "start upx ..."
cd ../../bin
upx -9 -k server_windows.exe
upx -9 -k server_linux
upx -9 -k server_mac
upx -9 -k client_windows.exe
upx -9 -k client_linux
upx -9 -k client_mac

echo "app build to path[ go_line_chat/bin ]"

