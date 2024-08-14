import { ethers } from "ethers";

export interface IBaseWrapper {
    pairName: string;
    provider: ethers.JsonRpcProvider;
    contractAddress: string;
    GetPrice(): Promise<number>;
}
