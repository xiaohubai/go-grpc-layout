GOPATH=$(go env GOPATH)
mkdir temp && cd temp
wget https://github.com/protocolbuffers/protobuf/releases/download/v23.2/protoc-23.2-linux-x86_64.zip
unzip protoc-23.2-linux-x86_64.zip

rm -rf $GOPATH/bin/protoc
mv bin/protoc $GOPATH/bin
cd ../
rm -rf temp