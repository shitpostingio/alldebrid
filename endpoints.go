package alldebrid

const (
	// MagnetURL is the endpoint for the magnet API
	magnet = "https://api.alldebrid.com/v4/magnet"

	magnetupload  = "%s/upload?agent=%s&apikey=%s"
	magnetstatus  = "%s/status?agent=%s&apikey=%s&id=%s"
	magnetdelete  = "%s/delete?agent=%s&apikey=%s&id=%s"
	magnetrestart = "%s/restart?agent=%s&apikey=%s&id=%s"
	magnetinstant = "%s/instant?agent=%s&apikey=%s"

	// LinksURL is the endpoint for the link API
	links = "https://api.alldebrid.com/v4/link"

	linkunlock    = "%s/unlock?agent=%s&apikey=%s&link=%s"
	linkstreaming = "%s/streaming?agent=%s&apikey=%s&stream=%s&id=%s"
	linkdelayed   = "%s/delayed?agent=%s&apikey=%s&id=%s"

	// Hosts is the endpoint for the hosts API
	hosts = "https://api.alldebrid.com/v4/hosts"

	hostsall     = "%s?agent=%s"
	hostsdomains = "%s/domains?agent=%s"

	// User is the endpoint for user API
	user = "https://api.alldebrid.com/v4/user"

	userinfo = "%s?agent=%s&apikey=%s"

	// Pin is the endpoint for PIN auth API
	pin = "https://api.alldebrid.com/v4/pin"

	pinget   = "%s/get?agent=%s"
	pincheck = "%s/check?agent=%s&check=%s&pin=%s"
)

func getMagnetEndpointRegular() string {
	return magnet
}

var getMagnetEndpoint = getMagnetEndpointRegular

func getLinksEndpointRegular() string {
	return links
}

var getLinksEndpoint = getLinksEndpointRegular

func getHostsEndpointRegular() string {
	return hosts
}

var getHostsEndpoint = getHostsEndpointRegular

func getUserEndpointRegular() string {
	return user
}

var getUserEndpoint = getUserEndpointRegular

func getPinEndpointRegular() string {
	return pin
}

var getPinEndpoint = getPinEndpointRegular
