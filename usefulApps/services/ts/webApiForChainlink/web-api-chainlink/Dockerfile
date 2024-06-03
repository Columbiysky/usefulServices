FROM node:alpine

WORKDIR /

COPY package*.json ./
RUN npm i

COPY nest-cli.json ./
COPY tsconfig.json ./


COPY ./src ./src

RUN npm run build

EXPOSE 30003
