#!/bin/bash
HTTP_CODE=$(curl -s -o /dev/null -w "%{http_code}" http://localhost:8080)

# Check if the HTTP status code is 200 (OK)
if [ $HTTP_CODE -eq 200 ]; then
    echo "Service is running. HTTP status code: $HTTP_CODE"
    exit 0
else
    echo "Service might be down or error occurred. HTTP status code: $HTTP_CODE"
    exit 1
fi
