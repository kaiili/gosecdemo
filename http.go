package gosecdemo

import "net/http"
func ssrf_http(url string)  {
	http.Get(url)
	http.Head(url)
}


func ssrf_client(url string){
	c := http.DefaultClient
	c.Get(url)
	c.Head(url)
}

