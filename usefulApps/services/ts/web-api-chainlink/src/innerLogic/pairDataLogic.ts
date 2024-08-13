import { AppDataSource } from "../core/data-source";
import { PairsEnum } from "../core/pairsEnum";
import { ChainsEnum } from './../core/chainsEnum';
import { GeneralFacade } from "./generalFacade";

AppDataSource.initialize()
    .then(() => {
        console.log("Db has been successfully connected");
    })
    .catch((error) => console.log(error));

export class PairDataLogic {
    constructor() {

    }

    private generalFacadeInstance: GeneralFacade;
    get generalFacade() {
        if (!this.generalFacadeInstance) {
            this.generalFacadeInstance = new GeneralFacade();
        }
        return this.generalFacadeInstance;
    }

    async getDataFromContracts() {
        this.generalFacade.arbitrumFacade.getDataFromContracts();
    }

    async getPairPrice(chainsEnum: ChainsEnum, pairsEnum: PairsEnum): Promise<number> {
        switch (chainsEnum) {
            case ChainsEnum.Arbitrum: {
                return this.generalFacade.arbitrumFacade.getPairPrice(pairsEnum);
            }
            default: {
                return undefined;
            }
        }
    }
}
