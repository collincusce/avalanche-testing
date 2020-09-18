// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package platform

import (
	"time"

	"github.com/ava-labs/avalanchego/api"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/formatting"
	cjson "github.com/ava-labs/avalanchego/utils/json"
	"github.com/ava-labs/avalanchego/vms/platformvm"

	"github.com/ava-labs/avalanche-testing/avalanche_client/utils"
)

// Client ...
type Client struct {
	requester utils.EndpointRequester
}

// NewClient returns a Client for interacting with the P Chain endpoint
func NewClient(uri string, requestTimeout time.Duration) *Client {
	return &Client{
		requester: utils.NewEndpointRequester(uri, "/ext/P", "platform", requestTimeout),
	}
}

// GetHeight returns the current block height of the P Chain
func (c *Client) GetHeight() (uint64, error) {
	res := &platformvm.GetHeightResponse{}
	err := c.requester.SendRequest("getHeight", struct{}{}, res)
	return uint64(res.Height), err
}

// ExportKey returns the private key corresponding to [address] from [user]'s account
func (c *Client) ExportKey(user api.UserPass, address string) (string, error) {
	res := &platformvm.ExportKeyReply{}
	err := c.requester.SendRequest("exportKey", &platformvm.ExportKeyArgs{
		UserPass: user,
		Address:  address,
	}, res)
	return res.PrivateKey, err
}

// ImportKey imports the specified [privateKey] to [user]'s keystore
func (c *Client) ImportKey(user api.UserPass, privateKey string) (string, error) {
	res := &api.JsonAddress{}
	err := c.requester.SendRequest("importKey", &platformvm.ImportKeyArgs{
		UserPass:   user,
		PrivateKey: privateKey,
	}, res)
	return res.Address, err
}

// GetBalance returns the balance of [address] on the P Chain
func (c *Client) GetBalance(address string) (*platformvm.GetBalanceResponse, error) {
	res := &platformvm.GetBalanceResponse{}
	err := c.requester.SendRequest("getBalance", &platformvm.GetBalanceArgs{
		Address: address,
	}, res)
	return res, err
}

// CreateAddress creates a new address for [user]
func (c *Client) CreateAddress(user api.UserPass) (string, error) {
	res := &api.JsonAddress{}
	err := c.requester.SendRequest("createAddress", &user, res)
	return res.Address, err
}

// ListAddresses returns an array of platform addresses controlled by [user]
func (c *Client) ListAddresses(user api.UserPass) ([]string, error) {
	res := &api.JsonAddresses{}
	err := c.requester.SendRequest("listAddresses", &user, res)
	return res.Addresses, err
}

// GetUTXOs returns the byte representation of the UTXOs controlled by [addresses]
func (c *Client) GetUTXOs(addresses []string) ([][]byte, error) {
	res := &platformvm.GetUTXOsResponse{}
	err := c.requester.SendRequest("getUTXOs", &platformvm.GetUTXOsArgs{
		Addresses: addresses,
	}, res)
	if err != nil {
		return nil, err
	}
	utxos := make([][]byte, len(res.UTXOs))
	for i, utxo := range res.UTXOs {
		utxos[i] = utxo.Bytes
	}
	return utxos, nil
}

// GetSubnets returns information about the specified subnets
func (c *Client) GetSubnets(ids []ids.ID) ([]platformvm.APISubnet, error) {
	res := &platformvm.GetSubnetsResponse{}
	err := c.requester.SendRequest("getSubnets", &platformvm.GetSubnetsArgs{
		IDs: ids,
	}, res)
	return res.Subnets, err
}

// GetCurrentValidators returns the list of current validators and the list of delegators for subnet with ID [subnetID]
func (c *Client) GetCurrentValidators(subnetID ids.ID) ([]interface{}, []interface{}, error) {
	res := &platformvm.GetCurrentValidatorsReply{}
	err := c.requester.SendRequest("getCurrentValidators", &platformvm.GetCurrentValidatorsArgs{
		SubnetID: subnetID,
	}, res)
	return res.Validators, res.Delegators, err
}

// GetPendingValidators returns the list of pending validators for subnet with ID [subnetID]
func (c *Client) GetPendingValidators(subnetID ids.ID) ([]interface{}, []interface{}, error) {
	res := &platformvm.GetPendingValidatorsReply{}
	err := c.requester.SendRequest("getPendingValidators", &platformvm.GetPendingValidatorsArgs{
		SubnetID: subnetID,
	}, res)
	return res.Validators, res.Delegators, err
}

