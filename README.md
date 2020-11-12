# cc-example
Hyperledger Fabric Proof of Concept (PoC) Project with Golang


## testcerts folder
Folder with fake certificates for mock up
- testcert.go (to load test certificates found in the testcerts folder)
```

import (
	"github.com/crypto-matcom/cc-example/testcerts"
)

// load certificate with attributes for abac
certBytes, _ := testcerts.Certificates[3].CertBytes()
```

## counterfeiter
Use counterfeiter to generate directives.

Counterfeiter allows you to simply generate test doubles for a given interface.

see https://github.com/maxbrunsfeld/counterfeiter

## Behavior-Driven Development (BDD)
We use [Ginkgo](https://github.com/onsi/ginkgo) which enables the BDD.
Paired with the [Gomega](https://github.com/onsi/gomega) matcher library.