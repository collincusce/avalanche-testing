package gecko_client

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateSubnet(t *testing.T) {
	resultStr := `{
		"jsonrpc": "2.0",
		"result": {
			"unsignedTx": "1112LA7e8GvkGHDkxZa9Q7kszqvWHooumX5PhqA9NJG7erwXYcwQUPRQyukYX1ncu1DmWvvPNMuivUqvGp1t9M3wys5joqXrXtV2jescQ5AWaUKHiSBUWBRHseMLhGxWNT4Bv6LNVvaaA1ZW33avQBAzz7V84KpKGW7fD3Fz1okxknLgoG"
		},
		"id": 1
	}`
	client := clientFromRequester(mockedJsonRpcRequester{resultStr: resultStr})
	unsignedTxn, err := client.PChainApi().CreateSubnet([]string{"key1", "key2"}, 1, 1)
	assert.Nil(t, err, "Error message should be nil")

	assert.Equal(
		t,
		"1112LA7e8GvkGHDkxZa9Q7kszqvWHooumX5PhqA9NJG7erwXYcwQUPRQyukYX1ncu1DmWvvPNMuivUqvGp1t9M3wys5joqXrXtV2jescQ5AWaUKHiSBUWBRHseMLhGxWNT4Bv6LNVvaaA1ZW33avQBAzz7V84KpKGW7fD3Fz1okxknLgoG",
		unsignedTxn)
}

func TestGetSubnets(t *testing.T) {
	resultStr := `{
		"jsonrpc": "2.0",
		"result": {
			"subnets": [
				{
					"id": "hW8Ma7dLMA7o4xmJf3AXBbo17bXzE7xnThUd3ypM4VAWo1sNJ",
					"controlKeys": [
						"KNjXsaA1sZsaKCD1cd85YXauDuxshTes2",
						"Aiz4eEt5xv9t4NCnAWaQJFNz5ABqLtJkR"
					],
					"threshold": "2"
				}
			]
		},
		"id": 6
	}`
	client := clientFromRequester(mockedJsonRpcRequester{resultStr: resultStr})
	subnetList, err := client.PChainApi().GetSubnets()
	assert.Nil(t, err, "Error message should be nil")

	assert.Equal(t, 1, len(subnetList))
	assert.Equal(t, 2, len(subnetList[0].ControlKeys))
	assert.Equal(t, "2", subnetList[0].Threshold)
	assert.Equal(
		t,
		"hW8Ma7dLMA7o4xmJf3AXBbo17bXzE7xnThUd3ypM4VAWo1sNJ",
		subnetList[0].Id)
}

func TestValidatedBy(t *testing.T) {
	resultStr := `{
		"jsonrpc": "2.0",
		"result": {
			"subnetID": "2bRCr6B4MiEfSjidDwxDpdCyviwnfUVqB2HGwhm947w9YYqb7r"
		},
		"id": 1
	}`
	client := clientFromRequester(mockedJsonRpcRequester{resultStr: resultStr})
	subnetId, err := client.PChainApi().ValidatedBy("KDYHHKjM4yTJTT8H8qPs5KXzE6gQH5TZrmP1qVr1P6qECj3XN")
	assert.Nil(t, err, "Error message should be nil")

	assert.Equal(t,"2bRCr6B4MiEfSjidDwxDpdCyviwnfUVqB2HGwhm947w9YYqb7r", subnetId)
}

