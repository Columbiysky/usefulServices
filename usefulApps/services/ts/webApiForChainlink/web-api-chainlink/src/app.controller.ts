import { Controller, Get, Query } from "@nestjs/common";
import { AppService } from "./app.service";

@Controller()
export class AppController {
  constructor(private readonly appService: AppService) {}

  @Get(":pairName")
  async getPairData(@Query() req) {
    const res = await this.appService.getPairData(req.pairName);
    return res;
  }
}
