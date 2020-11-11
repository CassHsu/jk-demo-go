FROM golang:alpine AS build
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -v -o /app/jk-demo-go

FROM alpine
WORKDIR /app
COPY --from=build /app /app
EXPOSE 8080
CMD ["./jk-demo-go"]
