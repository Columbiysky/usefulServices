import { Controller, Get, Query } from "@nestjs/common";
import { AppService } from "./app.service";

@Controller()
export class AppController {
    constructor(private readonly appService: AppService) { }

    @Get("/getPairData")
    async getPairData(@Query('chain') chainName: string, @Query('pairName') pairName: string) {
        const res = await this.appService.getPairData(chainName, pairName);
        return res ?? null;
    }
}
