# Remove obsolete images
docker image prune --filter="dangling=true" -f