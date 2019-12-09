# lapis builder image
FROM golang:1.13.4 as builder
LABEL maintainer "plainbanana <kazukidegozaimasuruzo@gmail.com>"
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /plainbanana/lapis
COPY . .
RUN make

# lapis image
# docker run ${containername} --env-file .env
FROM alpine
LABEL maintainer "plainbanana <kazukidegozaimasuruzo@gmail.com>"
ENV DOTENV=false
RUN apk add --no-cache ca-certificates
COPY --from=builder /plainbanana/lapis/lapis /lapis
CMD ["/lapis"]