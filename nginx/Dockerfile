FROM nginx:1.13.1-alpine

LABEL maintainer="GPF <5173180@qq.com>"

#https://yeasy.gitbooks.io/docker_practice/content/image/build.html
RUN mkdir -p /etc/nginx/cert \
    && mkdir -p /etc/nginx/conf.d \
    && mkdir -p /etc/nginx/sites

COPY ./nginx.conf /etc/ngixn/nginx.conf
COPY ./cert/ /etc/nginx/cert/

VOLUME ["/var/log/nginx", "/var/www", "/etc/nginx/sites/", "/etc/nginx/conf.d/"]

WORKDIR /usr/share/nginx/html