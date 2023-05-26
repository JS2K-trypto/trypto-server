// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/Counters.sol";
import "@chainlink/contracts/src/v0.8/AutomationCompatible.sol";




contract TryptoForRemix is ERC721, ERC721URIStorage, Ownable, AutomationCompatibleInterface {
    using Counters for Counters.Counter;

    Counters.Counter private _tokenIdCounter;

    // Counter : mapping for badge upgrades
    
    // badgeLevel 
    // ex ) 5(tokenId) : 0(level)
    // ex ) 8(tokenId) : 2(level)
    mapping(uint => uint) public badgeLevel;

    // pendingUpgrade is for counting badgeLevel
    uint pendingUpgrade;

    // Metadata information for each stage of the NFT on IPFS.
    string[] IpfsUri = [
        "https://ipfs-2.thirdwebcdn.com/ipfs/QmPwAZ6xZkVsTYjtnCmuEK1AfAT3TVvBHQ6N32w76A51yn/bronze.json",
        "https://ipfs-2.thirdwebcdn.com/ipfs/QmPwAZ6xZkVsTYjtnCmuEK1AfAT3TVvBHQ6N32w76A51yn/silver.json",
        "https://ipfs-2.thirdwebcdn.com/ipfs/QmPwAZ6xZkVsTYjtnCmuEK1AfAT3TVvBHQ6N32w76A51yn/gold.json"
        
    ];

    uint lastTimeStamp;
    uint interval;


     

    constructor(uint _interval) ERC721("Trypto", "TRT") {
        interval = _interval;
        lastTimeStamp = block.timestamp;
    }

    // Mint NFT(Badge)
    function safeMint(address to, string memory uri) public onlyOwner {
        uint256 tokenId = _tokenIdCounter.current();
        _tokenIdCounter.increment();
        _safeMint(to, tokenId);
        _setTokenURI(tokenId, uri);
        
    }


    // (ONLY FOR TEST, WILL BE DELETED LATER) upgrade nft by Change tokenURI
    function upgradeBadge(uint _tokenId, string memory _uri) public onlyOwner {
        _setTokenURI(_tokenId, _uri);

    }

    // get All nfts of User
    function getNftsOf(address _address) public view returns (string[] memory) {
        uint tokenCounts = _tokenIdCounter.current();
        string[] memory badges = new string[](tokenCounts);
        uint index = 0;
        for (uint i=0; i<tokenCounts; i++) {
            if (ownerOf(i) == _address) {
                string memory nftInfo = tokenURI(i);
                badges[index] = nftInfo;
                index++;
            }     
        }  
        return badges;
    }

    /// increase badgeLevel if user visit the country again and click the button
    function increasebadgeLevel(uint _tokenId) public onlyOwner {
        require(badgeLevel[_tokenId] < 2);
        badgeLevel[_tokenId]++;
        pendingUpgrade++;
    }

    
    // Automation calls this functions every interval,
    // upgrade every badges that badeLevels are 1 or 2  
    function upgrade() public {
        uint nftcounts = _tokenIdCounter.current();
        for(uint i=0;i<nftcounts;i++){
            if(badgeLevel[i] == 1) {
                _setTokenURI(i, IpfsUri[1]);
            } else if (badgeLevel[i] == 2) {
                _setTokenURI(i, IpfsUri[2]);
            }
            
        }
    }



    function checkUpkeep(
        bytes calldata /* checkData */
    )
        external
        view
        returns (
            bool upkeepNeeded,
            bytes memory /* performData */
        )
    {
        upkeepNeeded = (block.timestamp - lastTimeStamp) > interval;
        // We don't use the checkData in this example. The checkData is defined when the Upkeep was registered.
    }

    function performUpkeep(
        bytes calldata /* performData */
    ) external {
        //We highly recommend revalidating the upkeep in the performUpkeep function
        if ((block.timestamp - lastTimeStamp) > interval) {
            lastTimeStamp = block.timestamp;
            upgrade();
        }
        // We don't use the performData in this example. The performData is generated by the Keeper's call to your checkUpkeep function
    }




    // The following functions are overrides required by Solidity.

    function _burn(uint256 tokenId) internal override(ERC721, ERC721URIStorage) {
        super._burn(tokenId);
    }

    // shows NFT's Metadata  
    function tokenURI(uint256 tokenId)
        public
        view
        override(ERC721, ERC721URIStorage)
        returns (string memory)
    {
        return super.tokenURI(tokenId);
    }
}