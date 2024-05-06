import fs from "fs";
import "reflect-metadata";
import { Repository } from "typeorm";
import { AppDataSource } from "./data-source";
import { Pair } from "./entities/pair";
import EzEthToEthArbMainWrapper from "./wrappedContracts/arbMain/ezEthToEthWrapper";
import WstEthToEthArbMainWrapper from "./wrappedContracts/arbMain/wstEthToEthWrapper";

AppDataSource.initialize()
    .then(() => { console.log("Db successfully connected") })
    .catch((error) => console.log(error));

const arbMainRpcLink = JSON.parse(fs.readFileSync('./configs/arbMainRpcLink.json', 'utf8'));
const wstEthToEthArbMainWrapper = new WstEthToEthArbMainWrapper(arbMainRpcLink.provider);
const ezEthToEthArbMainWrapper = new EzEthToEthArbMainWrapper(arbMainRpcLink.provider);

getDataFromContracts();
setInterval(
    getDataFromContracts,
    1000 * 60 * 60 // 1 hour in milliseconds
)

async function getDataFromContracts() {
    const pairRepository = AppDataSource.getRepository(Pair);

    const wstEthToEthPrice = await wstEthToEthArbMainWrapper.get();
    const pair = pairInstance(wstEthToEthArbMainWrapper.pairName, wstEthToEthPrice);
    pairRepository.findOneBy({ Name: pair.Name }).then((existingPair) => {
        savePair(pairRepository, existingPair, pair);
        console.log("wstEthToEthArbMainWrapper saved");
    });

    const ezEthToEthPrice = await ezEthToEthArbMainWrapper.get();
    const pairr = pairInstance(ezEthToEthArbMainWrapper.pairName, ezEthToEthPrice);
    pairRepository.findOneBy({ Name: pairr.Name }).then((existingPair) => {
        savePair(pairRepository, existingPair, pairr);
        console.log("ezEthToEthArbMainWrapper saved");
    });
}

function pairInstance(name: string, price: number): Pair {
    return new Pair(name, price);
}

async function savePair(repository: Repository<Pair>, existingPair: Pair | null, newPair: Pair) {
    if (existingPair !== null && existingPair.Price === newPair.Price) {
        await repository.save(existingPair);
    } else if (existingPair !== null && existingPair.Price !== newPair.Price) {
        existingPair.Price = existingPair.Price;
        await repository.save(existingPair);
    } else {
        await repository.save(newPair);
    }
}