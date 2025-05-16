FROM alpine
WORKDIR /
COPY build/trmnl-server .
EXPOSE 8080
ENTRYPOINT ["/trmnl-server"]
