#!/bin/bash
docker_username_lower=$(echo "$docker_username" | tr '[:upper:]' '[:lower:]')
echo "$docker_password" | docker login ghcr.io --username "$docker_username_lower" --password-stdin
docker pull $docker_username_lower/micronaut-app:latest" &
wait