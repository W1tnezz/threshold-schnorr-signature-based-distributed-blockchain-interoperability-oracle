// SPDX-License-Identifier: MIT
pragma solidity >0.8.0;

import "./crypto/Schnorr.sol";
import "./crypto/BN256G1.sol";
import "./Registry.sol";
import "./DKG.sol";

contract Oracle {
    Registry private registry;
    DKG private dkg;

    constructor(address _registryContract, address _dkgContract) {
        registry = Registry(_registryContract);
        dkg = DKG(_dkgContract);
    }

    // 请求验证需要的钱
    uint256 public constant BASE_FEE = 0.1 ether;
    uint256 public constant AGGREGATE_FEE = 0.3 ether;

    // 保存验证结果的映射；
    mapping(bytes32 => bool) private txValidationResults;

    // indexed属性是为了方便在日志结构中查找，这个是一个事件，会存储到交易的日志中，就是类似于挖矿上链
    event ValidationRequest(
        address indexed from,
        bytes32 hash,
        bool needEnroll
    );

    modifier minFee() {
        require(
            msg.value >=
                BASE_FEE *
                    ((registry.countOracleNodes() - 1)) +
                    AGGREGATE_FEE,
            "too few"
        );
        _;
    }

    function totalFee() public view returns (uint256) {
        return
            BASE_FEE *
            ((registry.countOracleNodes() - 1)) +
            AGGREGATE_FEE;
    }

    function validateTransaction(bytes32 _message) external payable minFee {
        require(registry.countOracleNodes() >= 2, "TOO FEW NODES");
        emit ValidationRequest(msg.sender, _message, dkg.needEnroll());
    }

    function submit(bool res, bytes32 transHash, bytes memory message, uint256 signature, uint256 rx , uint256 ry) external{
        uint256[2] memory pubKey = dkg.usePubKey();

        uint256 _hash = uint256(
            sha256(abi.encodePacked(message, rx, ry, pubKey[0], pubKey[1]))
        );

        require(
            Schnorr.verify(signature, pubKey[0], pubKey[1], rx, ry, _hash),
            "sig: address doesn't match"
        );

        // txValidationResults[transHash] = res;

        //  // 给当前合约的调用者（聚合器）转账 
        // payable(msg.sender).transfer(AGGREGATE_FEE);     //此处完成给聚合器的报酬转账
        
        //  // 给所有的参与验证的验证器节点转账
        // address[] memory validators = dkg.getValidators();
        // for(uint32 i = 0 ; i < validators.length ; i++){
        //     payable(validators[i]).transfer(BASE_FEE);
        // }
    }
}
