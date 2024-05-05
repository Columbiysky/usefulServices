import { ethers } from "ethers";

export interface IBaseWrapper {
    provider: ethers.JsonRpcProvider,
    contractAddress: string,
    get(): Promise<number>;
}