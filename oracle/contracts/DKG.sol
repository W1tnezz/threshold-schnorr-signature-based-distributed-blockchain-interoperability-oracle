// SPDX-License-Identifier: MIT
pragma solidity >0.8.0;

import "./crypto/BN256G1.sol";
import "./crypto/BN256G2.sol";
import "./Registry.sol";

contract DKG {
    Registry private registry;

    constructor(address _registryContract) {
        registry = Registry(_registryContract);
    }

    event DistKey(uint256[2] pubKey);

    uint256 private remain;

    uint256[2] public Y;

    address[] private validators;

    function needEnroll() public returns (bool) {
        if(remain == 0){
            delete validators;
            return true;
        }
        return false;
    }

    function enroll() external payable returns (bool) {
        if (remain != 0) {
            return false;
        }
        for (uint i = 0; i < validators.length; i++) {
            require(validators[i] != msg.sender, "ENROLLED");
        }

        validators.push(msg.sender);
        if (validators.length >= (registry.countOracleNodes() - 1)) {
            distKey();
            remain = 4;
        }
        return true;
    }

    function distKey() private {
        uint256[2] memory key;
        for (uint i = 0; i < validators.length; i++) {
            int256 lambda1 = 1;
            int256 lambda2 = 1;
            int256 indexI = int256(registry.getIndex(validators[i]));
            for (uint j = 0; j < validators.length; j++) {
                if (i != j) {
                    int256 indexJ = int256(registry.getIndex(validators[j]));
                    lambda1 *= (-indexJ);
                    lambda2 = lambda2 * (indexI - indexJ);
                }
            }
            uint256 ulambda1 = BN256G1.mod(lambda1);
            uint256 ulambda2 = BN256G1.modInverse(BN256G1.mod(lambda2));
            uint256 lambda = mulmod(ulambda1, ulambda2, BN256G1.NN);
            registry.setLambda(validators[i], lambda);

            uint256 x = registry.getNodeByAddress(validators[i]).pubKey[0];
            uint256 y = registry.getNodeByAddress(validators[i]).pubKey[1];
            (x, y) = BN256G1.mulPoint([x, y, lambda]);
            (key[0], key[1]) = BN256G1.addPoint([key[0], key[1], x, y]);
        }
        Y[0] = key[0];
        Y[1] = key[1];

        emit DistKey(Y);
    }

    function getPubKey() public view returns (uint256[2] memory) {
        return Y;
    }

    function usePubKey() public returns (uint256[2] memory) {
        remain--;
        return Y;
    }

    function getValidators() public view returns (address[] memory) {
        return validators;
    }
}
