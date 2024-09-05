# musical-barnacle

How to expose the actual HOST/VPS IP to the client?

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
