// SPDX-License-Identifier: MIT
pragma solidity ~0.8.0;

contract Tran {
    event EtherReceived(address indexed sender, uint256 amount);
    event FallbackCalled(address indexed sender, uint256 amount, bytes data);
    
    address public owner;
    uint256 totalReceived;
    address[] public senders;

    constructor() {
        owner = msg.sender;
    }

    receive() external payable {
        totalReceived += msg.value;
        senders.push(msg.sender);
        emit EtherReceived(msg.sender, msg.value);
    }

    fallback() external payable {
        emit FallbackCalled(msg.sender, msg.value, "");
    }
}
