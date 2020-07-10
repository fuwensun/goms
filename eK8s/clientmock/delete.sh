#!/bin/bash
# set -xe

PWD=$(cd "$(dirname "$0")";pwd)
echo $PWD

kubectl delete -f $PWD/clientmock-deploy.yaml --namespace="$1"

kubectl delete -f $PWD/clientmock-svc.yaml --namespace="$1"