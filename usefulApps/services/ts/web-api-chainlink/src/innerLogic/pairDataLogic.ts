import { AppDataSource } from "../core/data-source";
import { PairsEnum } from "../core/pairsEnum";
import { GeneralFacade } from "./generalFacade";

AppDataSource.initialize()
    .then(() => {
        console.log("Db has been successfully connected");
    })
    .catch((error) => console.log(error));

export class PairDataLogic {
    constructor() {

    }

    get generalFacade() {
        return new GeneralFacade();
    }

    async getDataFromContracts() {
        this.generalFacade.arbitrumFacade.getDataFromContracts();
    }

    async getPairPrice(pairsEnum: PairsEnum): Promise<number> {
        return this.generalFacade.arbitrumFacade.getPairPrice(pairsEnum);
    }
}
