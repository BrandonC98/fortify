#!/bin/bash

echo "=== Deploying Secret objects ==="
kubectl create secret generic fortify-secrets --from-env-file=secrets/secret-env
echo "=== Deploying MySQL database objects ==="
kubectl create -f mysql/storageclass.yaml
kubectl create -f mysql/statefulset.yaml
kubectl create -f mysql/service.yaml
echo "=== Deploying Generator objects ==="
kubectl create -f generator/configmap.yaml
kubectl create -f generator/deployment.yaml
kubectl create -f generator/service.yaml
echo "=== Deploying Fortify objects ==="
kubectl create -f fortify/configmap.yaml
kubectl create -f fortify/deployment.yaml
kubectl create -f fortify/service.yaml
echo "=== Deploying Ingress objects ==="
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.8.2/deploy/static/provider/aws/deploy.yaml
echo "Waiting for 30 seconds"
sleep 30
kubectl create -f ingress/ingress.yaml
