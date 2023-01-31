FROM golang:1.19-alpine3.16 as builder

WORKDIR /src

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . ./

RUN go build -o /jobs-handler .

FROM alpine:3.16

RUN apk --no-cache add ca-certificates

COPY --from=builder /jobs-handler /

ENTRYPOINT ["/jobs-handler"]
CMD ["-h"]
