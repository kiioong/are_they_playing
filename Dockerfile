FROM golang:1.24.1-alpine3.21 AS go-server

WORKDIR /usr/src/app

COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend/. .
RUN go build -v -o /usr/local/bin/app ./internal/server.go

EXPOSE 9090

CMD ["app"]

FROM node:23-alpine3.21 AS node

WORKDIR /app

COPY frontend/package.json frontend/yarn.lock ./
RUN yarn install

COPY frontend/. .

EXPOSE 80

CMD ["./node_modules/@ionic/cli/bin/ionic", "serve","--host", "0.0.0.0", "--port", "80"]