import { Controller, Get, Query } from '@nestjs/common';
import { AppService } from './app.service';

@Controller()
export class AppController {
  constructor(private readonly appService: AppService) { }

  @Get(":pairName")
  getHello(@Query() pairName): string {
    return this.appService.getHello(pairName);
  }
}
