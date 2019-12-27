#!/bin/bash

TARGET_PATH=/opt/argononefan
SERVICE_NAME=adjustfan.service

systemctl stop $SERVICE_NAME || true
systemctl disable $SERVICE_NAME || true
rm /etc/systemd/system/$SERVICE_NAME
$TARGET_PATH/setfan 0
rm  $TARGET_PATH/setfan
rm  $TARGET_PATH/readtemp
rm  $TARGET_PATH/adjustfan
rmdir $TARGET_PATH
