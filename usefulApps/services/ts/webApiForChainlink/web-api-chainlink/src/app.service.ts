import { Injectable } from "@nestjs/common";
import { PairDataLogic } from "./innerLogic/pairDataLogic";
import { PairsEnum } from "./core/pairsEnum";

@Injectable()
export class AppService {
  constructor(private readonly pairDataLogic: PairDataLogic) {
    this.pairDataLogic
      .getDataFromContracts()
      .then(() => console.log("Smart contracts data has been updated"));
  }

  async getPairData(pairName: string): Promise<number> {
    if (pairName in PairsEnum) {
      const price = await this.pairDataLogic
        .getPairPrice(PairsEnum[pairName])
        .then((res) => {
          return res;
        });

      return price;
    }

    return undefined;
  }
}
