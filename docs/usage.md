# Usage 

## Setup
### Installation

### Brew

```shell
brew install janosmiko/tap/vlt
```
to upgrade
```shell
brew update && brew upgrade vlt
```

### Download from GitHub

Download the relevant binary for your operating system (macOS = Darwin) from
the [latest Github release](https://github.com/janosmiko/vlt/releases). Unpack it, then move the binary to
somewhere accessible in your `PATH`, e.g. `mv ./vlt /usr/local/bin`.

### > Using [go installed on your machine](https://go.dev/doc/install)

```shell
go install github.com/janosmiko/vlt@latest
```

### Building from source and Run vlt

Make sure you have your go environment setup:

1. Clone the project
1. Run `$ make build` to build the binary
1. Run `$ make run` to run the binary
1. You can use `$ make install-osx` on a Mac to cp the binary to `/usr/local/bin/vlt`

or

```
$ go install ./cmd/vlt
```

### How to use it

Once `vlt` is installed and avialable in your path, simply run:

```
$ vlt
```

![image](../images/screen1.png)


### Environment variables

In order to use the tool you must expose the needed env variables, that would generally be used by the vault cli to auth to a given cluster. 

Required:  
`VAULT_ADDR`  
`VAULT_TOKEN`

For the full list see the [official docs](https://developer.hashicorp.com/vault/docs/commands#environment-variables)

Another option is to store your configs in yaml file named `.vlt.yaml` stored in your home directory.  
Example: [`~/myuser/.vlt.yaml`](./examples/vlt.yaml)

Or alternatively pass a config file as an argument using `-c <path/file.yaml>`  
Example: `vlt -c ./new-env.yml`

#### Authentication and variables priority
Variables will be loaded in the following order, with the next superseding the previous ones:

1. Will check for vault [token cache](https://developer.hashicorp.com/vault/docs/commands#authenticating-to-vault)
2. Read from env variables
3. Config file 

### Features

Currently the capabilities are limited. 

* Support for navigation between KV mounts
    * Currently only KV2
* Looking up secret objects
    * Show/hide secrets and coping data
    * Update/patch secrets
    * Create new secrets
    * Filter paths/secrets 
* Support for exploring and filtering ACL Policies
* Namespace support for Enteprise versions
