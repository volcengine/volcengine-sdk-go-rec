package common

import (
	"fmt"

	"github.com/volcengine/volcengine-sdk-go-rec/core"
)

const (
	// The URL format of operation information
	// Example: https://tob.sgsnssdk.com/data/api/retail_demo/operation?method=get
	operationUrlFormat = "%s://%s/data/api/%s/operation?method=%s"

	// The URL of mark certain days that data synchronization is complete
	// Example: https://tob.sgsnssdk.com/data/api/retail_demo/done?topic=user
	doneUrlFormat = "%s://%s/data/api/%s/done?topic={}"
)

func NewURL(context *core.Context) *URL {
	return &URL{
		schema: context.Schema(),
		tenant: context.Tenant(),
	}
}

type URL struct {
	schema            string
	tenant            string
	getOperationUrl   []string
	listOperationsUrl []string
	doneUrlFormat     []string
}

func (receiver *URL) Refresh(hosts []string) {
	receiver.getOperationUrl = receiver.generateOperationUrl(hosts, "get")
	receiver.listOperationsUrl = receiver.generateOperationUrl(hosts, "list")
	receiver.doneUrlFormat = receiver.generateDoneUrl(hosts)
}

func (receiver *URL) generateOperationUrl(hosts []string, method string) []string {
	format := make([]string, 0, len(hosts))
	for _, host := range hosts {
		format = append(format, fmt.Sprintf(operationUrlFormat, receiver.schema, host, receiver.tenant, method))
	}
	return format
}

func (receiver *URL) generateDoneUrl(hosts []string) []string {
	format := make([]string, 0, len(hosts))
	for _, host := range hosts {
		format = append(format, fmt.Sprintf(doneUrlFormat, receiver.schema, host, receiver.tenant))
	}
	return format
}
