FROM golang:1.20.2
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
EXPOSE 8080
ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY etc /app/etc
COPY assets /app/assets

FROM google/dart:2.12
WORKDIR /app
COPY . .
RUN pub get
RUN pub build
EXPOSE 8081

FROM mysql:8.0.33-debian
ENV MYSQL_ROOT_PASSWORD example
EXPOSE 3306