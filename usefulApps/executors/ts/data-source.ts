import fs from "fs";
import { DataSource } from "typeorm";
import { Pair } from "./entities/pair";

const dbConfig = JSON.parse(fs.readFileSync('./configs/dbConfig.json', 'utf8'));
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
    // schema: "finance"
})