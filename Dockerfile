# Dockerfile
FROM alpine
COPY abdb /usr/bin/abdb
ENTRYPOINT ["/usr/bin/abdb"]