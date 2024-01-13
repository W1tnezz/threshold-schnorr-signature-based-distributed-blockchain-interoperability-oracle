#!/bin/sh

cd "$(dirname "$0")" || exit 1


baseDir=".."

# solc --optimize --abi $baseDir/contracts/RegistryContract.sol --overwrite -o $baseDir/build/contracts/abi --base-path $baseDir/contracts --include-path $baseDir/node_modules
solc --optimize --abi $baseDir/contracts/Registry.sol --overwrite -o $baseDir/build/contracts/abi --base-path $baseDir/contracts
solc --optimize --abi $baseDir/contracts/DKG.sol --overwrite -o $baseDir/build/contracts/abi --base-path $baseDir/contracts
solc --optimize --abi $baseDir/contracts/Oracle.sol --overwrite -o $baseDir/build/contracts/abi --base-path $baseDir/contracts

# abigen --abi $baseDir/build/contracts/abi/RegistryContract.abi --pkg iop --type RegistryContract --out ../pkg/iop/registrycontract.go
abigen --abi $baseDir/build/contracts/abi/Registry.abi --pkg node --type Registry  --out $baseDir/../node/pkg/node/registry.abi.go
abigen --abi $baseDir/build/contracts/abi/DKG.abi --pkg node --type DKG --out $baseDir/../node/pkg/node/dkg.abi.go
abigen --abi $baseDir/build/contracts/abi/Oracle.abi --pkg node --type OracleContract --out $baseDir/../node/pkg/node/oracle.abi.go