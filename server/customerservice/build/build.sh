go build -o authservice server/customerservice/cmd/main.go
docker build server/customerservice/build/ -t partharanke/customerservice:latest
docker push partharanke/customerservice:latest