package core

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
	"github.com/volcengine/volcengine-sdk-go-rec/core/logs"
	"github.com/volcengine/volcengine-sdk-go-rec/core/option"
	"google.golang.org/protobuf/proto"
)

const netErrMark = "[netErr]"

func NewHTTPCaller(context *Context) *HTTPCaller {
	return &HTTPCaller{context: context}
}

type HTTPCaller struct {
	context *Context
}

func (c *HTTPCaller) DoJSONRequest(url string, request interface{},
	response proto.Message, options *option.Options) error {
	reqBytes, err := c.jsonMarshal(request)
	if err != nil {
		logs.Error("json marshal request fail, err:%s url:%s", err.Error(), url)
		return err
	}
	headers := c.buildHeaders(options, "application/json")
	url = c.withOptionQueries(options, url)
	rspBytes, err := c.doHttpRequest(url, headers, reqBytes, options.Timeout)
	if err != nil {
		return err
	}
	err = proto.Unmarshal(rspBytes, response)
	if err != nil {
		logs.Error("unmarshal response fail, err:%s url:%s", err.Error(), url)
		return err
	}
	return nil
}

func (c *HTTPCaller) jsonMarshal(request interface{}) ([]byte, error) {
	reqBytes, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	reqBytes = fasthttp.AppendGzipBytes(nil, reqBytes)
	return reqBytes, nil
}

func (c *HTTPCaller) DoPBRequest(url string, request proto.Message,
	response proto.Message, options *option.Options) error {
	reqBytes, err := c.marshal(request)
	if err != nil {
		logs.Error("marshal request fail, err:%s url:%s", err.Error(), url)
		return err
	}
	headers := c.buildHeaders(options, "application/x-protobuf")
	url = c.withOptionQueries(options, url)
	rspBytes, err := c.doHttpRequest(url, headers, reqBytes, options.Timeout)
	if err != nil {
		return err
	}
	err = proto.Unmarshal(rspBytes, response)
	if err != nil {
		logs.Error("unmarshal response fail, err:%s url:%s", err.Error(), url)
		return err
	}
	return nil
}

func (c *HTTPCaller) marshal(request proto.Message) ([]byte, error) {
	reqBytes, err := proto.Marshal(request)
	if err != nil {
		return nil, err
	}
	reqBytes = fasthttp.AppendGzipBytes(nil, reqBytes)
	return reqBytes, nil
}

func (c *HTTPCaller) buildHeaders(options *option.Options, contentType string) map[string]string {
	headers := make(map[string]string)
	headers["Content-Encoding"] = "gzip"
	headers["Accept-Encoding"] = "gzip"
	headers["Content-Type"] = contentType
	headers["Accept"] = "application/x-protobuf"
	headers["Tenant-Id"] = c.context.tenantId
	c.withOptionHeaders(headers, options)
	return headers
}

func (c *HTTPCaller) withOptionHeaders(headers map[string]string, options *option.Options) {
	if len(options.RequestId) == 0 {
		requestId := uuid.NewString()
		logs.Info("use requestId generated by sdk: '%s' ", requestId)
		headers["Request-Id"] = requestId
	} else {
		headers["Request-Id"] = options.RequestId
	}
	if !options.DataDate.IsZero() {
		headers["Content-Date"] = options.DataDate.Format("2006-01-02")
	}
	if options.DataIsEnd {
		headers["Content-End"] = "true"
	}
	if options.ServerTimeout > 0 {
		headers["Timeout-Millis"] = strconv.Itoa(int(options.ServerTimeout.Milliseconds()))
	}
	for k, v := range options.Headers {
		headers[k] = v
	}
}

func (c *HTTPCaller) withAuthHeaders(req *fasthttp.Request, reqBytes []byte) {
	if c.context.UseVolcAuth() {
		c.withVolcAuthHeaders(req)
		return
	}
	c.withAirAuthHeaders(req, reqBytes)
}

func (c *HTTPCaller) withAirAuthHeaders(req *fasthttp.Request, reqBytes []byte) {
	var (
		// Gets the second-level timestamp of the current time.
		// The server only supports the second-level timestamp.
		// The 'ts' must be the current time.
		// When current time exceeds a certain time, such as 5 seconds, of 'ts',
		// the signature will be invalid and cannot pass authentication
		ts = strconv.FormatInt(time.Now().Unix(), 10)
		// Use sub string of UUID as "nonce",  too long will be wasted.
		// You can also use 'ts' as' nonce'
		nonce = uuid.NewString()[:8]
		// calculate the authentication signature
		signature = c.calSignature(reqBytes, ts, nonce)
	)
	req.Header.Set("Tenant-Ts", ts)
	req.Header.Set("Tenant-Nonce", nonce)
	req.Header.Set("Tenant-Signature", signature)
}

func (c *HTTPCaller) withVolcAuthHeaders(req *fasthttp.Request) {
	VolcSign(req, c.context.volcCredentials)
}

