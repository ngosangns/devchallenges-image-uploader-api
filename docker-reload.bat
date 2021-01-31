:: Golang build
set GOOS=linux
go build

:: Docker build
echo y | docker-compose pull
:: Start containers
docker-compose up -d --remove-orphans
:: Remove obsolete images
echo y | docker image prune