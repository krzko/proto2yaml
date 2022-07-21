# proto2yaml

🔄 A command-line utility to export Protocol Buffers (proto) files to YAML, and JSON.

Currently supported exports are for:

* Packages
* Services
* RPCs

Supported filters are for:

* Options

## Overview

We needed an intermediatate format to allow us to provision [Service Level Objective](https://cloud.google.com/service-mesh/docs/observability/slo-overview) resources using [terraform](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/monitoring_slo).

The idea is to enumerate the associated protos and parse the export using a `yamldecode()` or `jsondecode()` function and prepolute our variables.

Addition features such as filtering and [OpenSLO](https://github.com/OpenSLO/OpenSLO) export formating coming.

## Output

`proto2yaml` will export your `.proto` files using the following format;

```yaml
version: v0.0.5
packages:
  - package: foo.service.profile.v1
    services:
      - service: ProfileAPI
        rpc:
          - name: GetProfiles
            type: unary
      - service: AdminAPI
        rpc:
          - name: CreateProfile
            type: unary
          - name: DeleteProfile
            type: unary
          - name: SyncProfile
            type: server-streaming
```

## Getting started

Running `proto2yaml` is availabile through several methods. You can using `brew`, download it as a binary from GitHub releases, or running it as a distroless docker image.

### brew

Install [brew](https://brew.sh/) and then run:

```sh
brew install krzko/tap/proto2yaml
```

### Download Binary

Download the latest version from the [Releases](https://github.com/krzko/proto2yaml/releases) page.

### Docker

Attach a [bind mount](https://docs.docker.com/storage/bind-mounts/#start-a-container-with-a-bind-mount) to the source directory and the directory you want to export the file to.

To see all the tags view the [Packages](https://github.com/krzko/proto2yaml/pkgs/container/proto2yaml) page.

To run the docker image follow these examples:

```sh
# Use current directory as source
docker run --rm \
    -v "$(pwd)":/searchme \
    ghcr.io/krzko/proto2yaml:latest yaml print --source /searchme

# Use an explicit path as source
docker run --rm \
    -v "/Users/foobar/code/protos":/searchme \
    ghcr.io/krzko/proto2yaml:latest yaml print --source /searchme

# Use an explicit path as source and current as export
docker run --rm \
    -v "/Users/foobar/code/protos":/searchme \
    -v "$(pwd)":/save \
    ghcr.io/krzko/proto2yaml:latest yaml print --source /searchme --file /save/example_protos.yaml
```

## Run

```sh
NAME:
   proto2yaml - A command-line utility to convert Protocol Buffers (proto) files to YAML

USAGE:
   proto2yaml [global options] command [command options] [arguments...]

VERSION:
   proto2yaml version v0.0.6-9ad396c (2022-07-21T05:33:05Z)

COMMANDS:
   json     The outputs are formatted as JSON
   yaml     The outputs are formatted as YAML
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

### Export

To export to a file, run the following command

```sh
# JSON
proto2yaml json export --source ./protos --file ./example_protos.json
# JSON pretty
proto2yaml json export --source ./protos --file ./example_protos.json --pretty

# YAML
proto2yaml yaml export --source ./protos --file ./example_protos.yaml
```

### Filter

To filter on an `option` you can use the `--exclude-option` or `--include-option` filter. For now its based on singletons but hope to expand out multiple combinations in the future. An example is as follows:

```sh
proto2yaml yaml print --source ./protos --exclude-option "deprecated=true"
```

Or run the inverse of the above using:

```sh
proto2yaml yaml print --source ./protos --include-option "deprecated=true"
```

### Print

To print to the console, clone the repo and run the following command:

```sh
# JSON
proto2yaml json print --source ./protos
# JSON pretty
proto2yaml json print --source ./protos --pretty

# YAML
proto2yaml yaml print --source ./protos
```

### Disable Colour

If you need to run the tool in your CI/CD pipelines and ANSI isn't supported, you can pass the following variable to disable colour:

```sh
export NO_COLOR="true"
```

To enable colour again, simply `unset` the variable:

```sh
unset NO_COLOR
```

## Build & Run

To build to the binaries use the following targets. All outputs are generated to the `bin` directory.

### All

To build all the the targets, simply run:

```sh
# Default target invoked
make

# Explicit target
make build
```

To run all the builds without Docker, simply run:

```sh
make build-no-docker
```

### Linux

The following targets will generate your Linux binaries:

```sh
make build-linux
```

### macOS

macOS has two targets, one for the older Intel `amd64` CPUs and one for the newer Mx `arm64` CPUs. The following targets will generate your binaries:

```sh
# For intel macs
make build-darwin-amd64
# For m1 macs
make build-darwin-arm64
```

### Windows

The following targets will generate your Windows executable:

```sh
make build-windows
```

### Others

The general `make build` and `back-build-no-docker` targets will also build **Raspberry Pi** and **FreeBSD** version, along with a **distroless Docker** image, if selected.
