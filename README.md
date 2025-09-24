# ChainCraft

**Author: Vasanthadithya Mundrathi from CBIT college**

Golang implementation of ChainCraft's blockchain infrastructure node types (`light` | `full` | `bridge`).

ChainCraft is a comprehensive blockchain infrastructure project that provides data availability and consensus network solutions. The node types described above comprise the ChainCraft data availability (DA) network.

The DA network wraps the core consensus network by listening for blocks from the consensus network and making them digestible for data availability sampling (DAS), enabling secure and scalable access to blockchain data.

## Table of Contents

- [ChainCraft](#chaincraft)
  - [Table of Contents](#table-of-contents)
  - [Minimum requirements](#minimum-requirements)
  - [System requirements](#system-requirements)
  - [Supported architectures](#supported-architectures)
  - [Installation](#installation)
  - [API docs](#api-docs)
  - [Node types](#node-types)
  - [Run a node](#run-a-node)
    - [Quick Start with Light Node](#quick-start-with-light-node)
  - [Environment variables](#environment-variables)
  - [Package-specific documentation](#package-specific-documentation)

## Minimum requirements

| Requirement | Notes          |
| ----------- | -------------- |
| Go version  | 1.24 or higher |

## System requirements

See the [official docs page](https://docs.chaincraft.org/how-to-guides/nodes-overview#data-availability-nodes) for system requirements for each node type.

## Supported architectures

ChainCraft officially supports the following architectures:

- `linux/amd64`
- `linux/arm64`
- `darwin/amd64` (macOS Intel)
- `darwin/arm64` (macOS Apple Silicon)

Only these four architectures are officially tested and supported.

## Installation

```sh
# Build from source
make build
sudo make install
```

For more information on setting up a node and the hardware requirements needed, go visit our docs at <https://docs.chaincraft.org>.

## API docs

The ChainCraft node public API is documented [here](https://node-rpc-docs.chaincraft.org/).

## Node types

- **Bridge** nodes - relay blocks from the consensus network to the data availability (DA) network
- **Full** nodes - fully reconstruct and store blocks by sampling the DA network for shares
- **Light** nodes - verify the availability of block data by sampling the DA network for shares

More information can be found in the documentation.

## Run a node

`<node_type>` can be: `bridge`, `full` or `light`.

```sh
chaincraft <node_type> init
```

```sh
chaincraft <node_type> start
```

Please refer to [this guide](https://docs.chaincraft.org/how-to-guides/chaincraft-node/) for more information on running a node.

### Quick Start with Light Node

View available commands and their usage:

```sh
make node-help
```

Install chaincraft node and key binaries:

```sh
make node-install
```

Start a light node with automated setup:

```sh
make light-node-up
```

This command:

- Automatically checks wallet balance
- Requests funds from faucet if needed
- Sets node height to latest-1 for quick startup
- Initializes the node if running for the first time

Options:

```sh
make light-node-up COMMAND=again    # Reset node state to latest height
make light-node-up CORE_IP=<ip>     # Use custom core IP
```

## Environment variables

| Variable                | Explanation                         | Default value | Required |
| ----------------------- | ----------------------------------- | ------------- | -------- |
| `CHAINCRAFT_BOOTSTRAPPER` | Start the node in bootstrapper mode | `false`       | Optional |

## Package-specific documentation

- [Header](./header/doc.go)
- [Share](./share/doc.go)
- [DAS](./das/doc.go)
