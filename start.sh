#!/bin/bash

# 杀死占用端口 8080 的进程
echo "Checking for processes using port 8080..."
lsof -i :8080 | grep LISTEN | awk '{print $2}' | xargs -r kill -9

# 确保杀死进程后端口被释放
echo "Waiting for port 8080 to be released..."
while lsof -i :8080 > /dev/null; do
    sleep 1
done

# 启动 fresh
echo "Starting fresh..."
fresh