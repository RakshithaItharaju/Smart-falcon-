
package main

import (
    "encoding/json"
    "fmt"
    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
    contractapi.Contract
}

type Asset struct {
    DealerID    string  `json:"dealerId"`
    MSISDN      string  `json:"msisdn"`
    MPIN        string  `json:"mpin"`
    Balance     float64 `json:"balance"`
    Status      string  `json:"status"`
    TransAmount float64 `json:"transAmount"`
    TransType   string  `json:"transType"`
    Remarks     string  `json:"remarks"`
}

func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
    return nil
}

func (s *SmartContract) CreateAsset(ctx contractapi.TransactionContextInterface, id string, dealerId string, msisdn string, mpin string, balance float64, status string, transAmount float64, transType string, remarks string) error {
    asset := Asset{
        DealerID:    dealerId,
        MSISDN:      msisdn,
        MPIN:        mpin,
        Balance:     balance,
        Status:      status,
        TransAmount: transAmount,
        TransType:   transType,
        Remarks:     remarks,
    }

    assetJSON, err := json.Marshal(asset)
    if err != nil {
        return err
    }

    return ctx.GetStub().PutState(id, assetJSON)
}

func (s *SmartContract) ReadAsset(ctx contractapi.TransactionContextInterface, id string) (*Asset, error) {
    assetJSON, err := ctx.GetStub().GetState(id)
    if err != nil {
        return nil, fmt.Errorf("failed to read from world state: %v", err)
    }
    if assetJSON == nil {
        return nil, fmt.Errorf("the asset %s does not exist", id)
    }

    var asset Asset
    err = json.Unmarshal(assetJSON, &asset)
    if err != nil {
        return nil, err
    }

    return &asset, nil
}

func main() {
    chaincode, err := contractapi.NewChaincode(new(SmartContract))
    if err != nil {
        panic("Error creating chaincode")
    }
    if err := chaincode.Start(); err != nil {
        panic(err.Error())
    }
}
