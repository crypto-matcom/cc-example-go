package testcerts

import (
	"io/ioutil"
	"path"
	"runtime"
)

type (
	FileReader func(filename string) ([]byte, error)

	// Cert certificate data for testing
	Cert struct {
		CertFilename string
		PKeyFilename string
		readFile     FileReader
	}

	Certs []*Cert
)

func (cc Certs) UseReadFile(readFile FileReader) Certs {
	for _, c := range cc {
		c.readFile = readFile
	}
	return cc
}

var (
	Certificates = Certs{{
		CertFilename: `s7techlab.pem`, PKeyFilename: `s7techlab.key.pem`,
	}, {
		CertFilename: `some-person.pem`, PKeyFilename: `some-person.key.pem`,
	}, {
		CertFilename: `victor-nosov.pem`, PKeyFilename: `victor-nosov.key.pem`,
	}, {
		// with attributes abac.init":"true", "admin":"true", etc.
		CertFilename: `with-attrs.pem`, PKeyFilename: ``,
	}}.
		UseReadFile(ReadLocal())
)

func ReadLocal() func(filename string) ([]byte, error) {
	_, curFile, _, ok := runtime.Caller(1)
	dir := path.Dir(curFile)
	if !ok {
		return nil
	}
	return func(filename string) ([]byte, error) {
		return ioutil.ReadFile(dir + "/" + filename)
	}
}

func (c *Cert) CertBytes() ([]byte, error) {
	return c.readFile(`./` + c.CertFilename)
}
