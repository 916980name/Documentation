#### Install cri
https://github.com/Mirantis/cri-dockerd/releases

```
sudo apt --fix-broken install ./cri-dockerd_0.3.15.3-0.ubuntu-jammy_amd64.deb

systemctl status containerd.service
```
> You need CRI support enabled to use containerd with Kubernetes.   
> Make sure that `cri` is <strong>not</strong> included in the disabled_plugins list within `/etc/containerd/config.toml`  
> if you made changes to that file, also restart containerd: `sudo systemctl restart containerd`

#### Install crio (maybe not need, containerd is ok)
```
KUBERNETES_VERSION=v1.31
CRIO_VERSION=v1.31
sudo apt-get update
sudo apt-get install -y software-properties-common curl
curl -fsSL https://pkgs.k8s.io/addons:/cri-o:/stable:/$CRIO_VERSION/deb/Release.key |
    sudo gpg --dearmor -o /etc/apt/keyrings/cri-o-apt-keyring.gpg
sudo echo "deb [signed-by=/etc/apt/keyrings/cri-o-apt-keyring.gpg] https://pkgs.k8s.io/addons:/cri-o:/stable:/$CRIO_VERSION/deb/ /" |   sudo  tee /etc/apt/sources.list.d/cri-o.list
sudo apt-get update
sudo apt-get install -y cri-o
sudo systemctl start crio.service
```

#### Install kubectl
```
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
```
#### Install kubelet
```
### create soft link to /usr/local/bin/kubelet
### get .service content from create-kubelet-conf()
sudo systemctl enable kubelet.service
```


#### All of machines need 
kubadmin, kubelet, kubectl

#### Bootstrap a cluster
```
swapoff -a
modprobe br_netfilter
sysctl -w net.ipv4.ip_forward=1
```