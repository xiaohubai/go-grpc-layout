GOPATH=$(go env GOPATH)
mkdir temp && cd temp
wget https://github.com/protocolbuffers/protobuf/releases/download/v26.0/protoc-26.0-linux-x86_64.zip
unzip protoc-26.0-linux-x86_64.zip

mv bin/protoc $GOPATH/bin
cd ../
rm -rf temp