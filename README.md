# Video Transcoder Service

A scalable video transcoding service built with Go, Gin, and FFmpeg. This service supports multiple resolutions (1080p, 720p, 480p, 360p, 240p, 144p) with concurrent processing and Docker integration.

## Project Running Example

Here’s a screenshot of the project running:

![Project Running Result](./images/running-result.png)

## Features

- Accepts video file uploads via a POST request.
- Transcodes videos into multiple resolutions: 1080p, 720p, 480p, 360p, 240p, and 144p.
- Utilizes concurrent transcoding with a semaphore to limit the number of simultaneous jobs.

## Project Structure

 ```
├── controller
│ └── transcode.go          # Handles transcoding requests
├── model
│ └── transcode.go          # Defines data models for transcoding results
├── service
│ └── transcode.go          # Contains the logic for transcoding videos
├── transcode.sh            # Shell script for FFmpeg transcoding
├── Dockerfile              # Dockerfile for building the application image
├── docker-compose.yml      # Docker Compose configuration file
├── main.go                 # Entry point of the application
└── .gitattributes          # Ensures consistent line endings for shell scripts
```
## Prerequisites

- Docker
- Docker Compose

## Installation

1. **Clone the repository:**

```bash
git clone https://github.com/skandergazzah/video-transcoder-service.git
cd video-transcoder-service
```
2. **Build and start the service using Docker Compose:**
```
docker-compose up --build
```
3. **Access the service:**
- The service will be available at http://localhost:9000.

## Usage
1. **Transcoding Endpoint**

- POST /transcode

3. **Request Example using curl:**
```
curl --location 'http://localhost:9000/transcode' --form 'video=@"/C:/Users/Skander/Downloads/surf.mp4"'
```
3. **Response Example:**
```
{
  "message": "Transcoding in progress",
  "resolutions_in_progress": [
    "1080p",
    "720p",
    "480p",
    "360p",
    "240p",
    "144p"
  ]
}
```
## Acknowledgments
- Gin - Web framework for Go.
- FFmpeg - Multimedia framework for handling video, audio, and other multimedia files and streams.

## License
[MIT](https://choosealicense.com/licenses/mit/)
