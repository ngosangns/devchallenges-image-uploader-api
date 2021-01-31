:: Docker build
echo y | docker-compose pull
:: Start containers
docker-compose up --remove-orphans
:: Remove obsolete images
echo y | docker image prune