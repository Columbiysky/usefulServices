import * as fs from "fs";
import { Pair } from "src/enitities/pair";
import { DataSource } from "typeorm";

const dbConfig = JSON.parse(
  fs.readFileSync("./src/configs/dbConfig.json", "utf8"),
);
export const AppDataSource = new DataSource({
  type: "postgres",
  host: dbConfig.host,
  port: dbConfig.port,
  username: dbConfig.username,
  password: dbConfig.password,
  database: dbConfig.database,
  synchronize: false,
  logging: false,
  entities: [Pair],
});
