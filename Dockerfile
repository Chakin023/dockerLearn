FROM debian:stable-slim
COPY goserver /bin/goserver
ENV PORT=8080
CMD ["/bin/goserver"]