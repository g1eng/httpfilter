package ipfilter

import "net"

//IPFilter is an (experimental) struct represents IP filter
type IPFilter struct {
	Policy  bool       // Policy specifies matching direction of the filtering rule. true for allow, false for deny
	IPAddr  net.IP     // IPAddr is target IP to apply the filter with a policy
	NetMask *net.IPNet //ignored (but required to initialize)
}

//NewIPFilter return a new IPFilter instance with specified cidr.
func NewIPFilter(policy bool, cidr string) (*IPFilter, error) {
	ip, mask, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	return &IPFilter{
		policy,
		ip,
		mask,
	}, nil
}
