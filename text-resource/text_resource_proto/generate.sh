#!/bin/bash
echo "generating text_resource_proto"
protoc -I . text_resource.proto --go_out=plugins=http:.
sed -i "" "s/,omitempty//" text_resource.pb.go
