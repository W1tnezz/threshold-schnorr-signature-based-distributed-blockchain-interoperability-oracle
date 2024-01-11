// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "./crypto/BN256G1.sol";
import "./crypto/BN256G2.sol";
import "./Registry.sol";

contract Sakai {
    uint256[] private randoms;

    uint256[2][] private hashPointSequence;

    uint256[] private checkPairingInput;

    Registry private registry;

    constructor(address _registryContract) {
        registry = Registry(_registryContract);
    }

    // --------------------------------------------------------------------------------------------------------------------------------------------------------

    function submit(
        uint256[4] calldata masterPubKey,
        uint256[2][] calldata signatures,
        uint256[4][] calldata setOfR
    ) external payable {
        bytes32 message = registry.getMessage();
        address[] memory SignOrder = registry.getSignOrder();
        require(SignOrder.length == signatures.length, "sig nums error");

        for (uint i = 0; i < signatures.length; i++) {
            randoms.push(
                uint256(
                    keccak256(
                        abi.encodePacked(
                            signatures[i],
                            block.prevrandao,
                            msg.sender
                        )
                    )
                )
            );
        }

        // cal (s1 ^ p1) * (s2 ^ p2) * (s3 ^ p3) ...
        {
            uint256 Sx = 0;
            uint256 Sy = 0;
            for (uint i = 0; i < signatures.length; i++) {
                uint256 tempX;
                uint256 tempY;
                (tempX, tempY) = BN256G1.mulPoint(
                    [signatures[i][0], signatures[i][1], randoms[i]]
                );
                (Sx, Sy) = BN256G1.addPoint([Sx, Sy, tempX, tempY]);
            }
            checkPairingInput.push(Sx);
            checkPairingInput.push(Sy);
            checkPairingInput.push(BN256G2.G2_NEG_X_IM);
            checkPairingInput.push(BN256G2.G2_NEG_X_RE);
            checkPairingInput.push(BN256G2.G2_NEG_Y_IM);
            checkPairingInput.push(BN256G2.G2_NEG_Y_RE);
        }

        // cal H(ID1) * H(ID2) * H(ID3) ...
        {
            uint256 idx = 0;
            uint256 idy = 0;
            for (uint i = 0; i < signatures.length; i++) {
                uint256[2] memory pubKey = registry.getSignerPubkeyByAddress(
                    SignOrder[i]
                );
                (pubKey[0], pubKey[1]) = BN256G1.mulPoint(
                    [pubKey[0], pubKey[1], randoms[i]]
                );
                (idx, idy) = BN256G1.addPoint([idx, idy, pubKey[0], pubKey[1]]);
            }
            checkPairingInput.push(idx);
            checkPairingInput.push(idy);
            checkPairingInput.push(masterPubKey[1]);
            checkPairingInput.push(masterPubKey[0]);
            checkPairingInput.push(masterPubKey[3]);
            checkPairingInput.push(masterPubKey[2]);
        }

        // cal H(mi)
        {
            uint256 firstX;
            uint256 firstY;
            (firstX, firstY) = BN256G1.mulPoint(
                [
                    BN256G1.GX,
                    BN256G1.GY,
                    uint256(sha256(abi.encodePacked(message)))
                ]
            );
            uint256[2] memory first;
            (first[0], first[1]) = BN256G1.mulPoint(
                [firstX, firstY, randoms[0]]
            );
            hashPointSequence.push(first);

            for (uint i = 1; i < SignOrder.length; i++) {
                uint256 tempX;
                uint256 tempY;
                bytes32 res = sha256(
                    abi.encodePacked(
                        message,
                        signatures[i - 1][0],
                        signatures[i - 1][1]
                    )
                );
                (tempX, tempY) = BN256G1.mulPoint(
                    [BN256G1.GX, BN256G1.GY, uint256(res)]
                );
                uint256[2] memory hashPoint;
                (hashPoint[0], hashPoint[1]) = BN256G1.mulPoint(
                    [tempX, tempY, randoms[i]]
                );
                hashPointSequence.push(hashPoint);
            }
        }

        for (uint i = 0; i < signatures.length; i++) {
            checkPairingInput.push(hashPointSequence[i][0]);
            checkPairingInput.push(hashPointSequence[i][1]);
            checkPairingInput.push(setOfR[i][1]);
            checkPairingInput.push(setOfR[i][0]);
            checkPairingInput.push(setOfR[i][3]);
            checkPairingInput.push(setOfR[i][2]);
        }
        // 先虚部后实部
        require(
            BN256G1.bn256CheckPairingBatch(checkPairingInput),
            "sig verify fail"
        );

        delete randoms;
        delete hashPointSequence;
        delete checkPairingInput;
    }

    // function submitWithoutBatchVerify(
    //     uint256[4] calldata masterPubKey,
    //     uint256[2][] calldata signatures,
    //     uint256[4][] calldata setOfR
    // ) external payable {
    //     bytes32 message = registry.getMessage();
    //     address[] memory SignOrder = registry.getSignOrder();
    //     require(SignOrder.length == signatures.length, "sig nums error");

    //     {
    //         uint256 firstX;
    //         uint256 firstY;
    //         (firstX, firstY) = BN256G1.mulPoint(
    //             [
    //                 BN256G1.GX,
    //                 BN256G1.GY,
    //                 uint256(sha256(abi.encodePacked(message)))
    //             ]
    //         );

    //         hashPointSequence.push([firstX, firstY]);

    //         for (uint i = 1; i < SignOrder.length; i++) {
    //             uint256 tempX;
    //             uint256 tempY;
    //             bytes32 res = sha256(
    //                 abi.encodePacked(
    //                     message,
    //                     signatures[i - 1][0],
    //                     signatures[i - 1][1]
    //                 )
    //             );
    //             (tempX, tempY) = BN256G1.mulPoint(
    //                 [BN256G1.GX, BN256G1.GY, uint256(res)]
    //             );
    //             hashPointSequence.push([tempX, tempY]);
    //         }
    //     }

    //     for (uint i = 0; i < SignOrder.length; i++) {
    //         uint256[2] memory pubKey = registry.getSignerPubkeyByAddress(
    //             SignOrder[i]
    //         );
    //         checkPairingInput.push(signatures[i][0]);
    //         checkPairingInput.push(signatures[i][1]);
    //         checkPairingInput.push(BN256G2.G2_NEG_X_IM);
    //         checkPairingInput.push(BN256G2.G2_NEG_X_RE);
    //         checkPairingInput.push(BN256G2.G2_NEG_Y_IM);
    //         checkPairingInput.push(BN256G2.G2_NEG_Y_RE);

    //         checkPairingInput.push(pubKey[0]);
    //         checkPairingInput.push(pubKey[1]);
    //         checkPairingInput.push(masterPubKey[1]);
    //         checkPairingInput.push(masterPubKey[0]);
    //         checkPairingInput.push(masterPubKey[3]);
    //         checkPairingInput.push(masterPubKey[2]);

    //         checkPairingInput.push(hashPointSequence[i][0]);
    //         checkPairingInput.push(hashPointSequence[i][1]);
    //         checkPairingInput.push(setOfR[i][1]);
    //         checkPairingInput.push(setOfR[i][0]);
    //         checkPairingInput.push(setOfR[i][3]);
    //         checkPairingInput.push(setOfR[i][2]);

    //         require(
    //             BN256G1.bn256CheckPairingBatch(checkPairingInput),
    //             "sig verify fail"
    //         );
    //         delete checkPairingInput;
    //     }

    //     delete hashPointSequence;
    // }
}
