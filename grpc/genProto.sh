#!/usr/bin/env bash

# 获取项目绝对路径
root_path=$(pwd)

# 编译文件位置和生成*.pb.go文件位置
build_generate_path=${root_path}'/proto/loginProto'

# 编译文件位置
file_path=${build_generate_path}'/*.proto'

protoc -I ${build_generate_path} ${file_path} --go_out=plugins=grpc:${build_generate_path}