func TestValidates(t *testing.T) {
	resultStr := `{
		"jsonrpc": "2.0",
		"result": {
			"blockchainIDs": [
				"KDYHHKjM4yTJTT8H8qPs5KXzE6gQH5TZrmP1qVr1P6qECj3XN",
				"2TtHFqEAAJ6b33dromYMqfgavGPF3iCpdG3hwNMiart2aB5QHi"
			]
		},
		"id": 1
	}`
	client := clientFromRequester(mockedJsonRpcRequester{resultStr: resultStr})
	blockchainIds, err := client.PChainApi().Validates("2bRCr6B4MiEfSjidDwxDpdCyviwnfUVqB2HGwhm947w9YYqb7r")
	assert.Nil(t, err, "Error message should be nil")

	assert.Equal(t, 2, len(blockchainIds))
	assert.Equal(t, "KDYHHKjM4yTJTT8H8qPs5KXzE6gQH5TZrmP1qVr1P6qECj3XN", blockchainIds[0])
	assert.Equal(t, "2TtHFqEAAJ6b33dromYMqfgavGPF3iCpdG3hwNMiart2aB5QHi", blockchainIds[1])
}

func TestGetBlockchains(t *testing.T) {
	resultStr := `{
		"jsonrpc": "2.0",
		"result": {
			"blockchains": [
				{
					"id": "mnihvmrJ4MiojP7qhnF3sKR43RJvJbHrbkM8yFoLdwc4nwEqV",
					"name": "AVM",
					"subnetID": "11111111111111111111111111111111LpoYY",
					"vmID": "jvYyfQTxGMJLuGWa55kdP2p2zSUYsQ5Raupu4TW34ZAUBAbtq"
				},
				{
					"id": "2rWWhAMu2NyNPEHrDNnTfhtdV9MZWKkp1L5D6ANWnhAkJAkosN",
					"name": "Athereum",
					"subnetID": "11111111111111111111111111111111LpoYY",
					"vmID": "mgj786NP7uDwBCcq6YwThhaN8FLyybkCa4zBWTQbNgmK6k9A6"
				},
				{
					"id": "CqhF97NNugqYLiGaQJ2xckfmkEr8uNeGG5TQbyGcgnZ5ahQwa",
					"name": "Simple DAG Payments",
					"subnetID": "11111111111111111111111111111111LpoYY",
					"vmID": "sqjdyTKUSrQs1YmKDTUbdUhdstSdtRTGRbUn8sqK8B6pkZkz1"
				},
				{
					"id": "VcqKNBJsYanhVFxGyQE5CyNVYxL3ZFD7cnKptKWeVikJKQkjv",
					"name": "Simple Chain Payments",
					"subnetID": "11111111111111111111111111111111LpoYY",
					"vmID": "sqjchUjzDqDfBPGjfQq2tXW1UCwZTyvzAWHsNzF2cb1eVHt6w"
				},
				{
					"id": "2SMYrx4Dj6QqCEA3WjnUTYEFSnpqVTwyV3GPNgQqQZbBbFgoJX",
					"name": "Simple Timestamp Server",
					"subnetID": "11111111111111111111111111111111LpoYY",
					"vmID": "tGas3T58KzdjLHhBDMnH2TvrddhqTji5iZAMZ3RXs2NLpSnhH"
				},
				{
					"id": "KDYHHKjM4yTJTT8H8qPs5KXzE6gQH5TZrmP1qVr1P6qECj3XN",
					"name": "My new timestamp",
					"subnetID": "2bRCr6B4MiEfSjidDwxDpdCyviwnfUVqB2HGwhm947w9YYqb7r",
					"vmID": "tGas3T58KzdjLHhBDMnH2TvrddhqTji5iZAMZ3RXs2NLpSnhH"
				},
				{
					"id": "2TtHFqEAAJ6b33dromYMqfgavGPF3iCpdG3hwNMiart2aB5QHi",
					"name": "My new AVM",
					"subnetID": "2bRCr6B4MiEfSjidDwxDpdCyviwnfUVqB2HGwhm947w9YYqb7r",
					"vmID": "jvYyfQTxGMJLuGWa55kdP2p2zSUYsQ5Raupu4TW34ZAUBAbtq"
				}
			]
		},
		"id": 85
	}`
	client := clientFromRequester(mockedJsonRpcRequester{resultStr: resultStr})
	blockchains, err := client.PChainApi().GetBlockchains()
	assert.Nil(t, err, "Error message should be nil")

	assert.Equal(t, 7, len(blockchains))
	assert.Equal(t, "My new AVM", blockchains[6].Name)
	assert.Equal(t, "KDYHHKjM4yTJTT8H8qPs5KXzE6gQH5TZrmP1qVr1P6qECj3XN", blockchains[5].Id)
	assert.Equal(t, "11111111111111111111111111111111LpoYY", blockchains[4].SubnetID)
	assert.Equal(t, "jvYyfQTxGMJLuGWa55kdP2p2zSUYsQ5Raupu4TW34ZAUBAbtq", blockchains[0].VmID)
}


