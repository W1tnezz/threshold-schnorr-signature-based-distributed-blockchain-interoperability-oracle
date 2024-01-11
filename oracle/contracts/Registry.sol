// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

contract Registry {
    struct Signer {
        address addr; // eth address
        string ipAddr; // IP
        string identity; // identity；
        uint256[2] pubKey; // H(identity), H: string -> G1 point
    }

    mapping(address => Signer) private SignerMap;

    address[] private SignerArr;

    address[] private SignOrder;

    bytes32 private message;


    enum SignType {
        SAKAI,
        IBSAS,
        NOTBATCH
    }

    event Sign(SignType typ, bytes32 message, address[] signOrder);

    // --------------------------------------------------------------------------------------------------------------------------------------------------------

    function register(
        string calldata ipAddr,
        string calldata identity,
        uint256[2] calldata pubKey
    ) external payable {
        require(SignerMap[msg.sender].addr != msg.sender, "already");
        require(bytes(identity).length != 0, "empty identity");
        Signer storage signer = SignerMap[msg.sender];
        signer.addr = msg.sender;
        signer.ipAddr = ipAddr;
        signer.identity = identity;
        signer.pubKey = pubKey;
        SignerArr.push(msg.sender);
    }

    function getSignerByAddress(
        address addr
    ) public view returns (Signer memory) {
        return SignerMap[addr];
    }

    function getSignerPubkeyByAddress(
        address addr
    ) public view returns (uint256[2] memory) {
        return SignerMap[addr].pubKey;
    }

    function requestSign(
        SignType typ,
        bytes32 _message,
        uint[] calldata signOrder
    ) external payable {
        delete SignOrder;
        for (uint i = 0; i < signOrder.length; i++) {
            SignOrder.push(SignerArr[i]);
        }
        message = _message;
        emit Sign(typ, message, SignOrder);
    }

    function getMessage() public view returns (bytes32) {
        return message;
    }

    function getSignOrder() public view returns (address[] memory) {
        return SignOrder;
    }
}
