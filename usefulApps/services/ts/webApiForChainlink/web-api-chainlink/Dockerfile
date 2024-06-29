FROM node:alpine

WORKDIR /

COPY package*.json ./
RUN npm i

COPY nest-cli.json ./
COPY tsconfig.json ./
COPY tsconfig.build.json ./
COPY .eslintrc.js ./

COPY ./src ./src

RUN npm run build

CMD ["npm","run","start:prod"]

EXPOSE 3000
