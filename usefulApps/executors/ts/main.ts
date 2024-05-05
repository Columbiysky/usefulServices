import fs from "fs";
import WstEthToEthArbMainWrapper from "./wrappedContracts/arbMain/wstEthToEthArbMainWrapper";

const arbMainRpcLink = JSON.parse(fs.readFileSync('./configs/arbMainRpcLink.json', 'utf8'));
const wstEthToEthArbMainWrapper = new WstEthToEthArbMainWrapper(arbMainRpcLink.provider);
wstEthToEthArbMainWrapper.get().then((res) => console.log(res));