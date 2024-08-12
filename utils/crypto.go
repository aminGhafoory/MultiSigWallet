package utils

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"regexp"
	"strings"

	MultiSig "github.com/aminghafoory/multisig2/contracts/abigenBuild"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/google/uuid"
)

func ValidateUserAddress(Client *ethclient.Client, address string) error {
	RegexValidator := regexp.MustCompile("^0x[0-9a-zA-Z]{40}$")
	isRegexValid := RegexValidator.MatchString(address)
	if !isRegexValid {
		return fmt.Errorf("address is not valid")
	}
	code, err := Client.CodeAt(context.Background(), common.HexToAddress(address), nil)
	if err != nil {
		log.Fatalf("error: can not get codeAt %+v", err)
		return fmt.Errorf("can not verify address")
	}

	isContract := len(code) > 0

	if isContract {
		return fmt.Errorf("address is a Contract address")
	}
	return nil

}

func CreateEthClient(url string) (*ethclient.Client, error) {
	Client, err := ethclient.Dial(url)
	if err != nil {
		log.Printf("Error Creating ETH Client %+v", err)
		return nil, err
	}
	return Client, nil
}

func GetBalance(Client *ethclient.Client, address common.Address) (*big.Int, error) {
	balance, err := Client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Printf("Error: Can not get balance %+v", err)
		return nil, fmt.Errorf("error: Can not get balance %+v", err)
	}
	return balance, nil

}

func TransactionOpts(Client *ethclient.Client, key *keystore.Key, gaslimit uint64) (*bind.TransactOpts, error) {
	nonce, err := Client.PendingNonceAt(context.Background(), key.Address)

	if err != nil {
		log.Printf("Error: can Not get Nonce  %+v", err)
		return nil, fmt.Errorf("error: can Not get Nounce %+v", err)
	}
	// chainID, err := Client.NetworkID(context.Background())
	chainID, err := Client.ChainID(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	gasprice, err := Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Printf("Error getting gas price: %+v", err)
		return nil, fmt.Errorf("error getting gas price: %+v", err)

	}
	opts, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		log.Printf("Error creating transaction options: %+v", err)
		return nil, fmt.Errorf("error creating transaction options: %+v", err)
	}
	opts.GasPrice = gasprice
	opts.GasLimit = gaslimit
	opts.Nonce = big.NewInt(int64(nonce))

	return opts, nil
}

func DeployContract(Client *ethclient.Client, key *keystore.Key, gasLimit uint64, owners []common.Address, numOfConfirmation *big.Int) error {

	opts, err := TransactionOpts(Client, key, gasLimit)
	if err != nil {
		return fmt.Errorf("error creating transaction options: %+v", err)
	}

	a, tx, _, err := MultiSig.DeployMultiSig(opts, Client, owners, numOfConfirmation)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("+", strings.Repeat("-", 60), "+")
	fmt.Println(a.Hex())
	fmt.Println(tx.Hash().Hex())
	fmt.Println("+", strings.Repeat("-", 60), "+")

	return nil
}

func TransferWei(Client *ethclient.Client, toHex string, Weiamount int64, password, filePath string) (*types.Transaction, error) {
	chainID, err := Client.NetworkID(context.Background())
	if err != nil {
		log.Printf("Error: can Not get ChainID for this Network %+v", err)
		return nil, fmt.Errorf("error: can Not get ChainID for this Network %+v", err)
	}
	Key, err := DecryptKeyStoreOnFileSystem(password, filePath)
	if err != nil {
		log.Printf("Error: can Not get Decrypt keyStore %+v", err)
		return nil, fmt.Errorf("error: can Not get Decrypt keyStore %+v", err)
	}

	from := Key.Address
	to := common.HexToAddress(toHex)
	nonce, err := Client.PendingNonceAt(context.Background(), from)
	if err != nil {
		log.Printf("Error: can Not get nounce for this address %s -> %+v", from, err)
		return nil, fmt.Errorf("error: can Not get nounce for this address %s -> %+v", from, err)
	}
	amount := big.NewInt(Weiamount)
	gasprice, err := Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Printf("Error: can Not get gasPrice for this address %s -> %+v", from, err)
		return nil, fmt.Errorf("error: can Not get gasPrice for this address %s -> %+v", from, err)
	}
	Transaction := types.NewTransaction(nonce, to, amount, 21000, gasprice, nil)
	Transaction, err = types.SignTx(Transaction, types.NewEIP155Signer(chainID), Key.PrivateKey)
	if err != nil {
		log.Printf("Error: %+v", err)
		return nil, fmt.Errorf("error: %+v", err)
	}
	err = Client.SendTransaction(context.Background(), Transaction)
	if err != nil {
		log.Printf("Error: Transaction Failed %+v", err)
		return nil, fmt.Errorf("error: Transaction Faileds %+v", err)
	}
	return Transaction, nil
}

