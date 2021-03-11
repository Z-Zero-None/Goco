package tracer

import (
	"io"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
)

func NewJaegerTracer(serviceName, agentHostPort string) (opentracing.Tracer, io.Closer, error) {
	//为jaeger client初始化
	cfg := &config.Configuration{
		ServiceName: serviceName,
		Sampler: &config.SamplerConfig{ //固定采样，对所有信息都进行采样
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{ //是否开启刷新缓冲区的频率，上报Agent地址
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  agentHostPort,
		},
	}
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		return nil, nil, err
	}
	//设置全局Tracer对象
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer, nil
}
