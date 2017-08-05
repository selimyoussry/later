#!/bin/sh

# Python package code
echo "# Automatically generated -- START" > __init__.py
echo "import sys" >> __init__.py
echo "import os" >> __init__.py
echo "sys.path.append(os.path.dirname(os.path.realpath(__file__)))" >> __init__.py
echo "# Automatically generated -- END" >> __init__.py

# Python
$WORKON_HOME/grpc/bin/python -m grpc_tools.protoc -I=. -I=$GOPATH/src/ --python_out=. --grpc_python_out=. ./later.proto

# Go
protoc -I=. -I=$GOPATH/src/ ./later.proto --go_out=plugins=grpc:$GOPATH/src
go install

# Generate reverse proxy
protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --grpc-gateway_out=logtostderr=true:. \
  ./later_json.proto
go install

# Swagger definition
protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --swagger_out=logtostderr=true:. \
  ./later_json.proto
