package httputil

import (
	"time"

	"github.com/imroc/req/v3"
	"go.opentelemetry.io/otel/trace"
)

type Client struct {
	*req.Client
}

// NewClient 创建客户端
func NewClient() *Client {
	return &Client{
		Client: req.C(),
	}
}

// NewDefaultClient 创建默认客户端
func NewDefaultClient() *Client {
	client := req.C().
		SetCommonRetryCount(2).                                      // 设置重试次数
		SetCommonRetryBackoffInterval(1*time.Second, 5*time.Second). // 设置重试间隔
		SetTimeout(10 * time.Second).                                // 设置超时时间
		EnableTraceAll().                                            // 开启链路追踪
		EnableDumpEachRequest().                                     // 开启请求和响应的转储
		OnAfterResponse(ResponseMiddleware)                          // 设置响应中间件
	return &Client{
		client,
	}
}

// SetDebug 设置调试模式
func (c *Client) SetDebug(enable bool) *Client {
	if enable {
		c.EnableDebugLog()
		c.EnableDumpAll()
	} else {
		c.DisableDebugLog()
		c.DisableDumpAll()
	}
	return c
}

// SetTracer 设置链路追踪。
func (c *Client) SetTracer(tracer trace.Tracer) *Client {
	c.WrapRoundTripFunc(Tracer(tracer))
	return c
}

// SetPrometheus 设置Prometheus
func (c *Client) SetPrometheus() *Client {
	c.OnAfterResponse(Prometheus())
	return c
}

// Impersonate 设置Http指纹
func (c *Client) SetImpersonate(browser string) *Client {
	switch browser {
	case "chrome":
		c.ImpersonateChrome()
	case "firefox":
		c.ImpersonateFirefox()
	case "safari":
		c.ImpersonateSafari()
	default:
		c.ImpersonateChrome() // 默认
	}
	return c
}

// SetTLSFingerprintChrome 设置TLS指纹
func (c *Client) SetTLSFingerprint() *Client {
	// 模拟 Chrome 浏览器的 TLS 握手指纹，让网站相信这是 Chrome 浏览器在访问，予以通行。
	c.SetUserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36").SetTLSFingerprintChrome()
	return c
}
