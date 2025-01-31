FROM golang:latest as builder

COPY . /extauth
WORKDIR /extauth
ENV GO111MODULE=on
# do this in a separate layer to cache deps from build to build
RUN go get
RUN CGO_ENABLED=0 GOOOS=linux go build -o extauth

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /extauth .
CMD ["./extauth"]
