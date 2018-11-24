import { SHA256 } from "crypto-js";
import Transaction from "./Transaction";

export default class Block {
  id: any;
  timestamp: Date;
  data: any[];
  previousHash: any;
  hash: any;
  /**
   * Construct a Block object
   * @constructor
   * @param {number} id
   * @param {string} previousHash
   */
  constructor(id, previousHash) {
    this.id = id;
    this.timestamp = new Date();
    this.data = [];
    this.previousHash = previousHash;
    this.hash = this.calculateHash();
  }

  /**
   * Concatenates the timestamp, previous hash and transaction data of a block and
   * calculates the according SHA256 Hash
   * @returns hash
   */
  calculateHash() {
    let hashStr = this.timestamp + this.previousHash + this.data;
    let hash = SHA256(hashStr).toString();
    return hash;
  }

  /**
   * Creates the first block in a blockchain.
   * It has a stub transaction with an amount of 0 coins.
   *
   * Called automatically upon blockchain creation
   *
   * @returns genesisBlock
   */
  static createGenesisBlock() {
    let date = new Date();
    let genesisTransaction = new Transaction("Genesis", "Genesis", 0);
    let genesisBlock = new Block(0, "0");
    genesisBlock.addTransaction(genesisTransaction);
    return genesisBlock;
  }

  /**
   * Pushes a transaction to the Block
   * @param {Transaction} transaction
   */
  addTransaction(transaction) {
    this.data.push(transaction);
  }
}
