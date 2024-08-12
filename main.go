package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aminghafoory/multisig2/controllers"
	"github.com/aminghafoory/multisig2/internal/database"
	"github.com/aminghafoory/multisig2/models"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/csrf"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
)

type Config struct {
	listenAddr  string
	RPCEndPoint string
	DBURL       string
	CSRF        struct {
		Key    string
		Secure bool
	}
}

func LoadEnvConfig() (Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return Config{}, err
	}

	listenAddr := os.Getenv("LISTEN_ADDR")
	if listenAddr == "" {
		return Config{}, err
	}

	dbURL := os.Getenv("DB_URL")
	if listenAddr == "" {
		return Config{}, err
	}

	RPCEndPoint := os.Getenv("RPC_ENDPOINT")
	if listenAddr == "" {
		return Config{}, err
	}

	CSRFKey := os.Getenv("CSRF")
	if listenAddr == "" {
		return Config{}, err
	}

	devEnv := os.Getenv("DEV")
	if devEnv == "" {
		return Config{}, fmt.Errorf("CSRFKey Not found in the .env file")

	}
	useSecureMode := func(isDev string) bool {
		if isDev == "false" {
			return true
		} else {
			return false
		}
	}(devEnv)

	c := Config{}
	c.listenAddr = listenAddr
	c.RPCEndPoint = RPCEndPoint
	c.CSRF.Key = CSRFKey
	c.CSRF.Secure = useSecureMode
	c.DBURL = dbURL

	return c, nil

}

func CreateEthClient(url string) (*ethclient.Client, error) {
	Client, err := ethclient.Dial(url)
	if err != nil {
		log.Printf("Error Creating ETH Client %+v", err)
		return nil, err
	}
	return Client, nil
}

func main() {

	config, err := LoadEnvConfig()
	if err != nil {
		log.Fatal("Could Not Read .env file")
	}

	CSRFMiddleware := csrf.Protect(
		[]byte(config.CSRF.Key),
		csrf.Secure(config.CSRF.Secure),
		csrf.CookieName("CSRF"),
		csrf.Path("/"),
		csrf.HttpOnly(true),
	)
	//ETH Client
	client, err := CreateEthClient(config.RPCEndPoint)
	if err != nil {
		log.Fatal("Could Not Create a eth client ", err)
	}
	bnumber, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("chainID:", bnumber)

	// nnumber, err := client.NetworkID(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(nnumber.Int64())

	//DB
	conn, err := sql.Open("pgx", config.DBURL)
	if err != nil {
		log.Fatal("Can NOT connect to the database")
	}
	db := database.New(conn)
	err = conn.Ping()
	if err != nil {
		log.Fatal("can not connect to DB")
	}

	UserC := controllers.Users{
		UserService: &models.UserService{
			DB: db,
		},
		SessionService: &models.SessionService{
			DB: db,
		},
	}
	ContractC := controllers.Contracts{
		ContractsService: &models.ContractsService{
			DB:        db,
			ETHClient: client,
		},
		WalletService: &models.WalletService{
			DB:        db,
			ETHClient: client,
		},
	}

	WalletC := controllers.WalletController{
		WalletService: &models.WalletService{
			DB:        db,
			ETHClient: client,
		},
	}

	UserMw := controllers.UserMiddleware{
		SessionService: &models.SessionService{
			DB: db,
		},
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(CSRFMiddleware)
	r.Use(UserMw.SetUser)

	r.Group(func(r chi.Router) {
		r.Use(UserMw.ReqireUser)
		r.Get("/users/me", UserC.CurrentUser)

		r.Get("/contracts/new", ContractC.NewContractPage)
		r.Get("/contracts", ContractC.UserContracts)
		r.Post("/contracts", ContractC.DeployContract)
		r.Get("/contracts/{contractID}", ContractC.UseContract)
		r.Get("/contracts/{contractID}/delete", ContractC.DeleteContract)

		r.Get("/users/wallets", WalletC.UserWallets)
		r.Get("/users/wallets/import", WalletC.ImportExistingWallet)
		r.Post("/users/wallets/import", WalletC.ProccessImportExistingWallet)
		r.Get("/users/wallets/new", WalletC.CreateNewWallet)
		r.Post("/users/wallets/new", WalletC.ProccssCreateNewWallet)

		r.Get("/wallets/{walletID}/delete", WalletC.DeleteWallet)

		r.Get("/wallets/{walletID}/prikey", WalletC.ShowWalletPrivateKey)
		r.Post("/wallets/{walletID}/prikey", WalletC.ProcessShowWalletPrivateKey)

	})

	r.Get("/", controllers.Home)
	r.Post("/sign-in/", UserC.ProccessSignIn)
	r.Post("/users", UserC.Create)
	r.Get("/sign-in", UserC.NewSignInPage)
	r.Get("/sign-out", UserC.SignOutUser)
	r.Post("/sign-in", UserC.ProccessSignIn)
	r.Get("/sign-up", UserC.New)
	r.Get("/public/*", public().ServeHTTP)

	fmt.Printf("Server started on %s\n", config.listenAddr)
	http.ListenAndServe(config.listenAddr, r)
}
