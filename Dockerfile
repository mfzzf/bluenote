FROM ubuntu:latest
LABEL authors="fzzf"
# 复制二进制文件
COPY bluenote /bin
# 设置工作目录
WORKDIR /bin
ENTRYPOINT ["/app/bluenote"]
