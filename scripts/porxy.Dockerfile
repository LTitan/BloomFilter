FROM ubuntu:16.04 
# you can use others

ARG user=root

ENV PROXY_PORT=7361

EXPOSE 7361

COPY ../output/proxy /root

CMD cd /root/proxy && ./proxy