/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package fabric

import (
	"github.com/blockchainpro/usage/fabric/pkcs11"
	"os"
	"testing"

	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/blockchainpro/usage/fabric/integration"
	"github.com/blockchainpro/usage/fabric/configless"
)

func init() {
	os.Setenv("FABRIC_SDK_GO_PROJECT_PATH", "/root/gopath/src/github.com/blockchainpro/usage/fabric")
	os.Setenv("CRYPTOCONFIG_FIXTURES_PATH", "organizations")
	os.Setenv("username", "User1")
}

func TestE2E(t *testing.T) {
	t.Run("Base", func(t *testing.T) {
		os.Setenv("FABRIC_SDK_GO_PROJECT_PATH", "/root/gopath/src/github.com/blockchainpro/usage/fabric")
		os.Setenv("CRYPTOCONFIG_FIXTURES_PATH", "organizations")
		os.Setenv("username", "User1")
		configPath := integration.GetConfigPath("config_e2e.yaml")
		Run(t, config.FromFile(configPath))
	})

	t.Run("NoOrderer", func(t *testing.T) {
		//Using setup done set above by end to end test, run below test with new config which has no orderer config inside
		runWithNoOrdererConfig(t, config.FromFile(integration.GetConfigPath("config_e2e_no_orderer.yaml")))
	})
}

// this test mimics the original e2e test with the difference of injecting interface functions implementations
// to programmatically supply configs instead of using a yaml file. With this change, application developers can fetch
// configs from any source as long as they provide their own implementations.

func TestE2E_WithoutSetup(t *testing.T) {

	//Using same Run call as e2e package but with programmatically overriding interfaces
	// since in this configless test, we are overriding all the config's interfaces, there's no need to add a configProvider
	//
	// But if someone wants to partially override the configs interfaces (by setting only some functions of either
	// EndpointConfig, CryptoSuiteConfig and/or IdentityConfig) then they need to provide a configProvider
	// with a config file that contains at least the sections that are not overridden by the provided functions
	RunWithoutSetup(t, nil,
		fabsdk.WithEndpointConfig(configless.EndpointConfigImpls...),
		fabsdk.WithCryptoSuiteConfig(configless.CryptoConfigImpls...),
		fabsdk.WithIdentityConfig(configless.IdentityConfigImpls...),
		fabsdk.WithMetricsConfig(configless.OperationsConfigImpls...),
	)
}

func TestE2E_pkcs11(t *testing.T) {
	// Create SDK setup for the integration tests
	Run(t,
		config.FromFile(integration.GetConfigPath(pkcs11.ConfigTestFilename)),
		fabsdk.WithCorePkg(&pkcs11.CustomCryptoSuiteProviderFactory{}))
}

