### create secret for pull image
```
kubectl create secret generic regcred     --from-file=.dockerconfigjson=/home/user/.docker/config.json --type=kubernetes.io/dockerconfigjson
```
### load image to minikube from local machine
load
```
minikube image load 916980name/api-gateway:v1
```
show current exists
```
minikube image ls --format table
```
### expose minikube to local network
```
kubectl port-forward --address 0.0.0.0 ingress-nginx-controller-648cf4cd7d-dqb4j -n ingress-nginx 18080:80
```
### get the docker container IP
```
docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}'  nginx
```