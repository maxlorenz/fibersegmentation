export GOPATH=$(pwd -P)
cd bin
clear
echo Building...
go build -o fibersegmentation_linux main
echo Done. Press any key to continue.
read