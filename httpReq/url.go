package httpReq

const (
	// env
	local = "http://localhost:8000"
	uat   = "https://itms-uat.hk-tess.com"
	prod  = "https://itms-prod.hk-tess.com"

	// endpoint
	loginUrl           = "/api/v1/account/login"
	listRephrasedJob   = "/api/v1/order/job/rephrased"
	getRephrasedJobUrl = "/api/v1/order/job/rephrased/"
)
