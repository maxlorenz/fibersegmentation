export GOPATH=$(pwd -P)
cd bin/OSX
clear
echo building...
go build main
echo done.
read -p