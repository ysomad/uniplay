// Code generated by go-swagger; DO NOT EDIT.

package player

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// GetPlayerListURL generates an URL for the get player list operation
type GetPlayerListURL struct {
	LastSteamID *string
	PageSize    *int32
	Search      *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetPlayerListURL) WithBasePath(bp string) *GetPlayerListURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetPlayerListURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetPlayerListURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/players"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/v1"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var lastSteamIDQ string
	if o.LastSteamID != nil {
		lastSteamIDQ = *o.LastSteamID
	}
	if lastSteamIDQ != "" {
		qs.Set("last_steam_id", lastSteamIDQ)
	}

	var pageSizeQ string
	if o.PageSize != nil {
		pageSizeQ = swag.FormatInt32(*o.PageSize)
	}
	if pageSizeQ != "" {
		qs.Set("page_size", pageSizeQ)
	}

	var searchQ string
	if o.Search != nil {
		searchQ = *o.Search
	}
	if searchQ != "" {
		qs.Set("search", searchQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetPlayerListURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetPlayerListURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetPlayerListURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetPlayerListURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetPlayerListURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetPlayerListURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}