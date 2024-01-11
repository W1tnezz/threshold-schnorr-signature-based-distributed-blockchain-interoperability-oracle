#!/bin/sh

cd "$(dirname "$0")" || exit 1


baseDir=".."

# solc --optimize --abi $baseDir/contracts/RegistryContract.sol --overwrite -o $baseDir/build/contracts/abi --base-path $baseDir/contracts --include-path $baseDir/node_modules
solc --optimize --abi $baseDir/contracts/Registry.sol --overwrite -o $baseDir/build/contracts/abi --base-path $baseDir/contracts
solc --optimize --abi $baseDir/contracts/Sakai.sol --overwrite -o $baseDir/build/contracts/abi --base-path $baseDir/contracts
solc --optimize --abi $baseDir/contracts/IBSAS.sol --overwrite -o $baseDir/build/contracts/abi --base-path $baseDir/contracts

# abigen --abi $baseDir/build/contracts/abi/RegistryContract.abi --pkg iop --type RegistryContract --out ../pkg/iop/registrycontract.go
abigen --abi $baseDir/build/contracts/abi/Registry.abi --pkg signer --type Registry  --out $baseDir/../signer/pkg/signer/registry.abi.go
abigen --abi $baseDir/build/contracts/abi/Sakai.abi --pkg signer --type Sakai  --out $baseDir/../signer/pkg/signer/sakai.abi.go
abigen --abi $baseDir/build/contracts/abi/IBSAS.abi --pkg signer --type IBSAS  --out $baseDir/../signer/pkg/signer/ibsas.abi.go