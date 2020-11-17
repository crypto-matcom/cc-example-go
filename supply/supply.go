package cc_testing

import (
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-protos-go/peer"
)

type SupplyChainCode struct {
	contractapi.Contract
}

func (s SupplyChainCode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	panic("implement me")
}

func (s SupplyChainCode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	panic("implement me")
}


