FROM golang AS build-backend
COPY . /app
WORKDIR /app
RUN GO111MODULE=on go mod download
RUN GO111MODULE=on go build -o overtype

FROM node:12.18.2 as build-frontend
WORKDIR /app
COPY web .
RUN npm install
RUN npm run build
RUN npm run export

FROM alpine
RUN apk add --no-cache libc6-compat
WORKDIR /
COPY --from=build-backend /app/overtype .
COPY --from=build-frontend /app/__sapper__/export /web_build
CMD ./overtype
