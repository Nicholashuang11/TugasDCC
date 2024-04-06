FROM golang:alpine3.10
WORKDIR /app
COPY . .
EXPOSE 8080
CMD ["go","run","assignment1.go"]