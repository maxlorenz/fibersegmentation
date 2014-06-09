export GOPATH=$(pwd -P)
cd bin/Linux
clear
echo building...
go build main
echo done.
read -p