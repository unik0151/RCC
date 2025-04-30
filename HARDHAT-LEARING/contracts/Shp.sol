// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8;

contract Shp {

    enum ShpStatus {
        NOLOAD,
        LOAD,
        DELIVERED
    }
    ShpStatus private shopStatus ;
    constructor() public {
        shopStatus = ShpStatus.NOLOAD;
    }

    event LoadLog(string desc);

    event DeleiveredLog(string desc);

    function loadLogFunc() public {
        shopStatus = ShpStatus.LOAD;
        emit LoadLog("shop has been loaded");
    }

    function DeleiveredLogFunc() public {
        shopStatus = ShpStatus.DELIVERED;
        emit DeleiveredLog("shop has been delivered");
    }

    // modifier checkStatus {
    //     require(shopStatus , "not");
    //     _;
    // }
    function getStatus()public view returns (string memory){
        ShpStatus sp = shopStatus;
        return getStatus(sp);
    }
    function getStatus(ShpStatus sp) internal pure returns(string memory txt){
        if (sp == ShpStatus.NOLOAD) {
            return "NOLOAD";
        }
  
        if (sp == ShpStatus.LOAD) {
            return "LOAD";
        }
           
        if (sp == ShpStatus.DELIVERED) {
            return "DELIVERED";
        }
    }



}