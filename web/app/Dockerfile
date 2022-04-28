FROM node:16-alpine as build

WORKDIR /app

ENV PATH /app/node_modules/.bin:$PATH
COPY *.json ./
COPY yarn.lock ./

RUN yarn
