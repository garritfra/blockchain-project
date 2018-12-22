package server

import (
	"encoding/gob"
	"encoding/json"
	"log"
	"net/http"

	"github.com/garritfra/blockchain-project/core"
)

var blockchain core.Blockchain

// ServeHTTP serves a http server on a given port with format ":PORT"
func ServeHTTP(port string) {
	gob.Register(core.Block{})
	gob.Register(core.Transaction{})
	gob.Register(core.Blockchain{})

	blockchain = core.NewBlockchain()

	registerRouteHandlers()

	log.Print("Listening on port ", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}

func registerRouteHandlers() {
	http.HandleFunc("/", handleListBlocks)
	http.HandleFunc("/mine_block", handleMineBlock)
	http.HandleFunc("/pending_transactions", handleListPendingTransactions)
	http.HandleFunc("/add_transaction", handleAddTransaction)
	http.HandleFunc("/is_valid", handleIsValid)

	http.HandleFunc("/add_peers", handleAddPeers)
}

func handleError(err error, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Error: " + err.Error()))
	log.Print(err.Error())
}

func handleListBlocks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(blockchain.AsJSON())
}

func handleAddTransaction(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var receivedTransaction core.Transaction
	err := decoder.Decode(&receivedTransaction)
	if err == nil {
		log.Println("Transaction from", receivedTransaction.Sender, "to", receivedTransaction.Receiver, "received")

		go blockchain.AddTransaction(receivedTransaction)
	} else {
		handleError(err, w, r)
	}
}

func handleMineBlock(w http.ResponseWriter, r *http.Request) {
	block := blockchain.MineBlock()
	json.NewEncoder(w).Encode(block)
}

func handleListPendingTransactions(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(blockchain.PendingTransactions)
}

func handleIsValid(w http.ResponseWriter, r *http.Request) {
	valid := blockchain.IsValid()
	json.NewEncoder(w).Encode(valid)
}

// Takes an a string-slice, and adds it to the known peers
func handleAddPeers(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var receivedPeers []string
	err := decoder.Decode(&receivedPeers)

	if err == nil {
		for _, peer := range receivedPeers {
			blockchain.AddPeer(peer)
		}
	} else {
		handleError(err, w, r)
	}

	json.NewEncoder(w).Encode(blockchain.Peers)

}