#!/usr/bin/env bash

export GOPATH=$HOME/go/
protoc --proto_path=$GOPATH/src:. --twirp_out=. --go_out=. ./protobuf/pdefs/*.proto
