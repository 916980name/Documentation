### show dns server
```
resolvectl status | grep "DNS Server" -A2
```
### change dns server
```
sudo resolvectl dns wlan0 192.168.0.1
```