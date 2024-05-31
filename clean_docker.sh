sudo docker rm -f $(sudo docker ps -aq)
sudo docker network prune
sudo docker volume prune
cd fixtures
doucker-compose down -v
docker volume prune
docker container rm -f $(docker container ls -aq)
docker-compose up -d
cd ..
rm education
go build
./education