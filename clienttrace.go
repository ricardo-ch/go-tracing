package tracing

import (
	"net/http/httptrace"

	opentracing "github.com/opentracing/opentracing-go"
)

// NewClientTrace Creates a New ClientTrace
func newClientTrace(span opentracing.Span) *httptrace.ClientTrace {
	trace := &clientTrace{span: span}
	return &httptrace.ClientTrace{
		DNSStart:             trace.dnsStart,
		DNSDone:              trace.dnsDone,
		GetConn:              trace.getConn,
		GotConn:              trace.gotConn,
		ConnectStart:         trace.connectStart,
		ConnectDone:          trace.connectDone,
		GotFirstResponseByte: trace.gotFirstResponseByte,
		WroteRequest:         trace.wroteRequest,
	}
}

// clientTrace holds a reference to the Span and
// provides methods used as ClientTrace callbacks
type clientTrace struct {
	span opentracing.Span
}

func (h *clientTrace) dnsStart(info httptrace.DNSStartInfo) {
	h.span.LogKV("event", "DNS Start",
		"host", info.Host,
	)
}

func (h *clientTrace) dnsDone(d httptrace.DNSDoneInfo) {
	h.span.LogKV("event", "DNS done")
}

func (h *clientTrace) getConn(hostPort string) {
	h.span.LogKV("event", "Get Connection")
}

func (h *clientTrace) gotConn(httptrace.GotConnInfo) {
	h.span.LogKV("event", "Got Connection")
}

func (h *clientTrace) connectStart(network, addr string) {
	h.span.LogKV("event", "Connection Start")
}

func (h *clientTrace) connectDone(network, addr string, err error) {
	if err != nil {
		h.span.LogKV("event", "Connection Done",
			"err", err.Error(),
		)
	} else {
		h.span.LogKV("event", "Connection Done")
	}
}

func (h *clientTrace) wroteRequest(httptrace.WroteRequestInfo) {
	h.span.LogKV("event", "Wrote Request")
}

func (h *clientTrace) gotFirstResponseByte() {
	h.span.LogKV("event", "Got First Response byte")
}
