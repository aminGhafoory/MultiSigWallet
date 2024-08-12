package models

import (
	"context"
	"time"

	"github.com/aminghafoory/multisig2/internal/database"
	"github.com/aminghafoory/multisig2/utils"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/google/uuid"
)

type WalletService struct {
	DB        *database.Queries
	ETHClient *ethclient.Client
}

func (ws WalletService) CreateNewWallet(userID uuid.UUID, password, walletName string) error {
	keyStoreBytes, pubkey, err := utils.CreateKeyStoreBytes(password)
	if err != nil {
		return err
	}
	err = ws.DB.CreateaWallet(context.Background(), database.CreateaWalletParams{
		UserID:         userID,
		CreatedAt:      time.Now(),
		WalletID:       uuid.New(),
		WalletName:     walletName,
		WalletAddress:  crypto.PubkeyToAddress(pubkey).Hex(),
		WalletKeystore: string(keyStoreBytes),
	})
	if err != nil {
		return err
	}

	return nil
}

func (ws WalletService) ImportWallet(userID uuid.UUID, privateKey, password, walletName string) error {
	keyStoreBytes, pubkey, err := utils.CreateKeyStoreWithExistingKey(privateKey, password)
	if err != nil {
		return err
	}
	err = ws.DB.CreateaWallet(context.Background(), database.CreateaWalletParams{
		UserID:         userID,
		CreatedAt:      time.Now(),
		WalletID:       uuid.New(),
		WalletName:     walletName,
		WalletAddress:  crypto.PubkeyToAddress(pubkey).Hex(),
		WalletKeystore: string(keyStoreBytes),
	})
	if err != nil {
		return err
	}

	return nil
}

func (ws WalletService) UserWallets(userID uuid.UUID) ([]database.GetUserWalletsRow, error) {
	Wallets, err := ws.DB.GetUserWallets(context.Background(), userID)
	if err != nil {
		return nil, err
	}
	return Wallets, nil
}

func (ws WalletService) DeleteWallet(walletID, userID uuid.UUID) error {
	err := ws.DB.DeleteWallet(context.Background(), database.DeleteWalletParams{
		UserID:   userID,
		WalletID: walletID,
	})
	return err
}

func (ws WalletService) GetUserWalletKeystoreByID(walletID, userID uuid.UUID) (string, error) {
	keystoreBytes, err := ws.DB.GetUserWalletByID(context.Background(), database.GetUserWalletByIDParams{
		UserID:   userID,
		WalletID: walletID,
	})
	if err != nil {
		return "", err
	}

	return keystoreBytes, nil
}
