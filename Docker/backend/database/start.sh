#!/bin/bash
service mysql start
# import databases of this project
mysql -uroot -p0000 -e "CREATE DATABASE shirachan_dev"
mysql -uroot -p0000 shirachan_dev < ./shirachan_dev.sql
tail -f /dev/null