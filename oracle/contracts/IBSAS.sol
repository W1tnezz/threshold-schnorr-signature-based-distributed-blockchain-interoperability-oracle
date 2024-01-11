// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "./crypto/BN256G1.sol";
import "./crypto/BN256G2.sol";
import "./Registry.sol";

contract IBSAS {

    uint256[] private hashSi;

    uint256[] private checkPairingInput;

    Registry private registry;

    constructor(address _registryContract) {
        registry = Registry(_registryContract);
    }

    // --------------------------------------------------------------------------------------------------------------------------------------------------------

    function submit(
        uint256[2] calldata _X,
        uint256[2] calldata _Y,
        uint256[4] calldata _Z,
        uint256[4] calldata _Z1,
        uint256[2] calldata _u,
        uint256[2] calldata _v,
        uint256[4] calldata mpk
    ) external payable {
        bytes32 message = registry.getMessage();
        address[] memory SignOrder = registry.getSignOrder();
        
        // cal all H(si) 
        // question : message || ID1
        bytes memory s = abi.encodePacked(message, registry.getSignerByAddress(SignOrder[0]).identity);

        for(uint i = 1; i < SignOrder.length; i++){
            s = abi.encodePacked(s, registry.getSignerByAddress(SignOrder[i]).identity);
            hashSi.push(uint256(sha256(s)));
        }

        uint256[2] memory combine1;
        uint256[2] memory combine2;
        for(uint i = 0; i < SignOrder.length; i++){
            uint256[2] memory idHash = registry.getSignerByAddress(SignOrder[i]).pubKey;
            (combine2[0], combine2[1]) = BN256G1.addPoint([combine2[0], combine2[1], idHash[0], idHash[1]]);

            uint256 scalar = 1;
            for(uint j = i; j < hashSi.length; j++){
                scalar = mulmod(scalar, hashSi[j], BN256G1.NN);
            }
            scalar = BN256G1.modInverse(scalar);
            uint256 tempX;
            uint256 tempY;
            (tempX, tempY) = BN256G1.mulPoint([idHash[0], idHash[1], scalar]);
            (combine1[0], combine1[1]) = BN256G1.addPoint([combine1[0], combine1[1], tempX, tempY]);
        }
        
        checkPairingInput.push(_Y[0]);
        checkPairingInput.push(_Y[1]);
        checkPairingInput.push(BN256G2.G2_NEG_X_IM);
        checkPairingInput.push(BN256G2.G2_NEG_X_RE);
        checkPairingInput.push(BN256G2.G2_NEG_Y_IM);
        checkPairingInput.push(BN256G2.G2_NEG_Y_RE);
        checkPairingInput.push(_v[0]);
        checkPairingInput.push(_v[1]);
        checkPairingInput.push(_Z[1]);
        checkPairingInput.push(_Z[0]);
        checkPairingInput.push(_Z[3]);
        checkPairingInput.push(_Z[2]);
        checkPairingInput.push(combine1[0]);
        checkPairingInput.push(combine1[1]);
        checkPairingInput.push(mpk[1]);
        checkPairingInput.push(mpk[0]);
        checkPairingInput.push(mpk[3]);
        checkPairingInput.push(mpk[2]);
        require(BN256G1.bn256CheckPairingBatch(checkPairingInput), "first check failed!");

        delete checkPairingInput;
        checkPairingInput.push(_X[0]);
        checkPairingInput.push(_X[1]);
        checkPairingInput.push(BN256G2.G2_NEG_X_IM);
        checkPairingInput.push(BN256G2.G2_NEG_X_RE);
        checkPairingInput.push(BN256G2.G2_NEG_Y_IM);
        checkPairingInput.push(BN256G2.G2_NEG_Y_RE);
        checkPairingInput.push(_u[0]);
        checkPairingInput.push(_u[1]);
        checkPairingInput.push(_Z1[1]);
        checkPairingInput.push(_Z1[0]);
        checkPairingInput.push(_Z1[3]);
        checkPairingInput.push(_Z1[2]);
        checkPairingInput.push(combine2[0]);
        checkPairingInput.push(combine2[1]);
        checkPairingInput.push(mpk[1]);
        checkPairingInput.push(mpk[0]);
        checkPairingInput.push(mpk[3]);
        checkPairingInput.push(mpk[2]);
        require(BN256G1.bn256CheckPairingBatch(checkPairingInput), "second check failed!");
        delete checkPairingInput;
        delete hashSi;
    }
}
