#! /bin/sh

dst=/home/cloud/minikube/bin
rm $dst/*

cd /home/cloud/go/src/k8s.io/minikube/out
cp ./minikube $dst
cp ./minikube.iso $dst
ls -l $dst
