# proto2yaml

🔄 A command-line utility to export Protocol Buffers (proto) files to YAML, and JSON.

Currently supported exports are for:

* Packages
* Services
* RPCs

Supported filters are for:

* Options

## Overview

We needed and intermediatate format to allow us to provision [Service Level Objectives](https://cloud.google.com/service-mesh/docs/observability/slo-overview) resources using [terraform](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/monitoring_slo).

The idea is to enumerate the associated protos and parse the export using a `yamldecode()` or `jsondecode()` function and prepolute our vars.

Addition features such as filtering and [OpenSLO](https://github.com/OpenSLO/OpenSLO) export formating coming.

## Getting started

Download the latest [release](https://github.com/krzko/proto2yaml/releases).

```sh
$ ./proto2yaml
NAME:
   proto2yaml - A command-line utility to convert Protocol Buffers (proto) files to YAML

USAGE:
   proto2yaml [global options] command [command options] [arguments...]

VERSION:
   0.0.2

COMMANDS:
   json     The outputs are formatted as a JSON
   yaml     The outputs are formatted as a YAML
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

### Export

To export to a file, run the following command

```sh
# json
$ ./proto2yaml json export --source ./protos --file ./example_protos.json
# json pretty
$ ./proto2yaml json export --source ./protos --file ./example_protos.json --pretty

# yaml
$ ./proto2yaml yaml export --source ./protos --file ./example_protos.yaml
```

### Filter

To filter on an `option` you can use the `--exclude-option` or `--include-option` filter. For now its based on singletons but hope to expand out multiple combinations in the future. An example is as follows:

```sh
$ proto2yaml yaml print --source ./protos --exclude-option "deprecated=true"
```

Or run the inverse of the above using:

```sh
$ proto2yaml yaml print --source ./protos --include-option "deprecated=true"
```

### Print

To print to the console, clone the repo and run the following command:

```sh
# json
$ ./proto2yaml json print --source ./protos
# json pretty
$ ./proto2yaml json print --source ./protos --pretty

# yaml
$ ./proto2yaml yaml print --source ./protos
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
