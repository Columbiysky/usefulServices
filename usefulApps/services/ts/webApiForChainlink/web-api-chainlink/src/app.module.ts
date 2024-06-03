import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import {PairDataLogic} from "./innerLogic/pairDataLogic";

@Module({
  imports: [],
  controllers: [AppController],
  providers: [AppService, PairDataLogic],
})
export class AppModule {}