// SampleValidators returns the nodeIDs of a sample of [sampleSize] validators from the current validator set for subnet with ID [subnetID]
func (c *Client) SampleValidators(subnetID ids.ID, sampleSize uint16) (*platformvm.SampleValidatorsReply, error) {
	res := &platformvm.SampleValidatorsReply{}
	err := c.requester.SendRequest("sampleValidators", &platformvm.SampleValidatorsArgs{
		SubnetID: subnetID,
		Size:     cjson.Uint16(sampleSize),
	}, res)
	return res, err
}

// AddValidator issues a transaction to add a validator to the primary network and returns the txID
func (c *Client) AddValidator(user api.UserPass, rewardAddress, nodeID string, stakeAmount, startTime, endTime uint64, delegationFeeRate float32) (ids.ID, error) {
	res := &api.JsonTxID{}
	jsonStakeAmount := cjson.Uint64(stakeAmount)
	err := c.requester.SendRequest("addValidator", &platformvm.AddValidatorArgs{
		JsonSpendHeader: api.JsonSpendHeader{
			UserPass: user,
		},
		APIStaker: platformvm.APIStaker{
			NodeID:      nodeID,
			StakeAmount: &jsonStakeAmount,
			StartTime:   cjson.Uint64(startTime),
			EndTime:     cjson.Uint64(endTime),
		},
		RewardAddress:     rewardAddress,
		DelegationFeeRate: cjson.Float32(delegationFeeRate),
	}, res)
	return res.TxID, err
}

// AddDelegator issues a transaction to add a delegator to the primary network and returns the txID
func (c *Client) AddDelegator(
	user api.UserPass,
	rewardAddress,
	nodeID string,
	stakeAmount,
	startTime,
	endTime uint64,
	from []string,
	changeAddr string,
) (ids.ID, error) {
	res := &api.JsonTxID{}
	jsonStakeAmount := cjson.Uint64(stakeAmount)
	err := c.requester.SendRequest("addDelegator", &platformvm.AddDelegatorArgs{
		JsonSpendHeader: api.JsonSpendHeader{
			UserPass:       user,
			JsonFromAddrs:  api.JsonFromAddrs{From: from},
			JsonChangeAddr: api.JsonChangeAddr{ChangeAddr: changeAddr},
		}, APIStaker: platformvm.APIStaker{
			NodeID:      nodeID,
			StakeAmount: &jsonStakeAmount,
			StartTime:   cjson.Uint64(startTime),
			EndTime:     cjson.Uint64(endTime),
		},
		RewardAddress: rewardAddress,
	}, res)
	return res.TxID, err
}

// AddSubnetValidator issues a transaction to add validator [nodeID] to subnet with ID [subnetID] and returns the txID
func (c *Client) AddSubnetValidator(
	user api.UserPass,
	destination,
	nodeID string,
	stakeAmount,
	startTime,
	endTime uint64,
	subnetID string,
	from []string,
	changeAddr string,
) (ids.ID, error) {
	res := &api.JsonTxID{}
	jsonStakeAmount := cjson.Uint64(stakeAmount)
	err := c.requester.SendRequest("addSubnetValidator", &platformvm.AddSubnetValidatorArgs{
		JsonSpendHeader: api.JsonSpendHeader{
			UserPass:       user,
			JsonFromAddrs:  api.JsonFromAddrs{From: from},
			JsonChangeAddr: api.JsonChangeAddr{ChangeAddr: changeAddr},
		},
		APIStaker: platformvm.APIStaker{
			NodeID:      nodeID,
			StakeAmount: &jsonStakeAmount,
			StartTime:   cjson.Uint64(startTime),
			EndTime:     cjson.Uint64(endTime),
		},
		SubnetID: subnetID,
	}, res)
	return res.TxID, err
}

// CreateSubnet issues a transaction to create [subnet] and returns the txID
func (c *Client) CreateSubnet(
	user api.UserPass,
	subnet platformvm.APISubnet,
	from []string,
	changeAddr string,
) (ids.ID, error) {
	res := &api.JsonTxID{}
	err := c.requester.SendRequest("createSubnet", &platformvm.CreateSubnetArgs{
		JsonSpendHeader: api.JsonSpendHeader{
			UserPass:       user,
			JsonFromAddrs:  api.JsonFromAddrs{From: from},
			JsonChangeAddr: api.JsonChangeAddr{ChangeAddr: changeAddr},
		},
		APISubnet: subnet,
	}, res)
	return res.TxID, err
}

