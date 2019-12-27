#!/bin/bash

TARGET_PATH=/opt/argononefan
SERVICE_NAME=adjustfan.service

systemctl stop $SERVICE_NAME || true
systemctl disable $SERVICE_NAME || true
rm /etc/systemd/system/$SERVICE_NAME
$TARGET_PATH/setfan 0
rm -f $TARGET_PATH/setfan
rm -f $TARGET_PATH/readtemp
rm -f $TARGET_PATH/adjustfan
rm -f $TARGET_PATH/adjustfan.json
rmdir $TARGET_PATH
