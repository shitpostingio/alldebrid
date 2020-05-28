package alldebrid

// Useful test struct
var noapicl = &Client{
	ic: &innerClient{
		apikey:  "",
		appName: "test",
	},
}

var cl = &Client{
	ic: &innerClient{
		apikey:  "123456abcdef",
		appName: "test",
	},
}
