FROM golang:1.23  AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN go mod tidy 
COPY . .
RUN go build -o main .

FROM ubuntu:22.04
RUN apt-get update && apt-get install -y ffmpeg
COPY --from=build /app/main /usr/local/bin/main
COPY . /app/
RUN chmod +x /app/transcode.sh
EXPOSE 9000

ENTRYPOINT ["/usr/local/bin/main"]