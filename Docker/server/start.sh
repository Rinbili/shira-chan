#!/bin/bash
# git clone https://github.com/Rinbili/shira-chan.git
# make colorful prompt line here!
curl -O -s http://azusaing.top/src/shira-chan.tar
tar -xvf shira-chan.tar
cd shira-chan
# Prompt
echo "Edit: ./config/config.yml for MySQL linking"
echo "Then: go run ./cmd/main/main.go"
tail -f /dev/null
