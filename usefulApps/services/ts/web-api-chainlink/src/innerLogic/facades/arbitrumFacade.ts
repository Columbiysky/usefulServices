import { Injectable } from "@nestjs/common";
import * as fs from "fs";
import { AppDataSource } from "src/core/data-source";
import { PairsEnum } from "src/core/pairsEnum";
import { Pair } from "src/enitities/pair";
import { Repository } from "typeorm";
import EzEthToEthArbMainWrapper from "../wrappedContracts/arbMain/ezEthToEthWrapper";
import WstEthToEthArbMainWrapper from "../wrappedContracts/arbMain/wstEthToEthWrapper";
import { BaseFacade } from "./baseFacade";

@Injectable()
export class ArbitrumFacade extends BaseFacade {
    private get arbMainRpcLink() {
        return JSON.parse(
            fs.readFileSync("./src/configs/arbMainRpcLink.json", "utf8"),
        );
    }

    private wstEthToEthArbMainWrapperInstance: WstEthToEthArbMainWrapper;
    private ezEthToEthArbMainWrapperInstance: EzEthToEthArbMainWrapper;
    private get wstEthToEthArbMainWrapper() {
        if (!this.wstEthToEthArbMainWrapperInstance) {
            this.wstEthToEthArbMainWrapperInstance = new WstEthToEthArbMainWrapper(this.arbMainRpcLink.provider);
        }

        return this.wstEthToEthArbMainWrapperInstance;
    }

    private get ezEthToEthArbMainWrapper() {
        if (!this.ezEthToEthArbMainWrapperInstance) {
            this.ezEthToEthArbMainWrapperInstance = new EzEthToEthArbMainWrapper(this.arbMainRpcLink.provider);
        }

        return this.ezEthToEthArbMainWrapperInstance;
    }

    getDataFromContracts() {
        const pairRepository = AppDataSource.getRepository(Pair);
        this.getWstEthToEthPrice(pairRepository);
        this.getEzEthToEthPrice(pairRepository);
    }

    async getPairPrice(pairsEnum: PairsEnum): Promise<number> {
        const pairRepository = AppDataSource.getRepository(Pair);

        switch (pairsEnum) {
            case PairsEnum.ezETHETH: {
                const price = await pairRepository
                    .findOneBy({ Name: PairsEnum.ezETHETH })
                    .then((existingPair) => {
                        return (existingPair as Pair).Price;
                    });

                return price;
            }
            case PairsEnum.wstETHETH: {
                const price = await pairRepository
                    .findOneBy({ Name: PairsEnum.wstETHETH })
                    .then((existingPair) => {
                        return (existingPair as Pair).Price;
                    });
                return price;
            }
            default: {
                return undefined;
            }
        }
    }

    private async savePair(
        repository: Repository<Pair>,
        existingPair: Pair | null,
        newPair: Pair,
    ) {
        if (existingPair !== null && existingPair.Price === newPair.Price) {
            await repository.save(existingPair);
        } else if (existingPair !== null && existingPair.Price !== newPair.Price) {
            existingPair.Price = newPair.Price;
            await repository.save(existingPair);
        } else {
            await repository.save(newPair);
        }
    }

    private async getWstEthToEthPrice(pairRepository: Repository<Pair>) {
        const wstEthToEthPrice = await this.wstEthToEthArbMainWrapper.get();
        const pair = this.pairInstance(
            this.wstEthToEthArbMainWrapper.pairName,
            wstEthToEthPrice,
        );
        pairRepository.findOneBy({ Name: pair.Name }).then((existingPair) => {
            this.savePair(pairRepository, existingPair, pair);
        });
    }

    private async getEzEthToEthPrice(pairRepository: Repository<Pair>) {
        const ezEthToEthPrice = await this.ezEthToEthArbMainWrapper.get();
        const pair = this.pairInstance(
            this.ezEthToEthArbMainWrapper.pairName,
            ezEthToEthPrice,
        );
        pairRepository.findOneBy({ Name: pair.Name }).then((existingPair) => {
            this.savePair(pairRepository, existingPair, pair);
        });
    }
}