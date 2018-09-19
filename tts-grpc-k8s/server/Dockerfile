FROM alpine

RUN apk update && apk add flite
ADD say /say

ENTRYPOINT ["/say"]