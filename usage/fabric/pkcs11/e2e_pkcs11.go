/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package pkcs11

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/core"
	cryptosuite "github.com/hyperledger/fabric-sdk-go/pkg/core/cryptosuite/bccsp/pkcs11"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk/factory/defcore"
)

const (
	// ConfigTestFile contains the path and filename of the config for integration tests
	ConfigTestFilename = "config_e2e_pkcs11.yaml"
)

// CustomCryptoSuiteProviderFactory is will provide custom cryptosuite (bccsp.BCCSP)
type CustomCryptoSuiteProviderFactory struct {
	defcore.ProviderFactory
}

// CreateCryptoSuiteProvider returns a new default implementation of BCCSP
func (f *CustomCryptoSuiteProviderFactory) CreateCryptoSuiteProvider(config core.CryptoSuiteConfig) (core.CryptoSuite, error) {
	return cryptosuite.GetSuiteByConfig(config)
}
