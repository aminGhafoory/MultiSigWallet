package main

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
)

func createKeyStoreWithExistingKey(address, privateKey, password string) ([]byte, error) {
	privateKeyBytes, err := hexutil.Decode(privateKey)
	if err != nil {
		return nil, err
	}
	PriKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		return nil, err
	}

	key := keystore.Key{
		Id:         uuid.New(),
		Address:    common.HexToAddress(address),
		PrivateKey: PriKey,
	}
	keystoreBytes, err := keystore.EncryptKey(&key, password, keystore.StandardScryptN, keystore.StandardScryptP)
	if err != nil {
		return nil, err
	}
	return keystoreBytes, nil

}

func GetEncodedPublicKeyFromPrivateKey(ECDSAPrivatekey *ecdsa.PrivateKey) string {
	PublicKeyByteSlice := crypto.FromECDSAPub(&ECDSAPrivatekey.PublicKey)
	PublicKeyEncoded := hexutil.Encode(PublicKeyByteSlice)
	return PublicKeyEncoded
}

func main() {
	keyStoreBytes, err := createKeyStoreWithExistingKey(
		"0x0B1d322814AdE8bA07f02303E0Ed16814dC3657e",
		"0x63c52a85885de291400726cd9ee6f538968963cf14c36852614ac5ebf55f2feb",
		"amin1234")
	fmt.Println(string(keyStoreBytes))
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(keyStoreBytes)

	DecKey, err := keystore.DecryptKey(keyStoreBytes, "amin1234")
	if err != nil {
		fmt.Println(err)
	}
	address := crypto.PubkeyToAddress(DecKey.PrivateKey.PublicKey)
	pk := crypto.FromECDSA(DecKey.PrivateKey)
	pkenc := hexutil.Encode(pk)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Address: ", address)
	fmt.Println("PrivateKey: ", pkenc)
	fmt.Println(hexutil.Encode(crypto.FromECDSAPub(&DecKey.PrivateKey.PublicKey)))

	walletID, _ := uuid.Parse("b7fb318e-9b9b-41ce-83c8-cc003975b4de")
	fmt.Println(walletID)
}
