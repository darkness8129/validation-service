FROM golang:1.22.0-alpine3.19 as build

WORKDIR /app

# copy mod files and download dependencies
COPY go.mod go.sum .
RUN go mod download

# copy the rest of app source code
COPY . .

# run tests
RUN go test ./...

# build app
RUN go build -o /validation-service

# create new release stage
FROM alpine:latest as release
WORKDIR /

# copy binary from prev stage
COPY --from=build /validation-service /validation-service

EXPOSE 8080

CMD ["./validation-service"]