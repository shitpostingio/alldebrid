package alldebrid

const (
	// MagnetURL is the endpoint for the magnet API
	magnetURL = "https://api.alldebrid.com/v4/magnet"

	upload  = "%s/upload?agent=%s&apikey=%s"
	status  = "%s/status?agent=%s&apikey=%s&id=%s"
	delete  = "%s/delete?agent=%s&apikey=%s&id=%s"
	restart = "%s/restart?agent=%s&apikey=%s&id=%s"

	// LinksURL is the endpoint for the link API
	linksURL = "https://api.alldebrid.com/v4/link"

	unlock    = "%s/unlock?agent=%s&apikey=%s&link=%s"
	streaming = "%s/streaming?agent=%s&apikey=%s&stream=%s&id=%s"
	delayed   = "%s/delayed?agent=%s&apikey=%s&id=%s"
)
