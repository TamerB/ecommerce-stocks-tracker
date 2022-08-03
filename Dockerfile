FROM alpine AS alpine
 	 
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN apk add --no-cache bash

FROM scratch

COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY bin/ecommerce-stocks-tracker /usr/bin/ecommerce-stocks-tracker

ENTRYPOINT ["ecommerce-stocks-tracker"]