#!/bin/bash
docker_username_lower=$(echo "$docker_username" | tr '[:upper:]' '[:lower:]')
repository_name="MrChristmas96/simple-fortune-cookie"
echo "$docker_password" | docker login ghcr.io --username "$docker_username_lower" --password-stdin
docker push "ghcr.io/$repository_name/simple-fortune-cookie-backend:1.0-${GIT_COMMIT::8}"
docker push "ghcr.io/$repository_name/simple-fortune-cookie-frontend:1.0-${GIT_COMMIT::8}"
docker push "ghcr.io/$repository_name/simple-fortune-cookie-backend:latest"
docker push "ghcr.io/$repository_name/simple-fortune-cookie-frontend:latest"
wait