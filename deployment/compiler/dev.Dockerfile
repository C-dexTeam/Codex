FROM ubuntu:22.04-slim

WORKDIR /app

RUN apt-get update && apt-get upgrade -y && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*
