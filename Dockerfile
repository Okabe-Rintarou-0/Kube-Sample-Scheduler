FROM golang:1.21 as builder

WORKDIR /scheduler

COPY . /scheduler

RUN go env -w GOPROXY=https://goproxy.cn

RUN CGO_ENABLED=0 go build -o app main.go

FROM alpine

COPY --from=builder /scheduler/app /scheduler

ENTRYPOINT [ "/scheduler" ]