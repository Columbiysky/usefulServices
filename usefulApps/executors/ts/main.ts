import { WstEthToEthArbMainWrapper } from "./wrappedContracts/wstEthToEthArbMainWrapper.js";

const t = new WstEthToEthArbMainWrapper();
t.get().then((res) => console.log(res));