func TestExportAVA(t *testing.T) {
	resultStr := `{
		"jsonrpc": "2.0",
		"result": {
			"unsignedTx": "1112Y8Y5ibRqMDtby9NSdpK9u3n1yGywybAAVYnhCkFYcRzEYbR7J5Ci6SX98PmgS2LpRf5pcu6YAgLYGiTuQpiSucRcX4dv7HbVnEsrQnjcieGbgkf9PFS126hC8xce4pEZUzr9jReVdfXe3g9BSUsXLj2XcWrnD6iTgHpiC18jjyjg1wjm1Vs4TcXhG472MRvGspucJ8LuUE91WV7353Kxdc2e7Trw2Sd6iV"
		},
		"id": 1
	}`
	client := clientFromRequester(mockedJsonRpcRequester{resultStr: resultStr})
	unsignedTx, err := client.PChainApi().ExportAVA(1,"G5ZGXEfoWYNFZH5JF9C4QPKAbPTKwRbyB", 2)
	assert.Nil(t, err, "Error message should be nil")

	assert.Equal(
		t,
		"1112Y8Y5ibRqMDtby9NSdpK9u3n1yGywybAAVYnhCkFYcRzEYbR7J5Ci6SX98PmgS2LpRf5pcu6YAgLYGiTuQpiSucRcX4dv7HbVnEsrQnjcieGbgkf9PFS126hC8xce4pEZUzr9jReVdfXe3g9BSUsXLj2XcWrnD6iTgHpiC18jjyjg1wjm1Vs4TcXhG472MRvGspucJ8LuUE91WV7353Kxdc2e7Trw2Sd6iV",
		unsignedTx)
}

func TestImportAVA(t *testing.T) {
	resultStr := `{
		"jsonrpc": "2.0",
		"result": {
			"tx": "1117xBwcr5fo1Ch4umyzjYgnuoFhSwBHdMCam2wRe8SxcJJvQRKSmufXM8aSqKaDmX4TjvzPaUbSn33TAQsbZDhzcHEGviuthncY5VQfUJogyMoFGXUtu3M8NbwNhrYtmSRkFdmN4w933janKvJYKNnsDMvMkmasxrFj8fQxE6Ej8eyU2Jqj2gnTxU2WD3NusFNKmPfgJs8DRCWgYyJVodnGvT43hovggVaWHHD8yYi9WJ64pLCvtCcEYkQeEeA5NE8eTxPtWJrwSMTciHHVdHMpxdVAY6Ptr2rMcYSacr8TZzw59XJfbQT4R6DCsHYQAPJAUfDNeX2JuiBk9xonfKmGcJcGXwdJZ3QrvHHHfHCeuxqS13AfU"
		},
		"id": 1
	}`
	client := clientFromRequester(mockedJsonRpcRequester{resultStr: resultStr})
	tx, err := client.PChainApi().ImportAVA(
		"bob",
		"loblaw",
		"Bg6e45gxCUTLXcfUuoy3go2U6V3bRZ5jH",
		1)
	assert.Nil(t, err, "Error message should be nil")

	assert.Equal(t,
		"1117xBwcr5fo1Ch4umyzjYgnuoFhSwBHdMCam2wRe8SxcJJvQRKSmufXM8aSqKaDmX4TjvzPaUbSn33TAQsbZDhzcHEGviuthncY5VQfUJogyMoFGXUtu3M8NbwNhrYtmSRkFdmN4w933janKvJYKNnsDMvMkmasxrFj8fQxE6Ej8eyU2Jqj2gnTxU2WD3NusFNKmPfgJs8DRCWgYyJVodnGvT43hovggVaWHHD8yYi9WJ64pLCvtCcEYkQeEeA5NE8eTxPtWJrwSMTciHHVdHMpxdVAY6Ptr2rMcYSacr8TZzw59XJfbQT4R6DCsHYQAPJAUfDNeX2JuiBk9xonfKmGcJcGXwdJZ3QrvHHHfHCeuxqS13AfU",
		tx)
}


