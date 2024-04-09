echo "Building and deploying services"

chmod +x ./server/authservice/build/build.sh
chmod +x ./server/customerservice/build/build.sh
chmod +x ./server/taskservice/build/build.sh

./server/authservice/build/build.sh
./server/customerservice/build/build.sh
./server/taskservice/build/build.sh

docker build server/db/build -t partharanke/db:latest
docker push partharanke/db:latest