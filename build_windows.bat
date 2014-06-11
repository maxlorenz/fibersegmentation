echo off
set GOPATH=%cd%
cd bin
echo Building...
go build -o fibersegmentation_windows.exe main
echo Done. Press any key to continue.
pause