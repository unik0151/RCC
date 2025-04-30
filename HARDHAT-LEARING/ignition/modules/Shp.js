const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules"); 
module.exports = buildModule("ShpModule", (m) => { 
const shp = m.contract("Shp", []); 
m.call(shp, "getStatus", []);
return { shp }; 
}); 