func CreatePrivateKey() (*ecdsa.PrivateKey, error) {
	ecdsaPrivateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Printf("Error: can Not create private Key")
		return nil, err
	}
	return ecdsaPrivateKey, nil
}

func GetEncodedPrivateKeyFromPrivateKey(ECDSAPrivatekey *ecdsa.PrivateKey) string {
	pvkByteSlice := crypto.FromECDSA(ECDSAPrivatekey)
	PrivateKeyEncoded := hexutil.Encode(pvkByteSlice)
	return PrivateKeyEncoded
}

func GetEncodedPublicKeyFromPrivateKey(ECDSAPrivatekey *ecdsa.PrivateKey) string {
	PublicKeyByteSlice := crypto.FromECDSAPub(&ECDSAPrivatekey.PublicKey)
	PublicKeyEncoded := hexutil.Encode(PublicKeyByteSlice)
	return PublicKeyEncoded
}

func GetAddressFromPrivateKey(ECDSAPrivatekey *ecdsa.PrivateKey) string {
	return crypto.PubkeyToAddress(ECDSAPrivatekey.PublicKey).Hex()
}

func CreateKeyStoreWithExistingKey(privateKey, password string) ([]byte, ecdsa.PublicKey, error) {
	privateKeyBytes, err := hexutil.Decode(privateKey)
	if err != nil {
		return nil, ecdsa.PublicKey{}, err
	}
	PriKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		return nil, ecdsa.PublicKey{}, err
	}
	address := crypto.PubkeyToAddress(PriKey.PublicKey)
	key := keystore.Key{
		Id:         uuid.New(),
		Address:    address,
		PrivateKey: PriKey,
	}
	keystoreBytes, err := keystore.EncryptKey(&key, password, keystore.StandardScryptN, keystore.StandardScryptP)
	if err != nil {
		return nil, ecdsa.PublicKey{}, err
	}
	return keystoreBytes, PriKey.PublicKey, nil

}

func DecryptKeyStoreOnFileSystem(password, filePath string) (*keystore.Key, error) {
	b, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("Error: %+v", err)
	}

	key, err := keystore.DecryptKey(b, password)
	if err != nil {
		log.Printf("Error: Can not Decrypt the password %+v", err)
		return nil, fmt.Errorf("error: Can not Decrypt the password %+v", err)
	}

	return key, nil
}

func CreateEncryptedKeyStore(password string) (string, error) {

	key := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)

	Acc, err := key.NewAccount(password)
	if err != nil {
		log.Printf("Error: Can not Create a Keystore %+v", err)
		return "", fmt.Errorf("error: Can not Create a Keystore %+v", err)

	}
	return Acc.URL.Path, nil
}

func CreateKeyStoreBytes(password string) ([]byte, ecdsa.PublicKey, error) {
	PriKey, err := crypto.GenerateKey()

	if err != nil {
		return nil, ecdsa.PublicKey{}, err
	}

	address := crypto.PubkeyToAddress(PriKey.PublicKey)
	key := keystore.Key{
		Id:         uuid.New(),
		Address:    address,
		PrivateKey: PriKey,
	}
	keystoreBytes, err := keystore.EncryptKey(&key, password, keystore.StandardScryptN, keystore.StandardScryptP)
	if err != nil {
		return nil, ecdsa.PublicKey{}, err
	}
	return keystoreBytes, PriKey.PublicKey, nil

}
