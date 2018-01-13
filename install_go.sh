#!/bin/bash
if [ "$1" = "" ]; then
	version=1.92
else
	version=$1
fi
wget https://storage.googleapis.com/golang/go$version.linux-amd64.tar.gz
tar -xf go$version.linux-amd64.tar.gz
rm -rf /usr/local/go
cp -R go /usr/local/
rm -rf go go$version.linux-amd64.tar.gz
entry=$(cat /etc/profile|grep /usr/local/go)
if [ "$entry" = "" ]; then
	echo "adding an entry"
	echo "export PATH=$PATH:/usr/local/go/bin" >> /etc/profile
else
	echo "go root set up."
fi
