import { Controller, Get, Query } from '@nestjs/common';
import { AppService } from './app.service';

@Controller()
export class AppController {
  constructor(private readonly appService: AppService) { }

  @Get(":pairName")
  getPairData(@Query() req): string {
    const pairData = this.appService.getPairData(req.pairName);
    return pairData;
  }
}
