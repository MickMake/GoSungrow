#!/bin/bash

mkdir -p /usr/local/bin

wget -O /usr/local/bin/GoSungro --header "PRIVATE-TOKEN: ${GO_REPO_TOKEN}" https://mickmake.io/api/v4/projects/docker%2Fsungro/repository/files/dist%2FGoSungro_linux_amd64%2FGoSungro/raw?ref=master
chmod a+x /usr/local/bin/GoSungro

mkdir -p /root/.ssh
chmod 500 /root/.ssh

wget -O /root/.ssh/gopbx_rsa --header "PRIVATE-TOKEN: ${GO_REPO_TOKEN}" https://mickmake.io/api/v4/projects/docker%2Fsungro/repository/files/.ssh%2Fgopbx_rsa/raw?ref=master
wget -O /root/.ssh/gopbx_rsa.pub --header "PRIVATE-TOKEN: ${GO_REPO_TOKEN}" https://mickmake.io/api/v4/projects/docker%2Fsungro/repository/files/.ssh%2Fgopbx_rsa.pub/raw?ref=master

chmod 400 /root/.ssh/gopbx_rsa /root/.ssh/gopbx_rsa.pub

echo '00 07  *  *  *    /usr/local/bin/GoSungro sync default' > /etc/crontabs/root

apk add --no-cache colordiff tzdata

