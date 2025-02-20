FROM ubuntu:latest
LABEL authors="fzzf"

ENTRYPOINT ["top", "-b"]