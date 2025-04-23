package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

// curl -o dl.json https://golang.org/dl/?mode=json (that last '/' is important)
// curl -o dl.json https://golang.org/dl/?mode=json&include=all (every release available for
// download)
const (
	listURL = "https://go.dev/dl/"
	dlURL   = "https://dl.google.com/go/"
)

func listingsURL(all bool) url.URL {
	// highly unlikely to fail
	url, _ := url.ParseRequestURI(listURL)

	v := url.Query()
	v.Set("mode", "json")
	url.RawQuery = v.Encode()
	if all {
		v := url.Query()
		v.Set("include", "all")
		url.RawQuery = v.Encode()
	}
	return *url
}

func getListings(u url.URL) ([]byte, error) {
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func packageToDownload(f []GoVersions) (GoPkg, error) {
	for _, vrec := range f {
		if vrec.Version == "go"+Config.Version {
			for _, frec := range vrec.Files {
				if frec.OS == Config.OS && frec.Arch == Config.Arch {
					return frec, nil
				}
			}
		}
	}
	return GoPkg{}, fmt.Errorf(
		"unable to find package for %s %s/%s",
		Config.Version,
		Config.OS,
		Config.Arch,
	)
}

func retrievePackage(p GoPkg) error {
	// can also look at timeouts in the Transport struct,
	// but this seems to be what is needed;
	// probably should make the timeout configurable
	client := &http.Client{Timeout: 120 * time.Second}
	resp, err := client.Get(dlURL + "/" + p.Filename)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return writePackage(p.Filename, resp.Body)
}