// ExportAVAX issues an ExportAVAX transaction and returns the txID
func (c *Client) ExportAVAX(
	user api.UserPass,
	to string,
	amount uint64,
	from []string,
	changeAddr string,
) (ids.ID, error) {
	res := &api.JsonTxID{}
	err := c.requester.SendRequest("exportAVAX", &platformvm.ExportAVAXArgs{
		JsonSpendHeader: api.JsonSpendHeader{
			UserPass:       user,
			JsonFromAddrs:  api.JsonFromAddrs{From: from},
			JsonChangeAddr: api.JsonChangeAddr{ChangeAddr: changeAddr},
		},
		To:     to,
		Amount: cjson.Uint64(amount),
	}, res)
	return res.TxID, err
}

// ImportAVAX issues an ImportAVAX transaction and returns the txID
func (c *Client) ImportAVAX(
	user api.UserPass,
	to,
	sourceChain string,
	from []string,
	changeAddr string,
) (ids.ID, error) {
	res := &api.JsonTxID{}
	err := c.requester.SendRequest("importAVAX", &platformvm.ImportAVAXArgs{
		JsonSpendHeader: api.JsonSpendHeader{
			UserPass:       user,
			JsonFromAddrs:  api.JsonFromAddrs{From: from},
			JsonChangeAddr: api.JsonChangeAddr{ChangeAddr: changeAddr},
		},
		To:          to,
		SourceChain: sourceChain,
	}, res)
	return res.TxID, err
}

// CreateBlockchain issues a CreateBlockchain transaction and returns the txID
func (c *Client) CreateBlockchain(
	user api.UserPass,
	subnetID ids.ID,
	vmID string,
	fxIDs []string,
	name string,
	genesisData []byte,
	from []string,
	changeAddr string,
) (ids.ID, error) {
	res := &api.JsonTxID{}
	err := c.requester.SendRequest("createBlockchain", &platformvm.CreateBlockchainArgs{
		JsonSpendHeader: api.JsonSpendHeader{
			UserPass:       user,
			JsonFromAddrs:  api.JsonFromAddrs{From: from},
			JsonChangeAddr: api.JsonChangeAddr{ChangeAddr: changeAddr},
		},
		SubnetID:    subnetID,
		VMID:        vmID,
		FxIDs:       fxIDs,
		Name:        name,
		GenesisData: formatting.CB58{Bytes: genesisData},
	}, res)
	return res.TxID, err
}

// GetBlockchainStatus returns the current status of blockchain with ID: [blockchainID]
func (c *Client) GetBlockchainStatus(blockchainID string) (platformvm.Status, error) {
	res := &platformvm.GetBlockchainStatusReply{}
	err := c.requester.SendRequest("getBlockchainStatus", &platformvm.GetBlockchainStatusArgs{
		BlockchainID: blockchainID,
	}, res)
	return res.Status, err
}

// ValidatedBy returns the ID of the Subnet that validates [blockchainID]
func (c *Client) ValidatedBy(blockchainID ids.ID) (ids.ID, error) {
	res := &platformvm.ValidatedByResponse{}
	err := c.requester.SendRequest("validatedBy", &platformvm.ValidatedByArgs{
		BlockchainID: blockchainID,
	}, res)
	return res.SubnetID, err
}

// Validates returns the list of blockchains that are validated by the subnet with ID [subnetID]
func (c *Client) Validates(subnetID ids.ID) ([]ids.ID, error) {
	res := &platformvm.ValidatesResponse{}
	err := c.requester.SendRequest("validates", &platformvm.ValidatesArgs{
		SubnetID: subnetID,
	}, res)
	return res.BlockchainIDs, err
}

// GetBlockchains returns the list of blockchains on the platform
func (c *Client) GetBlockchains() ([]platformvm.APIBlockchain, error) {
	res := &platformvm.GetBlockchainsResponse{}
	err := c.requester.SendRequest("getBlockchains", struct{}{}, res)
	return res.Blockchains, err
}

// GetTx returns the byte representation of the transaction corresponding to [txID]
func (c *Client) GetTx(txID ids.ID) ([]byte, error) {
	res := &platformvm.GetTxResponse{}
	err := c.requester.SendRequest("getTx", &platformvm.GetTxArgs{
		TxID: txID,
	}, res)
	return res.Tx.Bytes, err
}

// GetTxStatus returns the status of the transaction corresponding to [txID]
func (c *Client) GetTxStatus(txID ids.ID) (platformvm.Status, error) {
	res := new(platformvm.Status)
	err := c.requester.SendRequest("getTxStatus", &platformvm.GetTxStatusArgs{
		TxID: txID,
	}, res)
	return *res, err
}
