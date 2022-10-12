#!/usr/bin/env bash

# This file assumes the OS is fedora based

### Add Arena root CA
cat << E0F > /etc/pki/ca-trust/source/anchors/arena-root-ca.pem
# Arena Solutions Primary CA (SHA) (Exp. 2038-08-24 UTC)
-----BEGIN CERTIFICATE-----
MIIEozCCA4ugAwIBAgIBADANBgkqhkiG9w0BAQUFADCBkDELMAkGA1UEBhMCVVMx
HjAcBgNVBAoTFUFyZW5hIFNvbHV0aW9ucywgSW5jLjE8MDoGA1UECxMzQXJlbmEg
VHJ1c3QgTmV0d29yayBQcmltYXJ5IENlcnRpZmljYXRpb24gQXV0aG9yaXR5MSMw
IQYDVQQDExpBcmVuYSBTb2x1dGlvbnMgUHJpbWFyeSBDQTAeFw0xMzA4MzAwMDA1
MTBaFw0zODA4MjQwMDA1MTBaMIGQMQswCQYDVQQGEwJVUzEeMBwGA1UEChMVQXJl
bmEgU29sdXRpb25zLCBJbmMuMTwwOgYDVQQLEzNBcmVuYSBUcnVzdCBOZXR3b3Jr
IFByaW1hcnkgQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkxIzAhBgNVBAMTGkFyZW5h
IFNvbHV0aW9ucyBQcmltYXJ5IENBMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIB
CgKCAQEA9ia5ahYRsfijF2oFUTcp2mhrkfCDJ8IcIoxEgKb3JOiG+WNFphqyJkCL
jskio57n5AKRakTsjBZN8H/0JTdESkWQfJjifcDN0xN6/HXR9tG9LbrERcA9RGUq
x9vv3Wilb9wZ1mg3qc0GrDbsizuMfdAOo7dY0RRODH7xMNbLca1PEN+vhJi7hheN
Py4WB8S+emFTgjFU3xAY062gJA48eqJ0G/UNlGPU7Jaacag/57eErfc009hJchR4
mYfaT8cEIcnmBL2sbpB6FyUTB15sPLATpoTQqqNc5zNTI0N+PH/aDbEw0uwT1H/4
kxBJwveKoo/rTs4HfEU0Bn9BQT1zIwIDAQABo4IBBDCCAQAwHQYDVR0OBBYEFKnH
8ZQHxaOkh+W+EIqFVAfgnQw9MA4GA1UdDwEB/wQEAwIBBjAPBgNVHRMBAf8EBTAD
AQH/MIG9BgNVHSMEgbUwgbKAFKnH8ZQHxaOkh+W+EIqFVAfgnQw9oYGWpIGTMIGQ
MQswCQYDVQQGEwJVUzEeMBwGA1UEChMVQXJlbmEgU29sdXRpb25zLCBJbmMuMTww
OgYDVQQLEzNBcmVuYSBUcnVzdCBOZXR3b3JrIFByaW1hcnkgQ2VydGlmaWNhdGlv
biBBdXRob3JpdHkxIzAhBgNVBAMTGkFyZW5hIFNvbHV0aW9ucyBQcmltYXJ5IENB
ggEAMA0GCSqGSIb3DQEBBQUAA4IBAQAG3YGFkiI6BaSFi4y9lc/JiMZnyHkPiO/E
mDcmt/Hzh6TUCkyS9eMWWtEnR1pD60l0AC89syNS9DuLPfNYs4kUYbM5hf8D+TEe
uXFucyLfYxTNkHPZxjYzapQJQDs1TaM8OAofEL5qO35ZBzsv67IuQqSzICkQFp8q
/nItM+Nh5lJC+BvRydPhBm8dYHT2eGMa/e2oZy+IX4JjhTBLydTHFKalj7AHwcA8
a/A1KePT0sfipAZbGK+4SaL8VKQUMsolLeHrdY4TZgEkfgsWE5YwVoq+g82PCNun
R+lVNegCWTZ/H/RY0DHhq9UDgpA3JiE2IdkO85gyIFoLpzc0NQIj
-----END CERTIFICATE-----
E0F

