echo off
set GOPATH=%cd%
cd bin\Windows
echo building...
go build main
echo done.
pause