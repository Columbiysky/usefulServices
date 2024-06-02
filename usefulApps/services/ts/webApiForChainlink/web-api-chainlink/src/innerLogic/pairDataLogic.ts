import {AppDataSource} from "../core/data-source";
import * as fs from "fs";
import WstEthToEthArbMainWrapper from "./wrappedContracts/arbMain/wstEthToEthWrapper";
import EzEthToEthArbMainWrapper from "./wrappedContracts/arbMain/ezEthToEthWrapper";
import {Pair} from "../enitities/pair";
import {Repository} from "typeorm";

AppDataSource.initialize()
    .then(() => {
        console.log("Db successfully connected")
    })
    .catch((error) => console.log(error));

export class PairDataLogic {

    constructor() {
    }

    private get arbMainRpcLink(){
        return JSON.parse(fs.readFileSync('./src/configs/arbMainRpcLink.json', 'utf8'));
    }

    private get wstEthToEthArbMainWrapper(){
        return new WstEthToEthArbMainWrapper(this.arbMainRpcLink.provider);
    }

    private get ezEthToEthArbMainWrapper(){
        return new EzEthToEthArbMainWrapper(this.arbMainRpcLink.provider);
    }

    async getDataFromContracts() {
        const pairRepository = AppDataSource.getRepository(Pair);

        const wstEthToEthPrice = await this.wstEthToEthArbMainWrapper.get();
        const pair = this.pairInstance(this.wstEthToEthArbMainWrapper.pairName, wstEthToEthPrice);
        pairRepository.findOneBy({ Name: pair.Name }).then((existingPair) => {
            this.savePair(pairRepository, existingPair, pair);
            console.log("wstEthToEthArbMainWrapper saved");
        });

        const ezEthToEthPrice = await this.ezEthToEthArbMainWrapper.get();
        const pairr = this.pairInstance(this.ezEthToEthArbMainWrapper.pairName, ezEthToEthPrice);
        pairRepository.findOneBy({ Name: pairr.Name }).then((existingPair) => {
            this.savePair(pairRepository, existingPair, pairr);
            console.log("ezEthToEthArbMainWrapper saved");
        });
    }

    private pairInstance(name: string, price: number): Pair {
        return new Pair(name, price);
    }

    private async savePair(repository: Repository<Pair>, existingPair: Pair | null, newPair: Pair) {
        if (existingPair !== null && existingPair.Price === newPair.Price) {
            await repository.save(existingPair);
        } else if (existingPair !== null && existingPair.Price !== newPair.Price) {
            existingPair.Price = existingPair.Price;
            await repository.save(existingPair);
        } else {
            await repository.save(newPair);
        }
    }
}