func (c *HTTPCaller) calSignature(reqBytes []byte, ts, nonce string) string {
	var (
		token    = c.context.token
		tenantId = c.context.tenantId
	)
	// Splice in the order of "token", "HttpBody", "tenant_id", "ts", and "nonce".
	// The order must not be mistaken.
	// String need to be encoded as byte arrays by UTF-8
	shaHash := sha256.New()
	shaHash.Write([]byte(token))
	shaHash.Write(reqBytes)
	shaHash.Write([]byte(tenantId))
	shaHash.Write([]byte(ts))
	shaHash.Write([]byte(nonce))
	return fmt.Sprintf("%x", shaHash.Sum(nil))
}

func (c *HTTPCaller) withOptionQueries(options *option.Options, url string) string {
	var queriesParts []string
	if options.Stage != "" {
		queriesParts = append(queriesParts, "stage="+options.Stage)
	}
	for name, value := range options.Queries {
		queriesParts = append(queriesParts, name+"="+value)
	}
	optionQuery := strings.Join(queriesParts, "&")
	if optionQuery == "" {
		return url
	}
	if strings.Contains(url, "?") {
		url = url + "&" + optionQuery
	} else {
		url = url + "?" + optionQuery
	}
	return url
}

func (c *HTTPCaller) doHttpRequest(url string, headers map[string]string,
	reqBytes []byte, timeout time.Duration) ([]byte, error) {
	request := c.acquireRequest(url, headers, reqBytes)
	response := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(request)
		fasthttp.ReleaseResponse(response)
	}()
	c.withAuthHeaders(request, reqBytes)
	start := time.Now()
	defer func() {
		logs.Debug("http url:%s, cost:%s", url, time.Now().Sub(start))
	}()
	logs.Trace("http request header:\n%s", string(request.Header.Header()))
	err := c.smartDoRequest(timeout, request, response)
	if err != nil {
		ReportRequestException(metricsKeyInvokeError, url, start, err)
		if strings.Contains(strings.ToLower(err.Error()), "timeout") {
			logs.Error("do http request timeout, msg:%s url:%s", err.Error(), url)
			return nil, errors.New(netErrMark + " timeout")
		}
		logs.Error("do http request occur error, msg:%s url:%s", err.Error(), url)
		return nil, err
	}
	logs.Trace("http response headers:\n%s", string(response.Header.Header()))
	if response.StatusCode() != fasthttp.StatusOK {
		c.logHttpResponse(url, response)
		ReportRequestError(metricsKeyInvokeError, url, start, response.StatusCode(), "invoke-fail")
		return nil, errors.New(netErrMark + "http status not 200")
	}
	ReportRequestSuccess(metricsKeyInvokeSuccess, url, start)
	return decompressResponse(url, response)
}

func (c *HTTPCaller) acquireRequest(url string,
	headers map[string]string, reqBytes []byte) *fasthttp.Request {
	request := fasthttp.AcquireRequest()
	request.Header.SetMethod(fasthttp.MethodPost)
	request.SetRequestURI(url)
	for k, v := range headers {
		request.Header.Set(k, v)
	}
	request.SetBodyRaw(reqBytes)
	if len(c.context.hostHeader) > 0 {
		request.SetHost(c.context.hostHeader)
	}
	return request
}

func (c *HTTPCaller) smartDoRequest(timeout time.Duration,
	request *fasthttp.Request, response *fasthttp.Response) error {
	var err error
	var httpCli = c.context.defaultHTTPCli
	if timeout > 0 {
		err = httpCli.DoTimeout(request, response, timeout)
	} else {
		err = httpCli.Do(request, response)
	}
	return err
}

func (c *HTTPCaller) logHttpResponse(url string, response *fasthttp.Response) {
	rspBytes, _ := decompressResponse(url, response)
	if len(rspBytes) > 0 {
		logs.Error("http status not 200, url:%s code:%d headers:\n%s\n body:\n%s",
			url, response.StatusCode(), string(response.Header.Header()), string(rspBytes))
		return
	}
	logs.Error("http status not 200, url:%s code:%d headers:\n%s\n",
		url, response.StatusCode(), string(response.Header.Header()))
}

func decompressResponse(url string, response *fasthttp.Response) ([]byte, error) {
	contentEncoding := strings.ToLower(strings.TrimSpace(string(response.Header.Peek("Content-Encoding"))))
	switch contentEncoding {
	case "gzip":
		respBodyBytes, err := response.BodyGunzip()
		if err != nil {
			logs.Error("decompress gzip resp occur error, msg:%v url:%s header:\n%s",
				err, url, &response.Header)
			return nil, err
		}
		return respBodyBytes, nil
	case "":
		return response.Body(), nil
	default:
		logs.Error("receive unsupported response content encoding:%s url:%s header:\n%s",
			contentEncoding, url, &response.Header)
		err := errors.New("unsupported resp content encoding:" + contentEncoding)
		return nil, err
	}
}
