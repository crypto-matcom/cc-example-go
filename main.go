package main

import (
	"github.com/crypto-matcom/cc-example/chaincode"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"

	"log"
)

func main() {
	// simpleContract := new(SimpleContract)
	assetChaincode, err := contractapi.NewChaincode(&chaincode.SimpleContract{})
	if err != nil {
		log.Panicf("Error creating meiner chaincode: %v", err)
	}

	if err := assetChaincode.Start(); err != nil {
		log.Panicf("Error starting meiner chaincode: %v", err)
	}
}
