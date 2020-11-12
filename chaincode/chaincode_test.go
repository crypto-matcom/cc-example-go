package chaincode_test

import (
	"fmt"

	"github.com/crypto-matcom/cc-example/chaincode"
	"github.com/crypto-matcom/cc-example/chaincode/mocks"
	"github.com/crypto-matcom/cc-example/testcerts"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

const (
	myOrg1Msp      = "org1MSP"
	myOrg1Clientid = "myOrg1Userid"
)

var _ = ginkgo.Describe("Chaincode", func() {

	// load certificate with attributes and abac
	certBytes, _ := testcerts.Certificates[3].CertBytes()

	var (
		transactionContext *mocks.TransactionContext
		chaincodeStub      *mocks.ChaincodeStub
		clientIdentity     *mocks.ClientIdentity
		err                error
	)

	ginkgo.When("at the beginning", func() {
		// mocks as Organization 1
		transactionContext, chaincodeStub, clientIdentity, err = prepMocksAsOrg1()
		err = setCreator(transactionContext, chaincodeStub, clientIdentity, myOrg1Msp, certBytes)
		gomega.Expect(err).NotTo(gomega.HaveOccurred())
	})

	ginkgo.Describe("Asset", func() {
		// Create chaincode mock
		var assetTransfer = chaincode.SimpleContract{}

		ginkgo.Context("Init ledger", func() {
			ginkgo.It("Successful", func() {
				err = assetTransfer.InitLedger(transactionContext)
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
			})

			ginkgo.It("Unsuccessful", func() {
				chaincodeStub.PutStateReturns(fmt.Errorf("failed inserting key"))
				err = assetTransfer.InitLedger(transactionContext)
				gomega.Expect(err.Error()).To(gomega.BeIdenticalTo("failed to put to world state. failed inserting key"))
			})
		})

		ginkgo.Context("Create asset", func() {

			ginkgo.It("Successful", func() {
				chaincodeStub.PutStateReturns(nil)
				err = assetTransfer.CreateAsset(transactionContext, "", "", 0, "", 0)
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
			})

			ginkgo.It("Unsuccessful", func() {
				chaincodeStub.GetStateReturns([]byte{}, nil)
				err = assetTransfer.CreateAsset(transactionContext, "cda6498a-235d-4f7e-ae19-661d41bc15c1", "", 0, "", 0)
				gomega.Expect(err.Error()).To(gomega.BeIdenticalTo("the asset cda6498a-235d-4f7e-ae19-661d41bc15c1 already exists"))

				chaincodeStub.GetStateReturns(nil, fmt.Errorf("unable to retrieve asset"))
				err = assetTransfer.CreateAsset(transactionContext, "cda6498a-235d-4f7e-ae19-661d41bc15c1", "", 0, "", 0)
				gomega.Expect(err.Error()).To(gomega.BeIdenticalTo("failed to read from world state: unable to retrieve asset"))
			})
		})
	})

})
