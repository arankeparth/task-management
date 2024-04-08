go build -o authservice server/taskservice/cmd/main.go
docker build server/taskservice/build/ -t partharanke/taskservice:latest
docker push partharanke/taskservice:latest