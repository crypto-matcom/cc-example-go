package chaincode_test

import (
	"github.com/crypto-matcom/cc-example/chaincode/mocks"
	"github.com/golang/protobuf/proto" //nolint:staticcheck
	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	"github.com/hyperledger/fabric-protos-go/msp"
	"os"
	"testing"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

func TestChaincode(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Chaincode Suite")
}

// TODO@kmilo: temporarily placed here 👇🏽

func prepMocksAsOrg1() (*mocks.TransactionContext, *mocks.ChaincodeStub, *mocks.ClientIdentity, error) {
	return prepMocks(myOrg1Msp, myOrg1Clientid)
}

func prepMocks(orgMSP string, clientID string) (*mocks.TransactionContext, *mocks.ChaincodeStub, *mocks.ClientIdentity, error) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	clientIdentity := &mocks.ClientIdentity{}
	clientIdentity.GetMSPIDReturns(orgMSP, nil)
	clientIdentity.GetIDReturns(clientID, nil)

	// set matching msp ID using peer shim env variable
	err := os.Setenv("CORE_PEER_LOCALMSPID", orgMSP)
	if err != nil {
		return nil, nil, nil, err
	}

	return transactionContext, chaincodeStub, clientIdentity, nil
}

func setCreator(transactionContext *mocks.TransactionContext, chaincodeStub *mocks.ChaincodeStub, clientIdentity cid.ClientIdentity, mspID string, idbytes []byte) error {
	transactionContext.GetClientIdentityReturns(clientIdentity)

	sid := &msp.SerializedIdentity{Mspid: mspID, IdBytes: idbytes}
	buffer, err := proto.Marshal(sid)

	if err != nil {
		return err
	}

	chaincodeStub.GetCreatorReturns(buffer, nil)
	return nil
}
