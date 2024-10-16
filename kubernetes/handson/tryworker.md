### Configure CNI Networking
```bash
POD_CIDR=192.168.56.0/24
cat <<EOF | sudo tee /etc/cni/net.d/10-bridge.conf
{
    "cniVersion": "1.3.0",
    "name": "bridge",
    "type": "bridge",
    "bridge": "cnio0",
    "isGateway": true,
    "ipMasq": true,
    "ipam": {
        "type": "host-local",
        "ranges": [
          [{"subnet": "${POD_CIDR}"}]
        ],
        "routes": [{"dst": "0.0.0.0/0"}]
    }
}
EOF
```

### Configure the Kubelet
```bash
WORKER_NAME=`hostname`
POD_CIDR=192.168.56.0/24

sudo ln -s /home/ubuntu/${WORKER_NAME}.pem /var/lib/kubelet/${WORKER_NAME}.pem
sudo ln -s /home/ubuntu/${WORKER_NAME}-key.pem /var/lib/kubelet/${WORKER_NAME}-key.pem
sudo ln -s /home/ubuntu/${WORKER_NAME}.kubeconfig /var/lib/kubelet/kubeconfig
sudo ln -s /home/ubuntu/ca.pem /var/lib/kubernetes/ca.pem

cat <<EOF | sudo tee /var/lib/kubelet/kubelet-config.yaml
kind: KubeletConfiguration
apiVersion: kubelet.config.k8s.io/v1beta1
authentication:
  anonymous:
    enabled: false
  webhook:
    enabled: true
  x509:
    clientCAFile: "/var/lib/kubernetes/ca.pem"
authorization:
  mode: Webhook
clusterDomain: "cluster.local"
clusterDNS:
  - "10.32.0.10"
podCIDR: "${POD_CIDR}"
resolvConf: "/run/systemd/resolve/resolv.conf"
runtimeRequestTimeout: "15m"
tlsCertFile: "/var/lib/kubelet/${WORKER_NAME}.pem"
tlsPrivateKeyFile: "/var/lib/kubelet/${WORKER_NAME}-key.pem"
EOF
```

### Configure the kubernetes proxy
1. ln
    ```bash
    sudo ln -s /home/ubuntu/kube-proxy.kubeconfig /var/lib/kube-proxy/kubeconfig
    ```
1. `kube-proxy-config.yaml`
```bash
cat <<EOF | sudo tee /var/lib/kube-proxy/kube-proxy-config.yaml
kind: KubeProxyConfiguration
apiVersion: kubeproxy.config.k8s.io/v1
clientConnection:
  kubeconfig: "/var/lib/kube-proxy/kubeconfig"
mode: "iptables"
clusterCIDR: "10.200.0.0/16"
EOF
```