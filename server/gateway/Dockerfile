FROM scratch

COPY gateway /gateway

#RUN chmod +x /gateway/main.go

#ADD server.crt /etc/ssl/certs/

EXPOSE 443

ENTRYPOINT ["/gateway"]