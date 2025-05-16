FROM alpine
RUN apk add --no-cache tzdata
ENV TZ=Europe/Zurich
WORKDIR /
COPY build/trmnl-server .
EXPOSE 8080
ENTRYPOINT ["/trmnl-server"]
