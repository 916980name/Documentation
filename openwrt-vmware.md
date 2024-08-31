### Download
openwrt-23.05.4-x86-64-generic-squashfs-combined-efi.img.gz

### Format
Convert openwrt.img to VBox drive
```
VBoxManage convertfromraw --format VDI openwrt-23.05.4-x86-64-generic-squashfs-combined-efi.img openwrt.vdi
```
When convert fail
```
Converting from raw image file="openwrt-23.05.4-sunxi-cortexa7-friendlyarm_nanopi-m1-plus-squashfs-sdcard.img" to file="openwrt-23.05.4-friendlyarm_nanopi.vdi"...
Creating dynamic image with size 27277434 bytes (27MB)...
VBoxManage: error: VD: The given disk size 27277434 is not aligned on a sector boundary (512 bytes)
VBoxManage: error: Error code VERR_VD_INVALID_SIZE at /home/vbox/tinderbox/build-VBox-7.0/svn/src/VBox/Storage/VD.cpp(6052) in function int VDCreateBase(PVDISK, const char*, const char*, uint64_t, unsigned int, const char*, PCVDGEOMETRY, PCVDGEOMETRY, PCRTUUID, unsigned int, PVDINTERFACE, PVDINTERFACE)
VBoxManage: error: Cannot create the disk image "openwrt-23.05.4-friendlyarm_nanopi.vdi": VERR_VD_INVALID_SIZE
```
[solution](https://stackoverflow.com/a/59179239/8936864)
```
truncate openwrt-23.05.4-sunxi-cortexa7-friendlyarm_nanopi-m1-plus-squashfs-sdcard.img --size=128M
```

### Settings
1. virtualbox -> file -> Host network manager -> create 192.168.56.1
1. virtualbox -> preferences -> network -> nat network -> create
1. vm -> settings -> network -> host-only, nat, bridged

### Bootup

### Config
iptables effect when start
```
vim /etc/firewall.user
```

show network
```
uci show network
```
config network
```
uci set network.lan.ipaddr='192.168.56.2'
uci commit
reboot
```
```
uci batch <<EOF 
set network.mng=interface 
set network.mng.device='br-lan'  
set network.mng.proto='static'
set network.mng.ipaddr='192.168.56.2'  
set network.mng.netmask='255.255.255.0'
set firewall.@zone[0].network='mng'
set firewall.@zone[0].name='mng'
delete network.lan
delete network.wan6
set network.wan=interface
set network.wan.device='eth1'
set network.wan.proto='dhcp'
EOF
```
```
uci changes
uci commit && reboot
```
```
uci batch <<EOF 
set network.lan=interface
set network.lan.device='eth2'
set network.lan.proto='dhcp'
EOF
uci commit
service network restart
```

### Proxy
1. sysctl -w net.ipv4.ip_forward=1
1.  "网关LAN_IP地址段" 通过运行命令"ip address | grep -w "inet" | awk '{print $2}'"获得，是其中的一个
```
iptables -t nat -N V2RAY # 新建一个名为 V2RAY 的链
iptables -t nat -A V2RAY -d 192.168.1.0/16 -j RETURN # 直连 192.168.0.0/16 
iptables -t nat -A V2RAY -p tcp -j RETURN -m mark --mark 0xff # 直连 SO_MARK 为 0xff 的流量(0xff 是 16 进制数，数值上等同与上面配置的 255)，此规则目的是避免代理本机(网关)流量出现回环问题
iptables -t nat -A V2RAY -p tcp -j REDIRECT --to-ports 1090 # 其余流量转发到 12345 端口（即 V2Ray）
iptables -t nat -A PREROUTING -p tcp -j V2RAY # 对局域网其他设备进行透明代理
iptables -t nat -A OUTPUT -p tcp -j V2RAY # 对本机进行透明代理

ip rule add fwmark 1 table 100
ip route add local 0.0.0.0/0 dev lo table 100
iptables -t mangle -N XRAY
iptables -t mangle -A XRAY ! -s 192.168.1.5 -j RETURN
iptables -t mangle -A XRAY -p tcp -j TPROXY --on-port 12345 --tproxy-mark 1
iptables -t mangle -A XRAY -p udp -j TPROXY --on-port 12345 --tproxy-mark 1
iptables -t mangle -A PREROUTING -j XRAY
```

### Problem
1. wget failed: Failed to allocate uclient context
```
opkg print-architecture
```

### Qemu
```
qemu-system-arm -m 1024 -machine orangepi-pc -drive format=raw,file=openwrt-23.05.4-sunxi-cortexa7-friendlyarm_nanopi-m1-plus-ext4-sdcard.img -nographic
```