package client

import (
	"github.com/04Akaps/gateway_module/common"
	"github.com/go-resty/resty/v2"
	"time"
)

func (h *HttpClient) fetchProducer() {
	h.fetchLock.Lock()
	defer h.fetchLock.Unlock()

	entity := h.mapper
	h.mapper = make([]ApiRequestTopic, 0)

	v, err := common.JsonHandler.Marshal(entity)

	if err == nil {
		h.producer.SendEvent(v)
	}
}

func (h *HttpClient) loopByBatchTime() {
	ticker := time.NewTicker(time.Duration(h.batchTime) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			h.fetchProducer()
		}
	}
}

func (h *HttpClient) handleRequest(resp *resty.Response, req interface{}) {
	if len(h.cfg.Producer.URL) > 0 {
		h.mapper = append(h.mapper, NewApiRequestTopic(resp, req))
	}
}
