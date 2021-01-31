:: Build
echo y | docker-compose build

:: Start containers with stopping obsolete containers
docker-compose up --remove-orphans -d

:: Remove obsolete images
docker image prune --filter="dangling=true" -f

pause