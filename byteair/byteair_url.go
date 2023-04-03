package byteair

import (
	"fmt"
	"github.com/volcengine/volcengine-sdk-go-rec/common"
)

const (
	// The URL template of "predict" request, which need fill with "scene" info when use
	// Example: https://api.byteair.volces.com/predict/api/20013144/home
	predictUrlFormat = "%s://%s/predict/api/%s/{}"

	// The URL format of reporting the real exposure list
	// Example: https://api.byteair.volces.com/predict/api/20013144/callback
	callbackUrlFormat = "%s://%s/predict/api/%s/callback"

	// The URL format of data uploading
	// Example: https://api.byteair.volces.com/data/api/20013144/user?method=write
	uploadUrlFormat = "%s://%s/data/api/%s/{}?method=%s"

	// The URL format of marking a whole day data has been imported completely
	// Example: https://api.byteair.volces.com/data/api/20013144/done?topic=user
	doneUrlFormat = "%s://%s/data/api/%s/done?topic={}"
)

type byteairURL struct {
	cu     *common.URL
	schema string
	tenant string

	// The URL template of "predict" request, which need fill with "scene" info when use
	// Example: https://api.byteair.volces.com/predict/api/20013144/home
	predictUrlFormat []string

	// The URL of reporting the real exposure list
	// Example: https://api.byteair.volces.com/predict/api/20013144/callback
	callbackURL []string

	// The URL of uploading real-time user data
	// Example: https://api.byteair.volces.com/data/api/20013144/user?method=write
	writeDataURLFormat []string

	// The URL of importing daily offline user data
	// Example: https://api.byteair.volces.com/data/api/20013144/user?method=import
	importDataURLFormat []string

	// The URL format of marking a whole day data has been imported completely
	// Example: https://api.byteair.volces.com/data/api/20013144/done?topic=user
	doneURLFormat []string
}

func (receiver *byteairURL) Refresh(hosts []string) {
	receiver.cu.Refresh(hosts)
	receiver.predictUrlFormat = receiver.generatePredictURLFormat(hosts)
	receiver.callbackURL = receiver.generateCallbackURL(hosts)
	receiver.writeDataURLFormat = receiver.generateUploadURL(hosts, "write")
	receiver.importDataURLFormat = receiver.generateUploadURL(hosts, "import")
	receiver.doneURLFormat = receiver.generateDoneURL(hosts)
}

func (receiver *byteairURL) generatePredictURLFormat(hosts []string) []string {
	format := make([]string, 0, len(hosts))
	for _, host := range hosts {
		format = append(format, fmt.Sprintf(predictUrlFormat, receiver.schema, host, receiver.tenant))
	}
	return format
}

func (receiver *byteairURL) generateCallbackURL(hosts []string) []string {
	format := make([]string, 0, len(hosts))
	for _, host := range hosts {
		format = append(format, fmt.Sprintf(callbackUrlFormat, receiver.schema, host, receiver.tenant))
	}
	return format
}

func (receiver *byteairURL) generateUploadURL(hosts []string, method string) []string {
	format := make([]string, 0, len(hosts))
	for _, host := range hosts {
		format = append(format, fmt.Sprintf(uploadUrlFormat, receiver.schema, host, receiver.tenant, method))
	}
	return format
}

func (receiver *byteairURL) generateDoneURL(hosts []string) []string {
	format := make([]string, 0, len(hosts))
	for _, host := range hosts {
		format = append(format, fmt.Sprintf(doneUrlFormat, receiver.schema, host, receiver.tenant))
	}
	return format
}
