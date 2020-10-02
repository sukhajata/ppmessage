export PATH=$PATH:~/go/bin

cd ppdownlink

protoc --go_out=plugins=grpc:. *.proto 

#move generated file to ppdownlink

python -m grpc_tools.protoc -I . --python_out=python --grpc_python_out=python *.proto

cd ../ppuplink

...