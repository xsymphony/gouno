#!/bin/sh

echo "\033[34m 开始编译 \033[0m"

project_path="$PWD"
project_name="${project_path##*/}"
to_path="/go/src/github.com/xiongsyao/$project_name"

echo "\033[34m 当前文件目录是:$project_path \n 项目名称为:$project_name \033[0m"

echo "\033[34m 开始构建容器并编译项目... \033[0m"

sudo docker run --rm -v "$project_path":"$to_path" \
                     -w "$to_path" golang:1.8 go build -v .