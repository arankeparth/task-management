go build -o server/customerservice/build/customerservice server/customerservice/cmd/main.go
docker build server/customerservice/build/ -t partharanke/customerservice:latest
docker push partharanke/customerservice:latest
