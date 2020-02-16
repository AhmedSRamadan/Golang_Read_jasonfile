FROM golang:alpine
WORKDIR /build
COPY file.json .
COPY smartfrog.go .
RUN go build -o main .
COPY script.sh .
CMD ["./script.sh"]
