go build -o server/authservice/build/authservice server/authservice/cmd/main.go
docker build server/authservice/build/ -t partharanke/authentication:latest
docker push partharanke/authentication:latest
