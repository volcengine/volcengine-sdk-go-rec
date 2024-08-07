package core

const (
	MaxWriteItemCount = 2000

	MaxImportItemCount = 10000

	// All requests will have a XXXResponse corresponding to them,
	// and all XXXResponses will contain a 'Status' field.
	// The status of this request can be determined by the value of `Status.Code`

	// StatusCodeSuccess The request was executed successfully without any exception
	StatusCodeSuccess = 0

	// StatusCodeIdempotent A Request with the same "Request-ID" was already received. This Request was rejected
	StatusCodeIdempotent = 409

	// StatusCodeOperationLoss Operation information is missing due to an unknown exception
	StatusCodeOperationLoss = 410

	// StatusCodeTooManyRequest The server hope slow down request frequency, and this request was rejected
	StatusCodeTooManyRequest = 429
)

const (
	volcAuthService = "air"
)

const (
	// metric name for  byteplus sdk
	metricsKeyInvokeSuccess = "invoke.success"
	metricsKeyInvokeError   = "invoke.error"
	metricsKeyPingSuccess   = "ping.success"
	metricsKeyPingError     = "ping.error"
)
