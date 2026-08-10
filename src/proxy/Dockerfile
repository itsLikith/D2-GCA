FROM nginx:1.27-alpine

WORKDIR /etc/nginx

COPY nginx.conf /etc/nginx/nginx.conf

RUN rm -f /etc/nginx/conf.d/default.conf

CMD ["nginx", "-g", "daemon off;"]