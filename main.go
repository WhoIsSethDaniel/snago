package main

import (
	"fmt"
	"os"

	"github.com/mholt/archiver"
)

func downloadPackage(f []GoVersions) (GoPkg, error) {
	p, err := packageToDownload(f)
	if err != nil {
		return GoPkg{}, err
	}
	fmt.Printf("Retrieving file %s\n", p.Filename)
	if err := retrievePackage(p); err != nil {
		return GoPkg{}, err
	}
	return p, nil
}

func installPackage(f, dir string) error {
	fmt.Printf("Installing package to %s\n", dir)
	if err := archiver.Unarchive(f, dir); err != nil {
		return err
	}
	return nil
}

func listVersions(f []GoVersions) {
	for _, v := range f {
		fmt.Printf("%s\n", v.Version)
	}
}

func errorAndExit(s string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, s, a...)
	os.Exit(1)
}

func main() {
	if err := parseConfig(); err != nil {
		fmt.Fprint(os.Stderr, err)
	}

	u := listingsURL(needAllVersions())
	rawJSON, err := getListings(u)
	if err != nil {
		errorAndExit("Failed to get package listings: %s\n", err)
	}

	if *Config.Raw {
		fmt.Printf("%s", rawJSON)
	} else {
		f, err := decodeJSON(rawJSON)
		if err != nil {
			errorAndExit("Failed to decode JSON: %s\n", err)
		}
		if hasVersion() {
			var p GoPkg
			if p, err = downloadPackage(f); err != nil {
				errorAndExit("Failed to download package: %s\n", err)
			}
			if err := installPackage(p.Filename, *Config.Dir); err != nil {
				errorAndExit("Failed to unarchive file: %s\n", err)
			}
		} else {
			listVersions(f)
		}
	}
}
