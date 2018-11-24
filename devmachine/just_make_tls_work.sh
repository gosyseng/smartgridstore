#!/bin/bash

cat >devmachine-localhost.ca.crt << EOF
-----BEGIN CERTIFICATE-----
MIIFXjCCA0agAwIBAgIJAK4m2J2WItmsMA0GCSqGSIb3DQEBCwUAMEQxCzAJBgNV
BAYTAlVTMQswCQYDVQQIDAJDQTETMBEGA1UEBwwKRGV2TWFjaGluZTETMBEGA1UE
CgwKRGV2TWFjaGluZTAeFw0xODEwMzAyMDQwMzNaFw0yODEwMjcyMDQwMzNaMEQx
CzAJBgNVBAYTAlVTMQswCQYDVQQIDAJDQTETMBEGA1UEBwwKRGV2TWFjaGluZTET
MBEGA1UECgwKRGV2TWFjaGluZTCCAiIwDQYJKoZIhvcNAQEBBQADggIPADCCAgoC
ggIBALY74G2LAcbxYq1oOs7+U4YDfgRt9Nvmr+6F4DylXhACOS8b0oRRrDum2Mba
F+Mu5S/uUPVMbyul69vcJ0CtquttJR9J2cOfj4pAA7a22th0idAlXX6uJheMQXHV
wMzam0dzdVk5LNzIRae16ruzqOFS8GSwMKrSg7zdurSeLrK3G8fEdB1YQZg3Fd1z
zT+KqY7QA9WnLYwFWwPI/ir3OlJrFxHBIRsqtipcrytyYhiOnE40seEJX9vaXshz
PopMXVzrPCypO5YW2TJFxyqi4XK3Ab2VZoD065Z6cTrCu/M7hj0EDyuu9N9x8vnn
wrK6t9B66YlBz2935TAtN5NubazOWJxb6Za7JsWNygYG8K3XS6l+Rb/b6VKan/VS
wkgRvkJuOgTo8pF7GPkDq+Ddn0JgTJ2gNRS+2UbSDgz3jg/ORiP3p1vdqcwGrsG6
8VmGc0iCvv3ItkZmjCztRXL+O0BAiw6gfUmB/gdqJIFR56S/Ilx1iji6iJOG4K9x
9Dev7/AT8NbQlQ1ilNKNVhb01I14vN6JSNaK6ZC5pgEnCaENz7s4WJ0UPMpE8ZrV
Xj6VfE9Pi/xrrwSg8Vu8Rs0U8btzgOWQLKyoPcUkJePHaJ8d5JukE0k8Mdeim8Zd
GLovHJ3+aupk1kdiUAVk/+CZ4Nbx1qRG07PMkA5AhYhdLdDBAgMBAAGjUzBRMB0G
A1UdDgQWBBTfxswkk1Y5rX2g5To8NEIeUsmkezAfBgNVHSMEGDAWgBTfxswkk1Y5
rX2g5To8NEIeUsmkezAPBgNVHRMBAf8EBTADAQH/MA0GCSqGSIb3DQEBCwUAA4IC
AQBzk5iyzYrKzb8SW5jLg7V7Bc/XneD/sOqHTCAatagF80WPmCIiVD2jwxoosDiZ
/fWvzzlNCAZPXa6IOr0crAHmETaJdki7SFjHYchk86x7qzqtdBH/ADrBmetIkRJN
QOScKyQpIfSqrzbEi4cS+PdLt3FSFRxyb3q5ujEE07woIlJMQ0UZ7+lZBM9rrria
A3uXp/ZKk+GFbe5qivGJYMEqaDNPA9AtFC92xcIwGkFsXVjT6rdmw1gKdqJd7Clp
8V4PNoVT2pKMqccSUHxM+PLoW4Sg22JJbSvAgdFmUHGqgpq2vX1+1Npexe/Lb5hG
tWyGqwz14u7wfG4a9pyshvNAjg3tQ9qvRdAGsgCBSgMCUCVPQDyn0aoumJ377e6R
/0TqLgro4rrWuLTR265b1IZ292NZxGfi77Wye3U/s9iQ1uq6qXXQ3HDB1Sc4oY6h
WZKn3UpQLeP6ypez82vCxJMF1UIXuRe1gy5ysJ91IpOFnyhcWBUv2ZAxptq3APeU
4HcGvWR67hkJvflvgLpaqqm/ATim1cwgoYIkn4woeu67k+i7qyYuDRGBPVhdLBXc
aQGz1W2arIgMmldirkPrPnhHqnb0s2BTWG75c6OmBmFMVp7uSO0BxMGk9vufBYuQ
AOFKy9svQ5cJtl8ZISPMSXWCwgA7OsYFxEnoDVC28AnQBg==
-----END CERTIFICATE-----
EOF

