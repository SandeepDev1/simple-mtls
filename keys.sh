#!/bin/bash

# Prompt for domain name
# shellcheck disable=SC2162
read -p "Enter the domain name (e.g., yourdomain.com): " DOMAIN

rm -rf server
rm -rf ca

# Configuration files
CA_CONFIG="[req]
default_bits = 2048
prompt = no
default_md = sha256
distinguished_name = dn

[dn]
C=US
ST=California
L=San Francisco
O=Your Company
OU=Your Department
emailAddress=ca@$DOMAIN
CN = $DOMAIN"

SERVER_CLIENT_CONFIG="[req]
default_bits = 2048
prompt = no
default_md = sha256
distinguished_name = dn
req_extensions = req_ext

[dn]
C=US
ST=California
L=San Francisco
O=Your Company
OU=Your Department
emailAddress=server@$DOMAIN
CN = $DOMAIN

[req_ext]
subjectAltName = @alt_names

[alt_names]
DNS.1 = $DOMAIN"

# CA certificate
echo "$CA_CONFIG" > ca.cnf
openssl genrsa -out ca.key 2048
openssl req -x509 -new -nodes -key ca.key -sha256 -days 1024 -out ca.crt -config ca.cnf

# Server certificate
echo "$SERVER_CLIENT_CONFIG" > server.cnf
openssl genrsa -out server.key 2048
openssl req -new -key server.key -out server.csr -config server.cnf
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 3650 -extensions req_ext -extfile server.cnf

# Remove unnecessary files
rm ca.cnf
rm server.cnf
rm server.csr
rm ca.srl

# Move the server certs to a /server directory
mkdir server
mv server.crt server/
mv server.key server/

# Move the ca certs to a /ca directory
mkdir ca
mv ca.crt ca/
mv ca.key ca/