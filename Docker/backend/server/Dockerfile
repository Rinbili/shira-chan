FROM debian:latest

RUN apt update
RUN apt install -y git
RUN apt install -y golang
RUN apt install -y curl
RUN apt install -y vim
COPY ./start.sh ./
CMD [ "./start.sh" ]
# last building checkpoint