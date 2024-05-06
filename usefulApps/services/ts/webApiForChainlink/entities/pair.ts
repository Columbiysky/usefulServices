import { Column, Entity, EntityOptions, PrimaryColumn } from "typeorm";

@Entity(<EntityOptions>{
    name: "Pair",
    schema: "finance"
})
export class Pair {
    constructor(name: string, price: number) {
        this.Name = name;
        this.Price = price;
    }

    @PrimaryColumn("int8")
    Id: number

    @Column("text")
    Name: string

    @Column("float8")
    Price: number
}