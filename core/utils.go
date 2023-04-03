package core

import (
	"github.com/volcengine/volcengine-sdk-go-rec/core/metrics"
	"log"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

func AsyncExecute(runnable func()) {
	go func(run func()) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("[VolcengineSDK] async execute occur panic, "+
					"please feedback to bytedance, err:%v trace:\n%s", r, string(debug.Stack()))
			}
		}()
		run()
	}(runnable)
}

func IsNetError(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), netErrMark)
}

func IsTimeoutError(err error) bool {
	return strings.Contains(strings.ToLower(err.Error()), "timeout")
}

func ReportRequestSuccess(metricsPrefix, url string, begin time.Time) {
	urlTag := buildUrlTags(url)
	tagKvs := appendBaseTags(urlTag)
	metrics.Latency(buildLatencyKey(metricsPrefix), begin, tagKvs...)
	metrics.Counter(buildCounterKey(metricsPrefix), 1, tagKvs...)
}

func ReportRequestError(metricsPrefix, url string, begin time.Time, code int, message string) {
	urlTag := buildUrlTags(url)
	tagKvs := append(urlTag, "code:"+strconv.Itoa(code), "message:"+message)
	tagKvs = appendBaseTags(tagKvs)
	metrics.Latency(buildLatencyKey(metricsPrefix), begin, tagKvs...)
	metrics.Counter(buildCounterKey(metricsPrefix), 1, tagKvs...)
}

func ReportRequestException(metricsPrefix, url string, begin time.Time, err error) {
	urlTag := buildUrlTags(url)
	tagKvs := appendErrorTags(urlTag, err)
	tagKvs = appendBaseTags(tagKvs)
	metrics.Latency(buildLatencyKey(metricsPrefix), begin, tagKvs...)
	metrics.Counter(buildCounterKey(metricsPrefix), 1, tagKvs...)
}

func appendErrorTags(urlTag []string, err error) []string {
	msg := strings.ToLower(err.Error())
	if strings.Contains(msg, "time") && strings.Contains(msg, "out") {
		if strings.Contains(msg, "connect") {
			return append(urlTag, "message:connect-timeout")
		}
		if strings.Contains(msg, "read") {
			return append(urlTag, "message:read-timeout")

		}
		return append(urlTag, "message:timeout")
	}
	return append(urlTag, "message:other")
}

func buildUrlTags(url string) []string {
	switch {
	case strings.Contains(url, "ping"):
		return []string{"url:" + adjustUrlTag(url), "req_type:ping"}
	case strings.Contains(url, "data/api"):
		tenant, scene := parseTenantAndScene(url)
		return []string{"url:" + adjustUrlTag(url), "req_type:data-api", "tenant:" + tenant, "scene:" + scene}
	case strings.Contains(url, "predict/api"):
		tenant, scene := parseTenantAndScene(url)
		return []string{"url:" + adjustUrlTag(url), "req_type:predict-api", "tenant:" + tenant, "scene:" + scene}
	default:
		return []string{"url:" + adjustUrlTag(url), "req_type:unknown"}
	}
}

func adjustUrlTag(url string) string {
	return strings.ReplaceAll(url, "=", "_is_")
}

func parseTenantAndScene(url string) (tenant, scene string) {
	sp := strings.Split(strings.Split(url, "?")[0], "/")
	if len(sp) < 2 {
		return "", ""
	}
	return sp[len(sp)-2], sp[len(sp)-1]
}

func buildCounterKey(metricsPrefix string) string {
	return metricsPrefix + "." + "count"
}

func buildLatencyKey(metricsPrefix string) string {
	return metricsPrefix + "." + "latency"
}

const version = "1.1.0"

func appendBaseTags(tags []string) []string {
	return append(tags, "language:go", "version:"+version)
}
