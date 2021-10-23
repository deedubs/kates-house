 FROM golang:1.17-alpine3.14 as builder

 WORKDIR /app

 COPY go.* ./
 RUN go mod download

 COPY . ./

 RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly  -v -o temperature-sensor

 FROM alpine:3
 RUN apk add --no-cache ca-certificates

 COPY --from=builder /app/temperature-sensor /temperature-sensor

 # Run the web service on container startup.
 CMD ["/temperature-sensor"]