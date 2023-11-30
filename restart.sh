#!/usr/bin/env bash
git pull
pkill brake_signal
go build
nohup /root/go_proj/brake_signal/brake_signal -c /root/go_proj/brake_signal/config/conf_prd.json >> /root/go_proj/logs/brake_signal.log &