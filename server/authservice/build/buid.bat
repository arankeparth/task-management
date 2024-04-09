set GOOS=linux&& set GOARCH=amd64&& go build -o authservice server/authservice/cmd/main.go
docker build server/authservice/build/ -t partharanke/authentication:latest
docker push partharanke/authentication:latest