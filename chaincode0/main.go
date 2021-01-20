package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type contract struct {
	contractapi.Contract
}

type Item struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func (c *contract) InitLedger(_ contractapi.TransactionContextInterface) error {
	return nil
}

func (c *contract) CreateItem(ctx contractapi.TransactionContextInterface, id, title string) error {
	state, err := ctx.GetStub().GetState(id)
	if err != nil {
		return fmt.Errorf("error on get state: %w", err)
	}

	if state != nil {
		return fmt.Errorf("item with id '%s' already exists", id)
	}

	item := Item{
		ID:    id,
		Title: title,
	}

	b, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("error on marshal json: %w", err)
	}

	err = ctx.GetStub().PutState(id, b)
	if err != nil {
		return fmt.Errorf("error on put state: %w", err)
	}

	return nil
}

func (c *contract) GetItem(ctx contractapi.TransactionContextInterface, id string) (*Item, error) {
	state, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("error on get state: %w", err)
	}

	if state == nil {
		return nil, fmt.Errorf("item with id '%s' not found", id)
	}

	var item Item

	err = json.Unmarshal(state, &item)
	if err != nil {
		return nil, fmt.Errorf("error on unmarshal state: %w", err)
	}

	return &item, nil
}

func main() {
	ci := new(contract)
	cc, err := contractapi.NewChaincode(ci)
	if err != nil {
		panic(fmt.Errorf("error on new chaincode: %w", err))
	}

	err = cc.Start()
	if err != nil {
		panic(fmt.Errorf("error on start chaincode: %w", err))
	}
}
