#!/bin/bash
# Obtain the web-release
echo -e "\033[32m=> Download and Extract latest files\033[0m"
curl -O -s http://azusaing.top/src/dist.tar
tar -xf ./dist.tar
mv ./dist /var/www/html/

# modify nginx config
rm /etc/nginx/sites-enabled/default
mv ./nginx_conf/* /etc/nginx/sites-enabled/
echo -e "\033[32m=> Modify necessary config file and rm the other one\033[0m"

tail -f /dev/null