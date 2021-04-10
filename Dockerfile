#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
RUN apk update && apk upgrade && \
    apk add --no-cache git
WORKDIR /go/src/app
COPY . .
COPY /assets /assets
COPY /mysql /mysql
RUN GOOS=linux go build -o /go/bin/app .
RUN go get -d -v ./...
RUN go install -v ./...

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /app
COPY --from=builder /assets /assets
COPY --from=builder /mysql /mysql
ENTRYPOINT ./app
LABEL Name=razerassignment Version=0.0.1
EXPOSE 5000
