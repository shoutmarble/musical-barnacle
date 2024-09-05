# musical-barnacle

How to expose the actual HOST/VPS IP to the client?


`openssl s_client -connect mail.landingdev.xyz:443 -verify_ip 185.28.22.166`
```
root@DESKTOP-3LUEUH3:/mnt/c/Users/Terry/Documents/GIT/musical-barnacle# openssl s_client -connect mail.landingdev.xyz:443 -verify_ip 185.28.22.166
CONNECTED(00000003)
depth=0 CN = landingdev.xyz
verify error:num=64:IP address mismatch
verify return:1
depth=2 C = US, O = Internet Security Research Group, CN = ISRG Root X1
verify return:1
depth=1 C = US, O = Let's Encrypt, CN = R10
verify return:1
depth=0 CN = landingdev.xyz
verify return:1
---
Certificate chain
 0 s:CN = landingdev.xyz
   i:C = US, O = Let's Encrypt, CN = R10
   a:PKEY: rsaEncryption, 2048 (bit); sigalg: RSA-SHA256
   v:NotBefore: Aug  7 00:39:59 2024 GMT; NotAfter: Nov  5 00:39:58 2024 GMT
 1 s:CN = landingdev.xyz
   i:C = US, O = Let's Encrypt, CN = R10
   a:PKEY: rsaEncryption, 2048 (bit); sigalg: RSA-SHA256
   v:NotBefore: Aug  7 00:39:59 2024 GMT; NotAfter: Nov  5 00:39:58 2024 GMT
 2 s:C = US, O = Let's Encrypt, CN = R10
   i:C = US, O = Internet Security Research Group, CN = ISRG Root X1
   a:PKEY: rsaEncryption, 2048 (bit); sigalg: RSA-SHA256
   v:NotBefore: Mar 13 00:00:00 2024 GMT; NotAfter: Mar 12 23:59:59 2027 GMT
---
---
    Start Time: 1725514679
    Timeout   : 7200 (sec)
    Verify return code: 64 (IP address mismatch)
    Extended master secret: no
    Max Early Data: 0
---
read R BLOCK

```


```
Hostname: who
IP: 127.0.0.1
IP: ::1
   ┌─────────────┐
IP:│192.168.10.20│◀──────────────────────────────────────┐
   └─────────────┘                                       │
Dnt: 1                                                   │
                                                 ┌───────┴──────┐
Forwarded: proto=https;for=66.51.236.224:60955;by│185.28.22.166 │
                                                 └──────────────┘
Sec-Ch-Ua: "Chromium";v="128", "Not;A=Brand";v="24", "Google Chrome";v="128"
X-Forwarded-For: 66.51.236.224
X-Forwarded-Port: 443
X-Forwarded-Proto: https
```
