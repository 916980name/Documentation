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
expose service
```
minikube service api-gateway-service --url -n test
```
### get the docker container IP
```
docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}'  nginx
```

### minikube problem
```
error execution phase certs/apiserver-kubelet-client: [certs] certificate apiserver-kubelet-client not signed by CA certificate ca: x509: certificate has expired or is not yet valid
```
https://github.com/kubernetes/minikube/issues/8770
```
minikube delete
minikube start
```