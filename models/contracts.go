package models

import (
	"context"
	"math/big"
	"time"

	"github.com/aminghafoory/multisig2/internal/database"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/google/uuid"
)

type ContractsService struct {
	DB        *database.Queries
	ETHClient *ethclient.Client
}

func (cs *ContractsService) CreateContract(
	password, userKeyStorePath, contractName string,
	owners []common.Address, numOfConfirmationRequired *big.Int,
	contractAddress common.Address) (database.Contract, error) {

	contract, err := cs.DB.CreateaContract(context.Background(), database.CreateaContractParams{
		UserID:       [16]byte{},
		CreatedAt:    time.Time{},
		ContractID:   contractAddress.Hex(),
		ContractName: contractName,
	})
	if err != nil {
		return database.Contract{}, err
	}

	return contract, nil

}

func (cs *ContractsService) UserContracts(userID uuid.UUID) ([]database.GetUserContractsRow, error) {
	contratcs, err := cs.DB.GetUserContracts(context.Background(), userID)
	if err != nil {
		return nil, err
	}
	return contratcs, nil
}

func (cs *ContractsService) DeleteContract(contractID string, userID uuid.UUID) error {
	err := cs.DB.DeleteContract(context.Background(), database.DeleteContractParams{
		UserID:     userID,
		ContractID: contractID,
	})
	return err
}
