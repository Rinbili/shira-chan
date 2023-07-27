FROM debian:latest

# Nginx Service and basic binaries
RUN apt update && \
    apt install -y nginx && \
    apt install -y curl && \
    apt install -y vim 
COPY ./start.sh ./
COPY ./nginx_conf ./nginx_conf
# Apache is named as 'apache2' in Debian and Ubuntu,
# instead of 'httpd' in Centos or other Red-Hat distributes.
CMD ["bash", "./start.sh"]