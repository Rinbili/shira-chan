FROM debian:latest

RUN apt update
RUN apt install -y git
RUN apt install -y golang
RUN apt install -y curl
COPY ./start.sh ./
CMD [ "./start.sh" ]
# last building checkpoint