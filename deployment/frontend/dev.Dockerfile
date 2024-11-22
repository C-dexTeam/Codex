FROM node:22.11-alpine AS builder
ENV TZ=UTC 
WORKDIR /app

COPY ./package.json /app

COPY . .

CMD npm install && npm run dev