func TestSign(t *testing.T) {
	resultStr := `{
		"jsonrpc": "2.0",
		"result": {
			"Tx": "111Bit5JNASbJyTLrd2kWkYRoc96swEWoWdmEhuGAFK3rCAyTnTzomuFwgx1SCUdUE71KbtXPnqj93KGr3CeftpPN37kVyqBaAQ5xaDjr7wVBTUYi9iV7kYJnHF61yovViJF74mJJy7WWQKeRMDRTiPuii5gsd11gtNahCCsKbm9seJtk2h1wAPZn9M1eL84CGVPnLUiLP"
		},
		"id": 1
	}`
	client := clientFromRequester(mockedJsonRpcRequester{resultStr: resultStr})
	signTx, err := client.PChainApi().Sign(
		"111Bit5JNASbJyTLrd2kWkYRoc96swEWoWdmEhuGAFK3rCAyTnTzomuFwgx1SCUdUE71KbtXPnqj93KGr3CeftpPN37kVyqBaAQ5xaDjr7wU8riGS89NDJ8AwVgZgnFkgF3uMfwCiCuPvvubGyQxNHE4TM9iDgj6h3URdGQ4JntP44wokCEP3ADn7sMM8kUTbmcNo84U87",
		"6Y3kysjF9jnHnYkdS9yGAuoHyae2eNmeV",
		"bob",
		"loblaw")
	assert.Nil(t, err, "Error message should be nil")

	assert.Equal(
		t,
		"111Bit5JNASbJyTLrd2kWkYRoc96swEWoWdmEhuGAFK3rCAyTnTzomuFwgx1SCUdUE71KbtXPnqj93KGr3CeftpPN37kVyqBaAQ5xaDjr7wVBTUYi9iV7kYJnHF61yovViJF74mJJy7WWQKeRMDRTiPuii5gsd11gtNahCCsKbm9seJtk2h1wAPZn9M1eL84CGVPnLUiLP",
		signTx)
}

func TestIssueTx(t *testing.T) {
	resultStr := `{
		"jsonrpc": "2.0",
		"result": {
			"txID": "G3BuH6ytQ2averrLxJJugjWZHTRubzCrUZEXoheG5JMqL5ccY"
		},
		"id": 1
	}`
	client := clientFromRequester(mockedJsonRpcRequester{resultStr: resultStr})
	txnId, err := client.PChainApi().IssueTx("111Bit5JNASbJyTLrd2kWkYRoc96swEWoWdmEhuGAFK3rCAyTnTzomuFwgx1SCUdUE71KbtXPnqj93KGr3CeftpPN37kVyqBaAQ5xaDjr7wVBTUYi9iV7kYJnHF61yovViJF74mJJy7WWQKeRMDRTiPuii5gsd11gtNahCCsKbm9seJtk2h1wAPZn9M1eL84CGVPnLUiLP")
	assert.Nil(t, err, "Error message should be nil")

	assert.Equal(t, txnId, "G3BuH6ytQ2averrLxJJugjWZHTRubzCrUZEXoheG5JMqL5ccY")
}
