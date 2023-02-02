FROM golang:1.19-alpine3.16 as builder

WORKDIR /src

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . ./

RUN go build -o /app .

FROM alpine:3.16

RUN apk --no-cache add ca-certificates

COPY --from=builder /app /

ENTRYPOINT ["/app"]
CMD ["-h"]
