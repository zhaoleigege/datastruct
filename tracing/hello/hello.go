package main

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"time"
)

func main() {
	cfg := config.Configuration{
		ServiceName: "hello",
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			//CollectorEndpoint:   "http://127.0.0.1:14268/api/traces", // 通过http传递span的地址和端口
			BufferFlushInterval: 1 * time.Second,
			LogSpans:            true,
			LocalAgentHostPort:  "localhost:6831", // 通过udp传递span的地址和端口
		},
	}
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		panic(err)
	}
	defer closer.Close()

	// 注册opentracing
	opentracing.SetGlobalTracer(tracer)

	ctx := context.Background()
	doHello(ctx)
}

func doHello(ctx context.Context) {
	parent, ctx := opentracing.StartSpanFromContext(ctx, "doHello")
	defer parent.Finish()

	time.Sleep(1 * time.Second)
	fun1(ctx)
	fun2(ctx)
}

func fun1(ctx context.Context) {
	child1, _ := opentracing.StartSpanFromContext(ctx, "fn1")
	defer child1.Finish()

	time.Sleep(2 * time.Second)
}
func fun2(ctx context.Context) {
	child1, _ := opentracing.StartSpanFromContext(ctx, "fn2")
	defer child1.Finish()

	time.Sleep(2 * time.Second)
}
