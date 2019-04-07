FROM nginx:1.14.2
COPY nginx/proxy-set-header.inc /etc/nginx/conf.d/
COPY nginx/nginx.conf /etc/nginx/nginx.conf
COPY swagger.json /usr/share/nginx/html/api/
