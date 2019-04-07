FROM golang:1.12.1 as builder
ENV GO111MODULE=on
WORKDIR $GOPATH/src/backend/
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go test -cover -tags=unit ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build 

#FROM ubuntu:18.04
FROM scratch
COPY --from=builder /go/src/backend/backend /backend
ENTRYPOINT ["/backend"]
