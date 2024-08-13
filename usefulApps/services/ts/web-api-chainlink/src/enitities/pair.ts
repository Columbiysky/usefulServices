import { Column, Entity, EntityOptions, PrimaryColumn } from "typeorm";

@Entity(<EntityOptions>{
  name: "Pair",
  schema: "finance",
})
export class Pair {
  constructor(name: string, price: number, chainName: string, date?: string) {
    this.Name = name;
    this.Price = price;
    this.ChainName = chainName;
    this.LastTimeUpdate = date ?? new Date().toISOString();
  }

  @PrimaryColumn("int8")
  Id: number;

  @Column("text")
  Name: string;

  @Column("float8")
  Price: number;

  @Column("timestamp with time zone")
  LastTimeUpdate: string;

  @Column("text")
  ChainName: string;
}
