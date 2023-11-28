# zkGen

#### Contents
- [Disclaimer](#disclaimer)
- [Overview](#overview)
- [Installation](#installation)
- [Command Line Toolkit](#command-line-toolkit)
- [Repository Structure](#repository-structure)
- [Tutorials](#tutorials)
- [Evaluation](#evaluation)
- [Contributions](#contributions)
- [Citing](#citing)

## Disclaimer
_zkGen_ is in an early stage development and does not support API changes of supported secure computation frameworks. We make no guarantees that the generated circuits are secure. Make sure that any generated circuits undergone a security audit before using them in a production system or before using them to process sensitive/confidential data.

## Overview
_zkGen_ is a decentralized oracle that doesn't rely on the trusted hardware and two-party computation. It allows users to prove the *provenance* of TLS records from a specific API and, in the meanwhile, prove the dynamic content fulfills a value constraint as pre-defined in a public policy in zero-knowledge.  It achieves confidentiality and integrity under certain honesty assumptions (e.g. honestly behaving proxy) which we define in the associated paper [HERMES](https://link-to-hermes.com). _zkGen_ relies on a proxy sitting between the client user and the server hosting the API. The TLS stack running on the server does not require any modifications, making _zkGen_ extendable to various existent APIs and web resources.


## initial setup
- `go mod init transpiler`
- `go mod tidy`

## Repository structure
The structure of the project is as follows:
```
    ├── circuits	# self-signed certificates for local development
    ├── commands	# command line toolkit logic
    ├── dsl     	# contains additional git submodules
    ├── templates	# containerization
    ├── docs		# documentation files
```

## Command Line Toolkit
_zkGen_ is a command-line toolkit which allows users to (i) transpile policies into SNARK circuits, (ii) generate input data for policy compliant data provenance proofs from private APIs, and (iii) prove policy-compliant proofs in zero-knowledge. The high-level workflow of the _zkGen_ command-line toolkit is as follows:

1. (proxy) start the proxy service
2. (server) start the server service (optional, required for local testing)
3. (prover) transpile public policy into a circuit generator
4. (prover) collect circuit input data by
	- refreshing API credentials (optional, if local deployment tutorial is used)
	- request API defined in policy and post-process tls and record layer traffic
5. (prover) generate snark-circuit in arithmetic representation and compute witness data
6. (prover) compute setup and generate zkp
7. (proxy) postprocess captured traffic transcript and collect zkp public input
8. (proxy) verify zkp


## Tutorials
We provide all details of how to execute _zkGen_ in different deployments in our [tutorial](./docs/tutorials) guidelines. Before you start running the tutorials, please follow our [installation instructions](./docs/00_installation.md) to correctly set up the repository and its dependencies.


## Evaluation
You can evaluate generated circuits via supported PETs frameworks. To do so, ... 
To reproduce results provided in the research paper [HERMES](https://link-to-hermes.com), we provide an evaluation script [evaluation.sh](./evaluation.sh). The evaluation script can be executed by calling `./evaluation.sh` in the root location of this repository after installing the repository with the `./installation.sh` script as described [here](./docs/00_installation.md).


## Limitations 
* The constraint system currently supported by _zkGen_ public policies is very limited with value proofs of float greater than `GT`, float less than `LT` and string equality `EQ` and will be extended in the future.

## Citing
We welcome you to cite our [paper](https://link-to-hermes.com) if you are using _zkGen_ in academia.
```
@inproceedings{2023zkGen,
    author = {},
    title = {zkGen: Transpiler Architecture to Generate
Zero-knowledge Circuits},
    year = {2023},
    publisher = {},
    booktitle = {},
    location = {},
    series = {}
}
```
