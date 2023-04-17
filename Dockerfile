FROM alpine:3.17

LABEL maintainer="albekov <me@albekov.net>" \
    org.label-schema.name="Drone Pushover Plugin" \
    org.label-schema.vendor="albekov" \
    org.label-schema.schema-version="1.0.1"

LABEL org.opencontainers.image.source=https://github.com/albekov/drone-pushover
LABEL org.opencontainers.image.description="Drone plugin for sending Pushover notifications"
LABEL org.opencontainers.image.licenses=MIT

RUN apk update && \
    apk add --no-cache ca-certificates && \
    rm -rf /var/cache/apk/*

WORKDIR /app
COPY ./bin/drone-pushover /app/
ENTRYPOINT ["/app/drone-pushover"]
