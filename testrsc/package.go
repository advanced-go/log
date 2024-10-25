package testrsc

const (
	PkgPath = "github/advanced-go/log/testrsc"
)
const (
	commonBasePath = "file:///f:/files/common"

	NotFoundResp            = commonBasePath + "/http-404-resp.txt"
	OKResp                  = commonBasePath + "/http-200-resp.txt"
	GatewayTimeoutResp      = commonBasePath + "/http-504-resp.txt"
	InternalServerErrorResp = commonBasePath + "/interanal-server-error-resp.txt"
)

const (
	log1BasePath = "file:///f:/files/log1"

	LOG1IngressEntry      = log1BasePath + "/ingress-entry.json"
	LOG1IngressGetAllReq  = log1BasePath + "/ingress-get-all-req.txt"
	LOG1IngressGetAllResp = log1BasePath + "/ingress-get-all-resp.txt"
	LOG1IngressPutReq     = log1BasePath + "/ingress-put-req.txt"

	LOG1EgressEntry          = log1BasePath + "/egress-entry.json"
	LOG1EgressGetAllReq      = log1BasePath + "/egress-get-all-req.txt"
	LOG1EgressGetAllResp     = log1BasePath + "/egress-get-all-resp.txt"
	LOG1EgressGetNotFoundReq = log1BasePath + "/egress-get-not-found-req.txt"

	log2BasePath = "file:///f:/files/log2"

	LOG2IngressEntry      = log2BasePath + "/ingress-entry.json"
	LOG2IngressGetAllReq  = log2BasePath + "/ingress-get-all-req.txt"
	LOG2IngressGetAllResp = log2BasePath + "/ingress-get-all-resp.txt"

	LOG2EgressEntry      = log2BasePath + "/egress-entry.json"
	LOG2EgressGetAllReq  = log2BasePath + "/egress-get-all-req.txt"
	LOG2EgressGetAllResp = log2BasePath + "/egress-get-all-resp.txt"
)

const (
	ts1BasePath = "file:///f:/files/timeseries1"

	TS1PercentileThreshold  = ts1BasePath + "/percentile-threshold.json"
	TS1PercentileGetAllReq  = ts1BasePath + "/percentile-get-all-req.txt"
	TS1PercentileGetAllResp = ts1BasePath + "/percentile-get-all-resp.txt"

	TS1StatusCodeThreshold  = ts1BasePath + "/stat-code-threshold.json"
	TS1StatusCodeGetAllReq  = ts1BasePath + "/stat-code-get-all-req.txt"
	TS1StatusCodeGetAllResp = ts1BasePath + "/stat-code-get-all-resp.txt"

	// test

)

/*
	ts1BasePath = "file:///f:/files/timeseries1"

	TS1EgressEntry      = ts1BasePath + "/egress-entry.json"
	TS1EgressEntryTest  = ts1BasePath + "/egress-percentile-threshold.json"
	TS1IngressEntry     = ts1BasePath + "/ingress-entry.json"
	TS1IngressEntryTest = ts1BasePath + "/ingress-percentile-threshold.json"

	TS1GetReq  = ts1BasePath + "/egress-get-all-req.txt"
	TS1GetResp = ts1BasePath + "/egress-get-all-resp.txt"

	ts2BasePath         = "file:///f:/files/timeseries2"
	TS2IngressEntry     = ts2BasePath + "/ingress-entry.json"
	TS2IngressEntryTest = ts2BasePath + "/ingress-percentile-threshold.json"
	TS2EgressEntry      = ts2BasePath + "/egress-entry.json"
	TS2EgressEntryTest  = ts2BasePath + "/percentile-threshold.json"


*/
