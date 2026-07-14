package security

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"os"
)

// CertManager orchestrates the automated identity generation for microservices.
type CertManager struct {
	ServiceName string
	Namespace   string
}

// GenerateIdentity creates a new Private Key and a CSR, ready for Vault/CA signing.
func (cm *CertManager) GenerateIdentity() ([]byte, []byte, error) {
	// 1. Generate P-256 ECC Private Key (Modern standard over RSA)
	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate key: %w", err)
	}

	// 2. Define CSR Template
	template := x509.CertificateRequest{
		Subject: pkix.Name{
			CommonName:   cm.ServiceName,
			Organization: []string{"Enterprise-Internal-Trust"},
			OrganizationalUnit: []string{cm.Namespace},
		},
		SignatureAlgorithm: x509.ECDSAWithSHA256,
	}

	// 3. Create CSR
	csrBytes, err := x509.CreateCertificateRequest(rand.Reader, &template, privKey)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create CSR: %w", err)
	}

	// 4. Encode Private Key to PEM
	privBytes, _ := x509.MarshalECPrivateKey(privKey)
	privPEM := pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: privBytes})
	
	// 5. Encode CSR to PEM
	csrPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE REQUEST", Bytes: csrBytes})

	return privPEM, csrPEM, nil
}
