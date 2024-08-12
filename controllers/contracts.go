package controllers

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"

	ctx "github.com/aminghafoory/multisig2/context"
	MultiSig "github.com/aminghafoory/multisig2/contracts/abigenBuild"
	"github.com/aminghafoory/multisig2/internal/database"
	"github.com/aminghafoory/multisig2/models"
	"github.com/aminghafoory/multisig2/utils"
	"github.com/aminghafoory/multisig2/views/contracts"
	"github.com/aminghafoory/multisig2/views/newContract"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type Contracts struct {
	ContractsService *models.ContractsService
	WalletService    *models.WalletService
}

func (c Contracts) DeployContract(w http.ResponseWriter, r *http.Request) {
	user := ctx.User(r.Context())

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var data struct {
		ContractName              string
		Owners                    []common.Address
		NumOfConfirmationRequired *big.Int
		walletPassword            string
		walletID                  uuid.UUID
	}

	data.walletPassword = r.FormValue("password")
	data.ContractName = r.FormValue("contract-name")
	nconfirm, err := strconv.ParseInt(r.FormValue("nconfirm"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data.NumOfConfirmationRequired = big.NewInt(nconfirm)

	ownerStr := r.FormValue("owners") //TODO fix
	owners := strings.Split(ownerStr, ",")
	ownersList := []common.Address{}
	for _, owner := range owners {
		ownersList = append(ownersList, common.HexToAddress(owner))
	}

	data.Owners = ownersList

	walletIDString := r.FormValue("wallet_id")

	walletID, err := uuid.Parse(walletIDString)

	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data.walletID = walletID

	fmt.Println(data)

	keyStoreString, err := c.ContractsService.DB.GetUserWalletByID(context.Background(), database.GetUserWalletByIDParams{
		UserID:   user.ID,
		WalletID: data.walletID,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	key, err := keystore.DecryptKey([]byte(keyStoreString), data.walletPassword)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	TxOpts, err := utils.TransactionOpts(c.ContractsService.ETHClient, key, 30000000)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	contractAddress, _, _, err := MultiSig.DeployMultiSig(TxOpts, c.ContractsService.ETHClient, data.Owners, data.NumOfConfirmationRequired)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	contract, err := c.ContractsService.DB.CreateaContract(context.Background(), database.CreateaContractParams{
		UserID:       user.ID,
		CreatedAt:    time.Now(),
		ContractID:   contractAddress.Hex(),
		ContractName: data.ContractName,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, contract)
}

func (c Contracts) NewContractPage(w http.ResponseWriter, r *http.Request) {
	comp := newContract.NewContractPage("newContract", r)
	comp.Render(context.Background(), w)
}

func (c Contracts) UserContracts(w http.ResponseWriter, r *http.Request) {
	user := ctx.User(r.Context())

	dbContracts, err := c.ContractsService.UserContracts(user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	comp := contracts.UserContracts("my contracts", r, dbContracts)
	comp.Render(context.Background(), w)
}

func (c Contracts) UseContract(w http.ResponseWriter, r *http.Request) {
	user := ctx.User(r.Context())
	fmt.Println(user) // TODO
	contractID := chi.URLParam(r, "contractID")
	fmt.Println(contractID) // TODO
	contract, err := MultiSig.NewMultiSig(common.HexToAddress(contractID), c.ContractsService.ETHClient)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	walletID := "8d4d1095-78ae-4e63-b32e-84031d891183"

	walletUUID, err := uuid.Parse(walletID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	password := "amin1234"

	fmt.Println(password, user, walletUUID)
	// fetch the keystore
	keystoreBytes, err := c.WalletService.GetUserWalletKeystoreByID(walletUUID, user.ID)
	if err != nil {
		http.Error(w, "can not get wallet", http.StatusBadRequest)
		return
	}

	// decrypt keystore
	key, err := keystore.DecryptKey([]byte(keystoreBytes), password)
	if err != nil {
		http.Error(w, "wrong password", http.StatusBadRequest)
		return
	}

	txopts, err := utils.TransactionOpts(c.ContractsService.ETHClient, key, 300000)
	if err != nil {
		http.Error(w, "wrong password", http.StatusBadRequest)
		return
	}

	num, err := contract.SubmitTransaction(txopts, common.HexToAddress("0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC"), big.NewInt(1000000000000), []byte("tip"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte(fmt.Sprint(num)))
}

func (c Contracts) DeleteContract(w http.ResponseWriter, r *http.Request) {
	user := ctx.User(r.Context())
	contractID := chi.URLParam(r, "contractID")
	if len(contractID) != 42 {
		http.Error(w, "invalid contract Address", http.StatusBadRequest)
		return
	}
	err := c.ContractsService.DeleteContract(contractID, user.ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, "/contracts", http.StatusFound)

}
