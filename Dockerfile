FROM golang:1.13-alpine as build-stage

ENV GOPROXY=https://goproxy.cn

WORKDIR /go/src/backend

COPY . .

RUN go mod tidy && go build -o backend .

FROM alpine:3.10

ENV TZ=Asia/Shanghai

COPY --from=build-stage /go/src/backend .

CMD ["./backend"]
