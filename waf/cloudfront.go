package waf

import (
	"net/http"
	"strings"
	"regexp"
)

func CloudFront(headers http.Header, str string) bool {
	match1, _ := regexp.MatchString("Cloudfront", headers.Get("Server"))
	match2, _ := regexp.MatchString(`([0-9\.]+?)? \w+?\.cloudfront\.net \(Cloudfront\)`, headers.Get("Via"))
	_, match3 := headers["X-Amz-Cf-Id"]
	if match1 || match2 || match3 || strings.Contains(str, "Generated by cloudfront (CloudFront)") || strings.Contains(headers.Get("X-Cache"), "Error from Cloudfront") {
		return true
	}
	return false
}