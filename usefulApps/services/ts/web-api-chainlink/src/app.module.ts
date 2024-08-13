import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { GeneralFacade } from './innerLogic/generalFacade';
import { PairDataLogic } from "./innerLogic/pairDataLogic";

@Module({
  imports: [],
  controllers: [AppController],
  providers: [AppService, PairDataLogic, GeneralFacade],
})
export class AppModule { }