docker exec -i devmachine-etcd /bin/bash << EOSCRIPT

export ETCDCTL_API=3

etcdctl --endpoints 172.29.0.20:2379 put api/hardcoded_pub << EOF
-----BEGIN CERTIFICATE-----
MIIEZDCCAkygAwIBAgIJAPSp+uG0b0OvMA0GCSqGSIb3DQEBCwUAMEQxCzAJBgNV
BAYTAlVTMQswCQYDVQQIDAJDQTETMBEGA1UEBwwKRGV2TWFjaGluZTETMBEGA1UE
CgwKRGV2TWFjaGluZTAeFw0xODExMjQxOTM1MTlaFw0yMTA4MjAxOTM1MTlaMEMx
CzAJBgNVBAYTAlVTMQswCQYDVQQIDAJDQTETMBEGA1UECgwKRGV2TWFjaGluZTES
MBAGA1UEAwwJbG9jYWxob3N0MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKC
AQEAngm1FYj8qOxnstFVDIIPdoJXgITVhzfhfF9txy0Wa1YbFTMhbyayxq/wL1bO
eagug7zIZDLcrwE7dUff2WA0Jl+Z5Gr6WEY15g0KilYOZfrdhsqtVGueLVqbrxrt
KqsQwbTg4I4e+CgdLMsGKxT3O8ijpC/XA8sZk5DRozNV7TWZutZ0ZeP0bQP7uLXt
ejYMQF9IyvDQY8TQgJoqX6WQA3xziXizolyVL62HQOUKcanpJtzeG4jEQre+73GB
HUauB5pUuMpBuw4vT8hdBt9bWMhAyAI+57H6cWXUTtcJDZ2pxIio0Q2aCwQEyw3i
rx6DBmlqIUiiXF2bl452WLpIxwIDAQABo1owWDBWBgNVHREETzBNgglsb2NhbGhv
c3SCFmRldm1hY2hpbmUtYXBpZnJvbnRlbmSCFGRldm1hY2hpbmUtbXJwbG90dGVy
ghJkZXZtYWNoaW5lLWNvbnNvbGUwDQYJKoZIhvcNAQELBQADggIBAKXb2mluGve1
kgiVb7kJ7qt6FiTeFh2F09I9Sh/yzbzuu4fzgwX/JuTI/9pmmTvnuJ49J7WE+KNz
zPVvgFRlRTVVNyScTD2WoqY5Et8e6W89y565IfSJjFangc6YjLqxgMgl5B4K4pDd
M/LT4Sm23sLH2bi5RcDpsFyJ+Qy26VMq4ylVQrEmVqTVd6sdJZzif4T3xVByb3qI
NUeZGUTO4I9Y5glj7UwE2kGuLCuAcG+XdKoY+scvIU4MdQdVuSuDvU65SA+jLPQJ
NyYx1cA18ov5b0jdvH90l4mMk8Fy2obMRRPTncZOCJ6oLuACub6jrE233OFXH7B2
QmZuKcPb7kbdE1jJ8CVsKZKRojCsJwzSrBYwauL9vn1ljJ6XdBxFUyhNmp+6I9XG
O+SA62kPso46cwlyhnqsUrEYWvXziYq9jyu6Yhd1Dp3dOrWPY96N/XtmKFFFy3PL
HSh+ITWekP4zvbzbyq1MgNjFhjdoBoi9YAbUcbd1JMZKBUPQgdfiptU4eTDWNnEL
K20pLlbMmb9CrfiOKp8kXmcUa/dNmwQahSzBA0fJ1XriXuJeYoDydksXVaDU3Chn
u0WtaYLwwPCzyFYkzYC0cILFgzwovSPdSRrTjfG4NfG9KpAfJCyZ+709RDn3iL03
o56SAD7Vg65KPskihelIBpK3q444V+UX
-----END CERTIFICATE-----
EOF

