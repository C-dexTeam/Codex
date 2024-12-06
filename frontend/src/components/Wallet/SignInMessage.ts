<<<<<<< HEAD
// import bs58 from "bs58";
import { base58_to_binary } from "base58-js";
=======
import { base58_to_binary } from 'base58-js';
>>>>>>> 3a9b9de425f75269bdd7cb465063b3ea01be1d75
import nacl from "tweetnacl";
type SignMessage = {
    domain: string;
    publicKey: string;
    nonce: string;
    statement: string;
};

export class SigninMessage {
    domain: any;
    publicKey: any;
<<<<<<< HEAD
    nonce: any;
    statement: any;

    constructor({ domain, publicKey, nonce, statement }: SignMessage) {
        this.domain = domain;
        this.publicKey = publicKey;
        this.nonce = nonce;
=======
    statement: any;

    constructor({ domain, publicKey, statement }: SignMessage) {
        this.domain = domain;
        this.publicKey = publicKey;
>>>>>>> 3a9b9de425f75269bdd7cb465063b3ea01be1d75
        this.statement = statement;
    }

    prepare() {
<<<<<<< HEAD
        return `${this.statement}${this.nonce}`;
=======
        return this.statement;
>>>>>>> 3a9b9de425f75269bdd7cb465063b3ea01be1d75
    }

    async validate(signature: string) {
        const msg = this.prepare();
        const signatureUint8 = base58_to_binary(signature);
        const msgUint8 = new TextEncoder().encode(msg);
        const pubKeyUint8 = base58_to_binary(this.publicKey);

        return nacl.sign.detached.verify(msgUint8, signatureUint8, pubKeyUint8);
    }
}