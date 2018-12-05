FROM alpine
COPY opentracing-jaeger-examples .
RUN mkdir -p /app/jaeger && mkdir /logs && mv opentracing-jaeger-examples /app/jaeger/
VOLUME /logs
WORKDIR /app/jaeger
ENTRYPOINT ["/app/jaeger/opentracing-jaeger-examples"]
