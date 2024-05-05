import fs from "fs";
import WstEthToEthArbMainWrapper from "./wrappedContracts/wstEthToEthArbMainWrapper";

const tConfig = JSON.parse(fs.readFileSync('./configs/wstEthToEthArbMain/config.json', 'utf8'));
const t = new WstEthToEthArbMainWrapper(tConfig.provider, tConfig.contractAddress);
t.get().then((res) => console.log(res));