FROM golang:alpine as builder

RUN apk update && apk add --no-cache tzdata && apk add --no-cache git

WORKDIR /app

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .

RUN GOOS=linux go build -o bin/audio-library ./cmd

FROM alpine:3.18

RUN apk update && apk add --no-cache libstdc++ ffmpeg python3

COPY --from=builder /internal/app/bin/audio-library /usr/local/bin/audio-library
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Ashgabat /usr/share/zoneinfo/Asia/Ashgabat
ENV TZ Asia/Ashgabat

COPY --from=builder /internal/app/configs/local.yaml /etc/config.yaml

CMD ["/usr/local/bin/audio-library", "-config", "/etc/app/config.yaml"]
