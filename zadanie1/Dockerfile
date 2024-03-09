FROM ubuntu:latest

ENV TZ = "Europe/Warsaw"
ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && \
    apt-get -y install software-properties-common

RUN add-apt-repository ppa:deadsnakes/ppa && \
    apt-get -y install python3.8

RUN update-alternatives --install /usr/bin/python python /usr/bin/python3.8 1 && \
    update-alternatives --set python /usr/bin/python3.8

RUN update-alternatives --install /usr/bin/python3 python3 /usr/bin/python3.8 1

WORKDIR /workspace
