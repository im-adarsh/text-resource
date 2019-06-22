#!/bin/bash
echo "generating proto"
protoc -I . text_resource.proto --go_out=plugins=http:.
sed -i "" "s/,omitempty//" text_resource.pb.go
