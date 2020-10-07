package main

import (
	"encoding/json"
)

type GoPkg struct {
	Filename string
	OS       string
	Arch     string
	Version  string
	Sha256   string
	Size     uint
	Kind     string
}

type GoVersions struct {
	Version string
	Stable  bool
	Files   []GoPkg
}

func decodeJSON(raw []byte) ([]GoVersions, error) {
	var f []GoVersions
	if err := json.Unmarshal(raw, &f); err != nil {
		return nil, err
	}
	return f, nil
}
