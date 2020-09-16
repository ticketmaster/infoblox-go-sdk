package test

var (
	// Username - infoblox user.
	Username = ""
	// Password - infoblox password.
	Password = ""
	// InfobloxServer - IP or DNS name of infoblox server.
	InfobloxServer = ""
	// TArgs - arguments for test cases.
	TArgs = Args{}
)

// Args - fields for test arguments.
type Args struct {
	// CIDR - ex. 1.1.1.0/24.
	CIDR string
	// Cname - alias or cname.
	Cname string
	// Host - hostname.
	Host string
	// HostRegex - ex. foo.*bar.
	HostRegex string
	// FilterAndRegex - ex. name~=foo.*bar.
	FilterAndRegex string
	// IP - ex. 1.1.1.1.
	IP string
	// ARef - infoblox A reference. ex. record:a/ZG5zLmJpbmRfYSQuX2RlZmF1bHQudG1jcy5uZXRvcHMsZG9ub3R1c2UsMTAuODQuMjQuMjAw:donotuse.netops.tmcs/default
	ARef string
	// CnameRef - infoblox CNAME reference. ex. record:cname/ZG5zLmJpbmRfYSQuX2RlZmF1bHQudG1jcy5uZXRvcHMsZG9ub3R1c2UsMTAuODQuMjQuMjAw:donotuse.netops.tmcs/default
	CnameRef string
	// HostRef - infoblox HOST reference. ex. record:cname/ZG5zLmJpbmRfYSQuX2RlZmF1bHQudG1jcy5uZXRvcHMsZG9ub3R1c2UsMTAuODQuMjQuMjAw:donotuse.netops.tmcs/default
	HostRef string
	// StartIP - first ip in subnet. ex. 1.1.1.1.
	StartIP string
	// Comment - comment field.
	Comment string
}
