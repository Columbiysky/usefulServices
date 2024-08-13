import { ArbitrumFacade } from "./facades/arbitrumFacade";

export class GeneralFacade {
    constructor() {
    }

    private arbFacadeInstance: ArbitrumFacade;
    get arbitrumFacade() {
        if (!this.arbFacadeInstance)
            this.arbFacadeInstance = new ArbitrumFacade();
        return this.arbFacadeInstance;
    }
}