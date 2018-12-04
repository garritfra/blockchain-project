mod core;
use core::transaction;

fn main() {
    let mut blockchain = core::blockchain::new();

    let transaction = transaction::new_with_amount("Me".to_string(), "You".to_string(), 10);

    blockchain.add_pending_transaction(transaction);
    blockchain.mine_block();

    for block in blockchain.get_blocks() {
        for transaction in &block.transactions {
            println!("{}", transaction.from);
        }
    }
}