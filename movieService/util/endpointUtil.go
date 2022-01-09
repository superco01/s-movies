package util

import (
	"net/url"
	"search/movieService/constant"
	"strconv"
	"strings"
)

type ComposeEndPoint interface {
	ComposeOMDBEndPoint(searchWord string, pagination int) string
}

type ComposeEndPointUtil struct{}

func NewUtil() ComposeEndPoint {
	return ComposeEndPointUtil{}
}

func (ComposeEndPointUtil) ComposeOMDBEndPoint(searchWord string, pagination int) string {

	var endPoint = urlFactory(constant.OMDB_BASE_URL, constant.OMDB_API_KEY, url.QueryEscape(searchWord), pagination)

	return endPoint
}

func urlFactory(url, apiKey, searchWord string, pagination int) string {
	var endPoint strings.Builder

	endPoint.WriteString(url)
	if apiKey != "" {
		endPoint.WriteString("?")
		endPoint.WriteString("apikey=")
		endPoint.WriteString(apiKey)
	}
	if searchWord != "" {
		endPoint.WriteString("&")
		endPoint.WriteString("s=")
		endPoint.WriteString(searchWord)
	}
	if pagination != 0 {
		endPoint.WriteString("&")
		endPoint.WriteString("page=")
		endPoint.WriteString(strconv.Itoa(pagination))
	}

	return endPoint.String()
}
