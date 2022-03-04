osu

docker build --tag docker-gs-ping .
docker run --publish 1323:1323 docker-gs-ping
