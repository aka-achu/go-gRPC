#!/bin/bash

SERVER_CN=localhost
PASSPHRASE=ECE1A42DABDADBAE4BEBC9D8EAFEF
KEY_SIZE=4096
EXPIRY=3650

openssl genrsa -passout pass:${PASSPHRASE} -des3 -out ca.key ${KEY_SIZE}
openssl req -passin pass:${PASSPHRASE} -new -x509 -days ${EXPIRY} -key ca.key -out ca.crt -subj "/CN=${SERVER_CN}"
openssl genrsa -passout pass:${PASSPHRASE} -des3 -out server.key ${KEY_SIZE}
openssl req -passin pass:${PASSPHRASE} -new -key server.key -out server.csr -subj "/CN=${SERVER_CN}"
openssl x509 -req -passin pass:${PASSPHRASE} -days ${EXPIRY} -in server.csr -CA ca.crt -CAkey ca.key -set_serial 01 -out server.crt
openssl pkcs8 -topk8 -nocrypt -passin pass:${PASSPHRASE} -in server.key -out server.pem