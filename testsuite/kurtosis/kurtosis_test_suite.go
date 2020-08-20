package kurtosis

import (
	"github.com/ava-labs/avalanche-e2e-tests/testsuite/tests/conflictvtx"
	"github.com/ava-labs/avalanche-e2e-tests/testsuite/tests/connected"
	"github.com/ava-labs/avalanche-e2e-tests/testsuite/tests/duplicate"
	"github.com/ava-labs/avalanche-e2e-tests/testsuite/tests/spamchits"
	"github.com/ava-labs/avalanche-e2e-tests/testsuite/tests/workflow"
	"github.com/ava-labs/avalanche-e2e-tests/testsuite/verifier"
	"github.com/kurtosis-tech/kurtosis/commons/testsuite"
)

// AvalancheTestSuite implements the Kurtosis TestSuite interface
type AvalancheTestSuite struct {
	ByzantineImageName string
	NormalImageName    string
}

// GetTests implements the Kurtosis TestSuite interface
func (a AvalancheTestSuite) GetTests() map[string]testsuite.Test {
	result := make(map[string]testsuite.Test)

	if a.ByzantineImageName != "" {
		result["stakingNetworkChitSpammerTest"] = spamchits.StakingNetworkUnrequestedChitSpammerTest{
			ByzantineImageName: a.ByzantineImageName,
			NormalImageName:    a.NormalImageName,
		}
		result["conflictingTxsVertexTest"] = conflictvtx.StakingNetworkConflictingTxsVertexTest{
			ByzantineImageName: a.ByzantineImageName,
			NormalImageName:    a.NormalImageName,
		}
	}
	result["stakingNetworkFullyConnectedTest"] = connected.StakingNetworkFullyConnectedTest{
		ImageName: a.NormalImageName,
		Verifier:  verifier.NetworkStateVerifier{},
	}
	result["stakingNetworkDuplicateNodeIDTest"] = duplicate.DuplicateNodeIDTest{
		ImageName: a.NormalImageName,
		Verifier:  verifier.NetworkStateVerifier{},
	}
	result["StakingNetworkRPCWorkflowTest"] = workflow.StakingNetworkRPCWorkflowTest{
		ImageName: a.NormalImageName,
	}

	return result
}
