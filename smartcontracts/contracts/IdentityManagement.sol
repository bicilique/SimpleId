// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.18;

contract IdentityManagement {
    
    struct Identity {
        address issuer; // Alamat pengelola atau penerbit identitas
        string name; // Nama pengguna
        string email; // Email pengguna
        string street; // Alamat jalan
        string country; // Negara
        string dateBirth; // Tanggal lahir
        string status; // Status pengguna
        bool active; // Status aktif/tidak aktif
        uint256 createdAt; // Timestamp kapan identitas dibuat
    }


    // mapping(address => Identity) public identities;

    bool private identityCreated;
    Identity private identity;
    mapping(address => bool) private  isIssuer;
    mapping(address => bool) public isStakeholder;

    event IdentityCreated(address indexed owner,string name,string email,string street,string country,string dateBirth,string status,bool active,uint256 createdAt);
    event IdentityUpdated(address indexed owner, string name, string email,string street,string country,string dateBirth,string status,bool active);
    event StakeholderAdded(address indexed stakeholder);


    constructor() {
        isIssuer[msg.sender] = true;
        isStakeholder[msg.sender] = true;
        identityCreated = false;
    }

    modifier onlyStakeholder() {
        require(isStakeholder[msg.sender], "Only the identity stakeholder can perform this action.");
        _;
    }

    modifier onlyIssuer() {
        require(isIssuer[msg.sender], "Only authorized issuers can perform this action.");
        _;
    }

    function addStakeholder(address _stakeholder) external onlyIssuer {
        require(!isStakeholder[_stakeholder], "Address is already an stakeholder.");
        isStakeholder[_stakeholder] = true;
        emit StakeholderAdded(_stakeholder);
    }

    modifier notCreated() {
        require(!identityCreated, "Identity has already been created.");
        _;
    }

    function createIdentity(
        string memory _name,
        string memory _email,
        string memory _street,
        string memory _country,
        string memory _dateBirth,
        string memory _status,
        bool _active) external notCreated {
        identity = Identity({
            issuer: msg.sender,
            name: _name,
            email: _email,
            street: _street,
            country: _country,
            dateBirth: _dateBirth,
            status: _status,
            active: _active,
            createdAt: block.timestamp // Menggunakan timestamp blok saat ini
        });
        identityCreated = true;
        emit IdentityCreated(
            msg.sender,
            _name,
            _email,
            _street,
            _country,
            _dateBirth,
            _status,
            _active,
            block.timestamp
        );
    }

    function updateIdentity(
        string memory _name,
        string memory _email,
        string memory _street,
        string memory _country,
        string memory _dateBirth,
        string memory _status,
        bool _active
    ) external onlyIssuer {
        identity.name = _name;
        identity.email = _email;
        identity.street = _street;
        identity.country = _country;
        identity.dateBirth = _dateBirth;
        identity.status = _status;
        identity.active = _active;
        emit IdentityUpdated(
            msg.sender,
            _name,
            _email,
            _street,
            _country,
            _dateBirth,
            _status,
            _active
        );
    }



    function getIdentity() external view onlyStakeholder returns (
        address issuer,
        string memory name,
        string memory email,
        string memory street,
        string memory country,
        string memory dateBirth,
        string memory status,
        bool active,
        uint256 createdAt
    ) {
        return (
            identity.issuer,
            identity.name,
            identity.email,
            identity.street,
            identity.country,
            identity.dateBirth,
            identity.status,
            identity.active,
            identity.createdAt
        );
    }
    
}
