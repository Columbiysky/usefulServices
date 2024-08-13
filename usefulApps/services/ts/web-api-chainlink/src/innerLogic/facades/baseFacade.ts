import { Pair } from "src/enitities/pair";

export class BaseFacade {

    protected pairInstance(name: string, price: number, chainName: string): Pair {
        return new Pair(name, price, chainName);
    }
}