# Arena
cat << E0F > /etc/pki/ca-trust/source/anchors/arena-root-sha2-ca.pem
# Arena Primary CA (SHA2) (Exp. 2047-02-12 UTC)
-----BEGIN CERTIFICATE-----
MIIEVzCCAz+gAwIBAgIBADANBgkqhkiG9w0BAQwFADB5MQswCQYDVQQGEwJVUzER
MA8GA1UEChMIUFRDIEluYy4xPDA6BgNVBAsTM0FyZW5hIFRydXN0IE5ldHdvcmsg
UHJpbWFyeSBDZXJ0aWZpY2F0aW9uIEF1dGhvcml0eTEZMBcGA1UEAxMQQXJlbmEg
UHJpbWFyeSBDQTAeFw0yMjAyMTgwMDQ1MDZaFw00NzAyMTIwMDQ1MDZaMHkxCzAJ
BgNVBAYTAlVTMREwDwYDVQQKEwhQVEMgSW5jLjE8MDoGA1UECxMzQXJlbmEgVHJ1
c3QgTmV0d29yayBQcmltYXJ5IENlcnRpZmljYXRpb24gQXV0aG9yaXR5MRkwFwYD
VQQDExBBcmVuYSBQcmltYXJ5IENBMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIB
CgKCAQEAyMA5SKPD+2qNypjcnVuvwUE1L+OpjT/1MOrx9D7WKXlcPabu6Tg6JuZr
G2GBeokXSyUZ2vsgp6ZzD13EeEx4RyfN47j9jshPpr7gd6zd7G82nlMeIBPz1lTt
mu0Ie5Q4G9xowEt7pwv2PsGZcRzulgy9jeuR3YoDIi4kz2MWzUu2JJFjvoxUGylM
+iw2OyeCG9iwZrgKP/l4ycIcLgccp1zcR/gX4RGuT0Itz16WkXR5eqrZDA75cKMo
7QYbIf1VLmlPuZKuPwEJr1uZTv/k2/fNXl7rur2ViR4oWgFyI82oRDjgzidZfXGU
dekIjOZS4ecyGY74ZsX4FPSd9T2D/wIDAQABo4HpMIHmMB0GA1UdDgQWBBRJucU9
QYbwykyX0A2y2ha1z9nnRTAOBgNVHQ8BAf8EBAMCAQYwDwYDVR0TAQH/BAUwAwEB
/zCBowYDVR0jBIGbMIGYgBRJucU9QYbwykyX0A2y2ha1z9nnRaF9pHsweTELMAkG
A1UEBhMCVVMxETAPBgNVBAoTCFBUQyBJbmMuMTwwOgYDVQQLEzNBcmVuYSBUcnVz
dCBOZXR3b3JrIFByaW1hcnkgQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkxGTAXBgNV
BAMTEEFyZW5hIFByaW1hcnkgQ0GCAQAwDQYJKoZIhvcNAQEMBQADggEBAHK2lafy
FoKFKDC8u1Tyd3awDT+Z3KkHhTqDgCaQtPNgSz/moGPo4vJiHPgvfFmzpT7fvHtM
3aFR7Olb0k/eQtQ1LK4YPDGKU2qsOpYGHy6+MfdD9dbm111vozbG54V9FIpK2zDX
IdsRHRnJjol6Bdm3iJSl+NdbCBpwjy4J6gFbesdqJeEvRQpHmQNTQmQpAoevddLK
68YqHyXfSShZw+xhyc9RfT6NLEHe0266oddLxJaFB04kpWyYUTp+UDHAUtco8lBL
6MIYrr0U/MqxzCIyc9kguns9AVULL1FrCex+QVGm2d1hlWSoXVHCrso6vsR5/INu
eL1zGbT5zuCzbmw=
-----END CERTIFICATE-----
E0F

/usr/bin/update-ca-trust
