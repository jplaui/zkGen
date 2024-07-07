## zkGen Transpiler Repository

## Prerequisites 
Install `Golang` on your system [follow docs](https://go.dev/doc/install).

## How to run zkGen
_zkGen_ is a command line toolkit and you can see available commands by running `go run main.go help` from the repository root location.
Before any commands can be run, make sure to clone the repository with `git clone git@github.com:anonsubsub/zkGen.git`, jump into the repository with `cd zkGen`, and `go mod tidy`. (if no go.mod file exists, make sure to run `go mod init transpiler`).

Available commands are:
- use `go run main.go gadgetlib-list` to show all available library files.
- use `go run main.go zkpolicy-list` to show all available zkPolicies.
- use `go run main.go zkpolicy-transpile <zkPolicy-name>` to transpile a zkPolicy and generate a circuit (e.g. use `go run main.go zkpolicy-transpile example3_zkFriendlyCommitData` to transpile policy-compliant commitment verification). The generated circuit is stored in the `circuits` folder.
- use `go run main.go zkpolicy-test <zkPolicy-name> <solidity>` to test the generated zk circuit via the supported ZKP system. The solidity is used to generate smart contract ZKP verification code of the respective circuit. For instance, the command `go run main.go zkpolicy-test example3_zkFriendlyCommitData solidity` runs the generated compliant commitment circuit via the plonk ZKP system and generates the solidity code inside the `circuits` folder.

## Citing
We welcome you to cite our [research paper](https://tum-esi.github.io/publications-list/PDF/2024-ICBC-zkGen.pdf) if you are using _zkGen_ in academia.
```
@inproceedings{les2023zkGen,
    author = {Lauinger, Jan and Ernstberger, Jens and Steinhorst, Sebastian},
    title = {zkGen: Policy-to-Circuit Transpiler},
    year = {2024},
    publisher = {2024 IEEE International Conference on Blockchain and Cryptocurrency (ICBC24)}
}
```