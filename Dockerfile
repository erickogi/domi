FROM debian:bullseye-slim

LABEL maintainer="DevOps Kung Fu Masters"

ADD bin/domi domi
EXPOSE 8080

CMD ["./domi"]
