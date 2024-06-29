import { Injectable } from "@nestjs/common";
import { CronJob } from 'cron';
import { PairsEnum } from "./core/pairsEnum";
import { PairDataLogic } from "./innerLogic/pairDataLogic";

@Injectable()
export class AppService {
  constructor(private readonly pairDataLogic: PairDataLogic) {
    const that = this;
    const job = new CronJob(
      '0 22 * * * *', // cronTime
      function () {
        that.pairDataLogic
          .getDataFromContracts()
          .then(() => {
            console.log("Smart contracts data has been updated from cron"); 
            console.log(new Date().toISOString()); 
          });
      }, // onTick
      null, // onComplete
      true, // start
      'Europe/Moscow' // timeZone
    );

    job.start();
    
    this.pairDataLogic
      .getDataFromContracts()
      .then(() => console.log("Smart contracts data has been updated from start"));
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
