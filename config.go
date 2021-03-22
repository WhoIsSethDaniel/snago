package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var Config struct {
	Version string
	All     *bool
	Help    *bool
	Raw     *bool
	Dir     *string
	OS      string
	Arch    string
}

func hasVersion() bool {
	return Config.Version != ""
}

func needAllVersions() bool {
	if hasVersion() || *Config.All {
		return true
	}
	return false
}

func parseConfig() error {
	flag.Usage = func() {
		progname := os.Args[0]
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [opts] [version] [os/arch]\n", progname)
		fmt.Fprint(flag.CommandLine.Output(), "\n-raw and -all mixed with version is disallowed\n")
		fmt.Fprint(flag.CommandLine.Output(), "\nDefaults:\n")
		fmt.Fprintf(flag.CommandLine.Output(), "OS = %s\n", runtime.GOOS)
		fmt.Fprintf(flag.CommandLine.Output(), "Arch = %s\n", runtime.GOARCH)
		fmt.Fprint(flag.CommandLine.Output(), "\nExamples:\n")
		fmt.Fprint(flag.CommandLine.Output(), "Print the most recent versions\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  %s\n", progname)
		fmt.Fprint(flag.CommandLine.Output(), "Print all versions with raw json\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  %s -raw -all\n", progname)
		fmt.Fprint(flag.CommandLine.Output(), "Retrieve the 1.15.1 release package for the default OS/Architecture\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  %s 1.15.1\n", progname)
		fmt.Fprint(flag.CommandLine.Output(), "Retrieve the 1.15.1 release package for darwin/amd64\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  %s 1.15.1 darwin/amd64\n", progname)
		fmt.Fprint(flag.CommandLine.Output(), "Retrieve the 1.15.1 release package and unpack it to ~/go\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  %s -dir ~/go 1.15.1\n", progname)
		fmt.Fprintf(flag.CommandLine.Output(), "\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Options:\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	Config.All = flag.Bool("all", false, "Return all known Go versions")
	Config.Help = flag.Bool("help", false, "This help")
	Config.Raw = flag.Bool("raw", false, "Return the raw JSON")
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	Config.Dir = flag.String("dir", filepath.Join(home, ".go"), "Directory to unpack archive into")
	flag.Parse()
	Config.OS = runtime.GOOS
	Config.Arch = runtime.GOARCH
	if flag.NArg() > 2 {
		flag.Usage()
	}
	if flag.NArg() == 2 {
		osArch := strings.Split(flag.Arg(1), "/")
		if len(osArch) != 2 {
			return errors.New("OS and architecture must be of the form 'os/arch'")
		}
		Config.OS = osArch[0]
		Config.Arch = osArch[1]
	}
	Config.Version = flag.Arg(0)

	if *Config.Help {
		flag.Usage()
	}

	if (*Config.All || *Config.Raw) && Config.Version != "" {
		return errors.New("-raw and -all may not be given along with a version")
	}
	return nil
}
