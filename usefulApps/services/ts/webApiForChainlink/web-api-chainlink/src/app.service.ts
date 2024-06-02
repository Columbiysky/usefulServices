import { Injectable } from '@nestjs/common';
import {PairDataLogic} from "./innerLogic/pairDataLogic";

@Injectable()
export class AppService {
  getPairData(param: string): string {
    const pdl  = new PairDataLogic();
    pdl.getDataFromContracts()
    return "Hello " + param;
  }
}
