#!/bin/bash

TARGET_PATH=/opt/argononefan
SERVICE_NAME=adjustfan.service

systemctl stop $SERVICE_NAME || true
mkdir $TARGET_PATH
cp setfan $TARGET_PATH
cp readtemp $TARGET_PATH
cp adjustfan $TARGET_PATH
chown -R root:root $TARGET_PATH
chmod -R 555 $TARGET_PATH
cp ./deploy/$SERVICE_NAME /etc/systemd/system
systemctl enable $SERVICE_NAME
systemctl start $SERVICE_NAME