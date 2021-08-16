#!/bin/bash

mkdir cat /etc/CdnPurge
mkdir /var/log/cdn_tools

cp ./etc/config.ini /etc/CdnPurge/
cp ./preheater/cdnPushCache /usr/local/bin
chmod +x /usr/local/bin/cdnPushCache