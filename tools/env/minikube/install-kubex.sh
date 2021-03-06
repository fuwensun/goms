#!/bin/bash
# set -x
# set -e

# version
# version=1.18.1
version=1.20.5

# version
kubeadm_version="${version}-00"
kubelet_version="${version}-00"
kubectl_version="${version}-00"

# apt gpg
curl -s https://mirrors.aliyun.com/kubernetes/apt/doc/apt-key.gpg | apt-key add -

# apt config
cat <<EOF > /etc/apt/sources.list.d/kubernetes.list
deb https://mirrors.aliyun.com/kubernetes/apt/ kubernetes-xenial main
EOF
sudo apt update

# apt install
sudo apt remove kubeadm kubelet kubectl -y
sudo apt install kubeadm=$kubeadm_version -y --allow-downgrades
sudo apt install kubelet=$kubelet_version -y --allow-downgrades
sudo apt install kubectl=$kubectl_version -y --allow-downgrades

