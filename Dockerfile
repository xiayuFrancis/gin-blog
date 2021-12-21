FROM golang:latest as builder

ENV GO111MODULE=on \
    GOPROXY=http://goproxy.cn,direct

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -a -installsuffix cgo  .

RUN mkdir publish
RUN cp -r conf publish
RUN cp gin-example publish

FROM alpine
WORKDIR /app
COPY --from=builder /app/publish .

EXPOSE 8000
ENTRYPOINT ["/app/gin-example"]