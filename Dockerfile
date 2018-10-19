FROM plugins/base:amd64

LABEL maintainer="Bo-Yi Wu <appleboy.tw@gmail.com>" \
  org.label-schema.name="Drone Minio" \
  org.label-schema.vendor="Bo-Yi Wu" \
  org.label-schema.schema-version="1.0"

RUN apk add --no-cache ca-certificates mc && \
  rm -rf /var/cache/apk/*

ADD release/linux/amd64/drone-minio /bin/
ENTRYPOINT ["/bin/drone-minio"]
