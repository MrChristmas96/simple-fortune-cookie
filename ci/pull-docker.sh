#!/bin/bash
docker_username_lower=$(echo "$docker_username" | tr '[:upper:]' '[:lower:]')
repository_name="mrchristmas96/simple-fortune-cookie"
echo "$docker_password" | docker login ghcr.io --username "$docker_username_lower" --password-stdin
docker pull "ghcr.io/$repository_name/simple-fortune-cookie-backend:latest"
docker pull "ghcr.io/$repository_name/simple-fortune-cookie-frontend:latest"
wait