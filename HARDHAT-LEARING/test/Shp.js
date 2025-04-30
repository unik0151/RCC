const { expect } = require("chai");

const hre = require("hardhat");

// async function(){

// }  == async ()=>{}
describe("shp" , function (){
    let shpContract;

    before(async () => { 
        // ⽣成合约实例并且复⽤ 
        shpContract = await hre.ethers.deployContract("Shp", []);
     });
    it("should return the status noload", async function(){
        expect(await shpContract.getStatus()).to.equal("NOLOAD");
    });

    it("should return the status load", async function(){
        await shpContract.loadLogFunc();
        expect(await shpContract.getStatus()).to.equal("LOAD");
    });
    
    it("should return event log " , async ()=>{
        await expect(shpContract.DeleiveredLogFunc())
        .to.emit(shpContract,"DeleiveredLog")
        .withArgs("shop has been delivered");
    });

})