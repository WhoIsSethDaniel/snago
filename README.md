# snago - grab a package for a release of Go; list Go releases

## What is it?

Use this tool to list current Go releases, old Go releases, and to download and unpack the archive for a particular OS
and architecture.

## Quick Usage

```
# List all recent stable releases
snago 

# List all releases ever
snago -all

# Show the raw JSON
snago -raw

# Show the raw JSON for all releases ever
snago -all -raw

# Download version 1.12.1 (for the default OS and arch - linux/amd64)
snago 1.12.1

# Download version 1.12.1 for Darwin on amd64
snago 1.12.1 darwin/amd64

# Download version 1.12.1 (default os and arch) and unpack in ~/.my_go
snago -dir ~/.my_go 1.12.1 
```

## Build and Install

Assuming you have 'go' already installed you can 

```
make
```

and this should build snago.

To install simply copy 'snago' somewhere in your $PATH.

## History

This is the first thing I ever wrote in Go. I wrote it back in 2016 or so. I don't even think the download URL I use
now existed then. Once I discovered that it existed I added the listing feature. All this work was performed on my old
GitLab account. When I decided to move everything over to Github (October 2020) I simply copied it all over, created new
repos, and committed. This had the exciting property of losing all history. I was okay with this, some may not be.

## Future

I may someday make this utility more user friendly. Right now the errors are pretty raw, some of the functionality may
not be quite as expected. It's a pretty simple tool though. I use it often. I use it to find new releases and to
download them, which is its entire purpose.
