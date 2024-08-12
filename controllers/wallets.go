package controllers

import (
	"context"
	"fmt"
	"net/http"

	ctx "github.com/aminghafoory/multisig2/context"
	"github.com/aminghafoory/multisig2/models"
	"github.com/aminghafoory/multisig2/utils"
	ImportWallet "github.com/aminghafoory/multisig2/views/importWallet"
	NewWallet "github.com/aminghafoory/multisig2/views/newWallet"
	ShowPrivateKey "github.com/aminghafoory/multisig2/views/showPrivatekey"
	"github.com/aminghafoory/multisig2/views/wallets"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type WalletController struct {
	WalletService *models.WalletService
}

func (wc *WalletController) UserWallets(w http.ResponseWriter, r *http.Request) {
	user := ctx.User(r.Context())
	dbWallets, err := wc.WalletService.UserWallets(user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	c := wallets.UserWallets("wallets", r, dbWallets)
	err = c.Render(context.Background(), w)
	if err != nil {
		http.Error(w, "error in parsing the template", http.StatusInternalServerError)
		return
	}

}

func (wc *WalletController) CreateNewWallet(w http.ResponseWriter, r *http.Request) {
	c := NewWallet.NewWallet("new wallet", r)
	c.Render(context.Background(), w)

}

func (wc *WalletController) ProccssCreateNewWallet(w http.ResponseWriter, r *http.Request) {

	type data struct {
		name             string
		keystorePassword string
	}
	formdata := data{}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "can not parse form", http.StatusBadRequest)
		return
	}
	name := r.FormValue("name")
	password := r.FormValue("password")

	formdata.name = name
	formdata.keystorePassword = password

	user := ctx.User(r.Context())

	err = wc.WalletService.CreateNewWallet(user.ID, formdata.keystorePassword, formdata.name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, "/users/wallets", http.StatusFound)

}

func (wc *WalletController) ImportExistingWallet(w http.ResponseWriter, r *http.Request) {
	c := ImportWallet.NewWallet("import wallet", r)
	err := c.Render(context.Background(), w)
	if err != nil {
		http.Error(w, "error in parsing the template", http.StatusInternalServerError)
		return
	}

}

func (wc *WalletController) ProccessImportExistingWallet(w http.ResponseWriter, r *http.Request) {
	type data struct {
		name             string
		privateKey       string
		keystorePassword string
	}
	formdata := data{}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "can not parse form", http.StatusBadRequest)
		return
	}
	name := r.FormValue("name")
	privateKey := r.FormValue("private-key")
	password := r.FormValue("password")

	if len(privateKey) == 64 {
		privateKey = fmt.Sprint("0x", privateKey)
	}
	formdata.name = name
	formdata.privateKey = privateKey
	formdata.keystorePassword = password

	user := ctx.User(r.Context())

	err = wc.WalletService.ImportWallet(user.ID, formdata.privateKey, formdata.keystorePassword, formdata.name)
	if err != nil {
		http.Error(w, "something went wrong : error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/users/wallets", http.StatusFound)

}

func (wc *WalletController) DeleteWallet(w http.ResponseWriter, r *http.Request) {
	user := ctx.User(r.Context())

	walletID := chi.URLParam(r, "walletID")

	walletUUID, err := uuid.Parse(walletID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = wc.WalletService.DeleteWallet(walletUUID, user.ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/users/wallets", http.StatusFound)

}

func (wc *WalletController) ShowWalletPrivateKey(w http.ResponseWriter, r *http.Request) {
	c := ShowPrivateKey.ShowWalletPrivateKey("show Private Key", r)
	err := c.Render(context.Background(), w)
	if err != nil {
		http.Error(w, "error in parsing the template", http.StatusInternalServerError)
		return
	}
}

func (wc *WalletController) ProcessShowWalletPrivateKey(w http.ResponseWriter, r *http.Request) {
	user := ctx.User(r.Context())

	walletID := chi.URLParam(r, "walletID")

	walletUUID, err := uuid.Parse(walletID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = r.ParseForm()
	if err != nil {
		http.Error(w, "can not parse form", http.StatusBadRequest)
		return
	}

	password := r.FormValue("password")

	fmt.Println(password, user, walletUUID)
	// fetch the keystore
	keystoreBytes, err := wc.WalletService.GetUserWalletKeystoreByID(walletUUID, user.ID)
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

	w.Write([]byte(utils.GetEncodedPrivateKeyFromPrivateKey(key.PrivateKey)))
}
