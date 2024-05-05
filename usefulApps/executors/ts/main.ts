import fs from "fs";
import EzEthToEthArbMainWrapper from "./wrappedContracts/arbMain/ezEthToEthWrapper";
import WstEthToEthArbMainWrapper from "./wrappedContracts/arbMain/wstEthToEthWrapper";

const arbMainRpcLink = JSON.parse(fs.readFileSync('./configs/arbMainRpcLink.json', 'utf8'));
const wstEthToEthArbMainWrapper = new WstEthToEthArbMainWrapper(arbMainRpcLink.provider);
const ezEthToEthArbMainWrapper = new EzEthToEthArbMainWrapper(arbMainRpcLink.provider);
wstEthToEthArbMainWrapper.get().then((res) => console.log("wstETH\\ETH = " + res));
ezEthToEthArbMainWrapper.get().then((res) => console.log("ezETH\\ETH = " + res));