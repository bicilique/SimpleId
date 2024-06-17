// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.18;

contract SecureDataTransfer {

    struct DataPackage {
        string encryptedData;
        string dataHash;
        string signedHash;
        uint256 timeCreated;
    }
    
    DataPackage private dataPackages;
    address private senderAddress;
    address private receiverAddress;
    bool private alreadySend;

    constructor() {
        senderAddress = msg.sender;
        alreadySend = false;
    }
    
    event DataSent(address indexed sender, address indexed receiver, uint256 timeCreated);

    modifier onlySenderOrReceiver() {
        require(msg.sender == senderAddress || msg.sender == receiverAddress,"Only the sender or receiver can access this data.");
        _;
    }

    modifier notSent() {
        require(!alreadySend, "Identity has already been created.");
        _;
    }

    function sendData(address _receiver, string memory _encryptedData, string memory _dataHash, string memory _signedHash) external notSent {
        require(_receiver != address(0), "Invalid receiver address.");

        DataPackage memory dataPackage = DataPackage({
            encryptedData: _encryptedData,
            dataHash: _dataHash,
            signedHash: _signedHash,
            timeCreated: block.timestamp
        });

        dataPackages = dataPackage;
        receiverAddress = _receiver;
        alreadySend = true;
        emit DataSent(msg.sender, _receiver, block.timestamp);
    }

    function getData() external view onlySenderOrReceiver returns (string memory encryptedData,string memory dataHash,string memory signedHash,uint256 timeCreated) {
        DataPackage memory dataPackage = dataPackages;
        return (
            dataPackage.encryptedData,
            dataPackage.dataHash,
            dataPackage.signedHash,
            dataPackage.timeCreated
        );
    }

}