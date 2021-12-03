package ipfilter

import (
	"github.com/jpillora/ipfilter"
)

//IPFilter is an (experimental) struct represents IP filter
type IPFilter struct {
	policy    bool              // Policy specifies matching direction of the filtering rule. true for allow, false for deny
	RawFilter ipfilter.IPFilter // IPAddr is target IP to apply the filter with a policy
}

//NewIPFilter returns a new IPFilter instance with specified IP addresses.
//This
func NewIPFilter(policy bool, ips []string) *IPFilter {

	var (
		allowedIPs []string
		blockedIPs []string
	)

	if policy {
		allowedIPs = ips
	} else {
		blockedIPs = ips
	}

	return &IPFilter{
		policy: policy,
		RawFilter: *ipfilter.New(ipfilter.Options{
			BlockByDefault: !policy,
			AllowedIPs:     allowedIPs,
			BlockedIPs:     blockedIPs,
		}),
	}
}
