FROM alpine
COPY jaeger-opentracing-examples .
RUN mkdir -p /app/jaeger && mkdir /logs && mv jaeger-opentracing-examples /app/jaeger/
VOLUME /logs
WORKDIR /app/jaeger
ENTRYPOINT ["/app/jaeger/jaeger-opentracing-examples"]
