package client

import (
	"github.com/go-resty/resty/v2"
	"net/http"
)

type CallErr struct {
	Url      string          `json:"url"`
	Message  string          `json:"message"`
	Response *resty.Response `json:"response"`
}

func (c *CallErr) Error() string {
	return c.Message
}

func NewCallError(url string, msg error, response *resty.Response) CallErr {
	return CallErr{
		Url:      url,
		Message:  msg.Error(),
		Response: response,
	}
}

type ApiRequestTopic struct {
	Trade   trace       `json:"trade"`
	Header  http.Header `json:"header"`
	Request interface{} `json:"request"`
	Code    int         `json:"code"`
}

type trace struct {
	ConnectionTime    float64 `json:"connectionTime"`
	TcpConnectionTime float64 `json:"tcpConnectionTime"`
	TlsHandshake      float64 `json:"tlsHandshake"`
	ServerTime        float64 `json:"serverTime"`
	TotalTime         float64 `json:"totalTime"`
	ResponseTime      float64 `json:"responseTime"`

	ConnectionRefused bool   `json:"connectionRefused"`
	RetryCount        int    `json:"retryCount"`
	RemoteAddr        string `json:"remoteAddr"`
}

func NewApiRequestTopic(resp *resty.Response, req interface{}) ApiRequestTopic {
	t := resp.Request.TraceInfo()

	connectionTime := t.ConnTime.Seconds()
	tCPConectionTime := t.TCPConnTime.Seconds()
	tLSHandShake := t.TLSHandshake.Seconds()
	serverTime := t.ServerTime.Seconds()
	totalTime := t.TotalTime.Seconds()

	reponseTime := t.ResponseTime.Seconds()
	connectionResued := t.IsConnReused
	retryCount := t.RequestAttempt

	remoteAddr := t.RemoteAddr

	header := resp.Header()

	return ApiRequestTopic{
		Trade: trace{
			ConnectionTime:    connectionTime,
			TcpConnectionTime: tCPConectionTime,
			TlsHandshake:      tLSHandShake,
			ServerTime:        serverTime,
			TotalTime:         totalTime,
			ResponseTime:      reponseTime,
			ConnectionRefused: connectionResued,
			RetryCount:        retryCount,
			RemoteAddr:        remoteAddr.String(),
		},
		Header:  header,
		Request: req,
		Code:    resp.StatusCode(),
	}
}
