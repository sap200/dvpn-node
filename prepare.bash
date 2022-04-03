#!/usr/bin/bash

# Install go

apt update
apt install wget
wget https://go.dev/dl/go1.17.7.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.17.7.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
go version

# Install starport
curl https://get.starport.network/starport@v0.18.1 | bash
cp starport /usr/local/bin

