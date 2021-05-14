package restMiddleware

import "net/http"

/*
	GetIP
	Get the ip from cloudflare or fall back to the actual request ip.

	@param {*http.Request} r - The request pointer we want to scrape.

	@returns {string} ip address with port.
*/
func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("HTTP_CF_CONNECTING_IP")

	if forwarded != "" {
		return forwarded
	}

	forwarded2 := r.Header.Get("cf-connecting-ip")

	if forwarded2 != "" {
		return forwarded2
	}

	return r.RemoteAddr
}
