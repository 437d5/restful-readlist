FROM ubuntu:latest
LABEL authors="slavebook"

ENTRYPOINT ["top", "-b"]