etcdctl --endpoints 172.29.0.20:2379 put api/hardcoded_priv << EOF
-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAngm1FYj8qOxnstFVDIIPdoJXgITVhzfhfF9txy0Wa1YbFTMh
byayxq/wL1bOeagug7zIZDLcrwE7dUff2WA0Jl+Z5Gr6WEY15g0KilYOZfrdhsqt
VGueLVqbrxrtKqsQwbTg4I4e+CgdLMsGKxT3O8ijpC/XA8sZk5DRozNV7TWZutZ0
ZeP0bQP7uLXtejYMQF9IyvDQY8TQgJoqX6WQA3xziXizolyVL62HQOUKcanpJtze
G4jEQre+73GBHUauB5pUuMpBuw4vT8hdBt9bWMhAyAI+57H6cWXUTtcJDZ2pxIio
0Q2aCwQEyw3irx6DBmlqIUiiXF2bl452WLpIxwIDAQABAoIBADg8DAuVqxIQzPV8
N180CHyK9MfufGyZHcs7ybxN9uRKx2hLwWHjy0mX4Ud3BBGWHrcRvdypnI1JSvb4
lFatPOFKs1/sERjIIvWxPkbrSERRd/I2vGaFxMsuqqcsVagRJu1Gh55f8/UpkPgX
ZiJ+W1sSkegifcgDxR4VMgN4iOLW2KTEz+Pwqb9dypSqxyBOvL0mDSoYYHZPjY2o
KCZRJLw7F+stTg/l9KdgGeZPO0k9z8B5LDCdNGJ43TgBhd9p4gihjaQLLgE3Vve6
OIaWlZDRAn+ozX5TP1/7tMB1sQdsjT3KiU+7k0vEzf8rDwU61EktXID64OaApA3K
tRqxL4ECgYEAzhd/0SbUr2H0DNfnSbVqjujuzD7ailA6XJv2khEhGT1vdR/mrc+X
zYbxkDCZC0D/gWWY2qkqEzzYlU+/F6b4bYtc/8N3FTmfKZ1JOF/3Q4GrRzWm+xax
LrW6n+2l32Jxn31Hei4t8WC0gq0wCDcKQ+xZmiGXfdKp31b3lmVaHNkCgYEAxE8l
OxTmMc9ZPidx28PaMDEH3ZanEEGPJxRwa0Veqsr4fF4Yo8Hc4YIKYDQdnAjIV3Bf
L91ovHF7wH3mQAADPIF4MPpSXGhvBS45aPfjTh7hC03oP/Y7OPKcfIX2ZM3jXEcr
Uw6qOuU83uaHBIhOR6oSNYnuEf4MfrTg9uaSjp8CgYAfyHn6zG1Ceu3DSUyvDl5X
mc6yCwNzDzpg6+CT2DZtiQ72ViwiJAC6PpRkc0o29RgyYXO4TRxjCSAxrrG+uQ5Y
/oQiIYfs6w0DwvD2073zWD0M51ZQJMjAdmBIp8LL94ekXBK44gRN1GciPAlLSwm3
Pez3mzScv+9YtTXLqOfDyQKBgCwzNcrLtdjZOtrHvtcgjevz3jWCSjNaz4SZEYbV
o68I7FH0Tc/xQm5z/SoUEeakA5fMcWIxk/d6Bccdkc4vMotZC1rbwdeUnAqqkbqU
tdVNFpaZAAYGxubXmWxgLU4dHWgVjiexUZrZKo91kEjw988i996eI8BhVybgSxbM
a4s5AoGBAMb6mAHTb3F5TLJd5TD6TXgXyWoaY/GYYSu44emDc5Ed4SY9G9wkdXbJ
RPT8dNPed67hm9vdQSt6EgUPBRkPkV5iSD8A2mY5oaXDZBV2xTKoAIy1zBTlFoGD
8oCPQVE6F52M7AezCBMMd51LhLpm3hzGef4wbIqbLowRHp5U21lN
-----END RSA PRIVATE KEY-----
EOF

etcdctl --endpoints 172.29.0.20:2379 put api/certsrc hardcoded
EOSCRIPT

set -x
docker restart devmachine-console
docker restart devmachine-apifrontend
set +x
echo "We have created a 'devmachine-localhost.ca.crt' file. Add this to your browser's Certificate Authorities and everything will just work"
echo "For non-browser use, copy the pem file to /usr/local/share/ca-certificates/btrdb/ and ensure it is world-readable. Then run sudo update-ca-certificates"
echo "Ensure that the API endpoint is given as localhost:4411. Using an IP address will not work"
echo "For python you also need to run export BTRDB_CA_BUNDLE=/etc/ssl/certs/ca-certificates.crt to make grpc use system ssl certificates"
