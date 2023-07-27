FROM debian:latest

# Apache Service
RUN apt update
RUN apt install -y apache2
COPY ./dist /var/www/html
COPY ./start.sh ./
# Apache is named as 'apache2' in Debian and Ubuntu,
# instead of 'httpd' in Centos or other Red-Hat distributes.
CMD ["bash", "./start.sh"]