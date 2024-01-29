FROM node:lts-alpine

WORKDIR /

COPY package*.json ./
RUN npm i

COPY nuxt.config.ts ./
COPY tsconfig.json ./
COPY app.vue ./
COPY .npmrc ./


COPY ./data/exampleObject/*.ts ./data/exampleObject/
COPY ./logic/webhooks/httpWrappers/*.ts ./logic/webhooks/httpWrappers/

COPY ./pages/login/*.vue ./pages/login/
COPY ./pages/youtubeDownloader/*.vue ./pages/youtubeDownloader/
COPY ./pages/*.vue ./pages/

COPY ./plugins/*.js ./plugins/
COPY ./public/*.ico ./public/

RUN npm run build

EXPOSE 30002
