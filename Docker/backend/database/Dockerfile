FROM ubuntu:latest

# so that you can install mysql automatically
COPY ./mysql.preseed ./
RUN debconf-set-selections ./mysql.preseed

# MySQL Service
RUN apt update
RUN apt install -y mysql-server

# .sql file of database
COPY ./shirachan_dev.sql ./

# runtime script 
COPY ./start.sh ./
CMD ["./start.sh"]