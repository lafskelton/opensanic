#!/bin/bash
protoc -I sanicdb/ sanicdb/sanic.proto --go_out=plugins=grpc:sanicdb --experimental_allow_proto3_optional