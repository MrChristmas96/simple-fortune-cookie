#!/bin/bash

# Specify the directory containing your Kubernetes YAML files
K8S_DIR="../deployment/"

# Iterate through all YAML files in the directory
for file in "$K8S_DIR"/*.yaml; do
    if [ -f "$file" ]; then
        kubectl apply -f "$file"
        if [ $? -eq 0 ]; then
            echo "Applied $file successfully"
        else
            echo "Error applying $file"
        fi
    fi
done
