// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

contract Registry {
    struct Node {
        address addr; // eth address
        string ipAddr; // IP
        uint256[2] pubKey; // H(identity), H: string -> G1 point
        uint256 index;
        uint256 lambda;
        uint256 stake;
    }

    mapping(address => Node) private NodeMap;

    address[] private NodeArr;

    address private aggregator;

    uint256 private constant MIN_STAKE = 1 ether;

    event RegisterOracleNode(address indexed sender);

    // --------------------------------------------------------------------------------------------------------------------------------------------------------

    function register(
        string calldata ipAddr,
        uint256[2] calldata pubKey
    ) external payable {
        require(NodeMap[msg.sender].addr != msg.sender, "ALREADY");
        require(msg.value >= MIN_STAKE, "STAKE TOO LOW");
        Node storage node = NodeMap[msg.sender];
        node.addr = msg.sender;
        node.ipAddr = ipAddr;
        node.pubKey = pubKey;
        node.stake = msg.value;
        if(NodeArr.length == 0){
            aggregator = msg.sender;
        }
        NodeArr.push(msg.sender);
        node.index = NodeArr.length;
        emit RegisterOracleNode(msg.sender);
    }

    function unregister() external payable {
        require(NodeMap[msg.sender].addr == msg.sender, "HAVEN'T REGISTER");
        payable(msg.sender).transfer(NodeMap[msg.sender].stake);
        uint index = NodeMap[msg.sender].index;
        for (; index < NodeArr.length; index++) {
            NodeMap[msg.sender].index = NodeMap[msg.sender].index - 1;
            NodeArr[index - 1] = NodeArr[index];
        }
        NodeArr.pop();
        delete NodeMap[msg.sender];
    }

    function minStake() public pure returns (uint256){
        return MIN_STAKE;
    }

    function getNodeByAddress(address addr) public view returns (Node memory) {
        return NodeMap[addr];
    }

    function getIndex(address addr) public view returns (uint256) {
        return NodeMap[addr].index;
    }

    function getLambda(address addr) public view returns (uint256) {
        return NodeMap[addr].lambda;
    }

    function setLambda(address addr, uint256 lambda) public {
        NodeMap[addr].lambda = lambda;
    }

    function countOracleNodes() public view returns (uint256) {
        return NodeArr.length;
    }

    function findOracleNodeByIndex(
        uint256 _index
    ) public view returns (Node memory) {
        require(_index >= 0 && _index < NodeArr.length, "not found");
        return NodeMap[NodeArr[_index]];
    }

    function isAggregator(address addr) public view returns (bool){
        return aggregator == addr;
    }

    function getAggregator() public view returns (address){
        return aggregator;
    }

}
