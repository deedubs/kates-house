 FROM golang:1.17-alpine3.14 as builder

 WORKDIR /app

 COPY go.* ./
 RUN go mod download

 COPY . ./

 RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly  -v -o event-ingress
#  RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -mod=readonly  -v -o event-ingress

 FROM scratch

 COPY --from=builder /app/event-ingress /event-ingress

 # Run the web service on container startup.
 CMD ["/event-ingress"]