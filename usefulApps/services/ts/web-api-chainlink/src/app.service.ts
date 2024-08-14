import { Injectable } from "@nestjs/common";
import { CronJob } from 'cron';
import { ChainsEnum } from './core/chainsEnum';
import { PairsEnum } from "./core/pairsEnum";
import { PairDataLogic } from "./innerLogic/pairDataLogic";

@Injectable()
export class AppService {
    constructor(private readonly pairDataLogic: PairDataLogic) {
        const that = this;
        const job = new CronJob(
            '*/5 * * * * *', // cronTime every 5 mins
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

    async getPairData(chainName: string, pairName: string): Promise<number> {
        if (chainName in ChainsEnum && pairName in PairsEnum) {
            const price = await this.pairDataLogic
                .getPairPrice(ChainsEnum[chainName], PairsEnum[pairName])
                .then((res) => {
                    return res;
                });

            return price;
        }

        return undefined;
    }
}
