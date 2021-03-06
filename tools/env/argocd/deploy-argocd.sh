#!/bin/bash
set -x

version=v1.7.8

# wget https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
wget https://raw.githubusercontent.com/argoproj/argo-cd/v1.8.7/manifests/install.yaml

kubectl create namespace argocd
kubectl apply -n argocd -f install.yaml

# kubectl patch svc -n argocd argocd-server -p '{"spec": {"type": "NodePort","ports": [{"port":80,"nodePort":31140}]}}'
kubectl patch svc -n argocd argocd-server -p '{"spec": {"type": "NodePort","ports": [{"port":80,"nodePort":31140},{"port":443,"nodePort":31141}]}}'
