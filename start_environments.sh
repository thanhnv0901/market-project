#/bin/bash

docker-compose --file docker_environments/docker-compose.yaml up  --build && docker-compose --file docker_environments/docker-compose.yaml rm -fsv