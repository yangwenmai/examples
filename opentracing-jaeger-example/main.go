package main

import (
	"fmt"
	"net/http"
	"time"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/openzipkin/zipkin-go-opentracing/thrift/gen-go/zipkincore"

	"github.com/uber/jaeger-client-go"
	jaegerClientConfig "github.com/uber/jaeger-client-go/config"
)

var t opentracing.Tracer
var cfg jaegerClientConfig.Configuration

func main() {
	cfg = jaegerClientConfig.Configuration{
		Sampler: &jaegerClientConfig.SamplerConfig{
			Type:  "const",
			Param: 1,
			// SamplingServerURL: "172.17.0.19:5775", // 刚开始没有搞懂上报的服务器URL，以为是这个，结果不是这个；
		},
		Reporter: &jaegerClientConfig.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			// LocalAgentHostPort:  "127.0.0.1:5775",
			LocalAgentHostPort: "172.17.0.9:5775",
		},
	}
	tracer, closer, _ := cfg.New(
		"j-examples",
		jaegerClientConfig.Logger(jaeger.StdLogger),
	)

	defer closer.Close()
	t = tracer

	testTracer2(testTracer2(testTracer1(t))) // working
	// testTracer()

	http.HandleFunc("/", HelloHandler)
	http.ListenAndServe(":8080", nil)
}

// HelloHandler world, the web server
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	testTracer1(t)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Hello, world!\n"))
}

// Working
func testTracer1(t opentracing.Tracer) (opentracing.Tracer, opentracing.Span) {
	span := t.StartSpan("new_span")
	defer span.Finish()
	span.SetOperationName("span_1")
	span.LogFields(log.String("ds", "asd"))
	span.LogEvent("hello")
	span.SetTag(zipkincore.HTTP_PATH, struct{ name string }{"ad"})

	span.SetBaggageItem("Some_Key", "12345")
	span.SetBaggageItem("Some-other-key", "42")

	return t, span
}

func testTracer2(t opentracing.Tracer, span opentracing.Span) (opentracing.Tracer, opentracing.Span) {
	span2 := t.StartSpan("span_2", opentracing.ChildOf(span.Context()))
	defer span2.Finish()
	span2.LogFields(log.String("hello", "span2"))
	time.Sleep(19 * time.Millisecond)
	return t, span2
}

func testTracer() interface{} {
	fmt.Println("testTracer...")
	tracer, closer, _ := cfg.New(
		"crossdock",
		jaegerClientConfig.Logger(jaeger.StdLogger),
	)
	defer closer.Close()

	span := tracer.StartSpan("hi")
	defer span.Finish()

	span.LogEvent("hello")
	span.SetBaggageItem("key", "xyz")
	// ctx := opentracing.ContextWithSpan(context.Background(), span)

	return nil
}
