export interface ICustomObject {
    id: number;
    name: string;
}

export default class CustomObject implements ICustomObject {
    constructor(id_: number, name_: string) {
        this.id = id_;
        this.name = name_;
    }
    id: number;
    name: string;

}