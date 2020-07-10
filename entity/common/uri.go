package common

import (
	"net/url"
	"path"
	"strconv"
)

const (
	PageParam           = "page"
	PerPageParam        = "per_page"
	DefaultPerPageParam = 10
)

// GetPaginationQueryParams returns the query params map to request a pagination.
func GetPaginationQueryParams(page int, txPerPage int) map[string]string {
	return map[string]string{
		PageParam:    strconv.Itoa(page),
		PerPageParam: strconv.Itoa(txPerPage),
	}
}

// BuildUri joins the base path and paths array and adds the query values to return a new url.
func BuildUri(basePath string, paths []string, queryValues map[string]string) (*url.URL, error) {
	uri, err := url.Parse(basePath)
	if err != nil {
		return nil, err
	}
	uri.Path = path.Join(append([]string{uri.Path}, paths...)...)
	uriQuery := uri.Query()
	for index, value := range queryValues {
		uriQuery.Set(index, value)
	}
	uri.RawQuery = uriQuery.Encode()
	return uri, nil
}
