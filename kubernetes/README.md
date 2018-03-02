# gcloud instruction summary



## vm for management in google cloud
- create a vm with ssh key

To generate ssh key, useful when using external tool
```
ssh-keygen -t rsa -b 4096 -f ./yourKeyFileNameHere -C "yourNameHere"
```

### Connect to workstation

Install kubectl
```
sudo yum install kubectl
```

Set project, authorize and create cluster
```
gcloud config set project [PROJECT_ID]
gcloud config set compute/zone [COMPUTE_ZONE]
gcloud auth login
```

## create a cluster

[create cluster](https://cloud.google.com/kubernetes-engine/docs/quickstart)

```
gcloud container clusters list
gcloud container clusters delete [CLUSTER_NAME]
gcloud container clusters create cluster-k8sconfig 
```

### create deployment with three replicas from command line
```
kubectl run k8sconfig --image johncicilio/k8sconfig:latest --port 80 --replicas=3
```

[kubectl cheat sheet](https://kubernetes.io/docs/reference/kubectl/cheatsheet/)


# create deployment from configuration file
```
kubectl create -f filename
```


Some useful commands
```
kubectl get pods -o wide
kubectl logs [PODNAME]
```

Delete the deployment
``` 
kubectl delete deployment [DEPLOYMENT_NAME]
```

Create configmap, which holds a configuration file 
```
kubectl create configmap k8sconfig-config --from-file=./k8sconfig.yaml
```

Delete the cluster
```
gcloud container clusters delete cluster-k8sconfig 
```