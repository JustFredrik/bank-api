package db

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/justfredrik/bank-api/internal/camt053"
)

type BankData struct {
	Accounts      map[uint64]Account
	TotalAccounts uint64
	Transactions  map[uint64]([]camt053.TransactionDetail)
}

type IDataBase interface {
	AccountsExists(accountId uint64) bool
	GetAccounts(perPage uint16, page uint64) (Accounts, error)
	GetAccount(accountId uint64) (*Account, error)
	CreateAccount(camtAcc *camt053.Account) (*Account, error)
	PatchAccount(accountId uint64, patch *camt053.Account) (*Account, error)
	GetAccountTransactions(accountId uint64) (Transactions, error)
}

type Account struct {
	Account  *camt053.Account  `json:"account"`
	Balances []camt053.Balance `json:"balances"`
}

type Accounts struct {
	Accounts   []*Account `json: "accounts"`
	TotalCount int        `json: "totalCount"`
	Page       int        `json:"page"`
	PerPage    int        `json:"perPage"`
}

type Transactions struct {
	Transactions []*Transactions `json:"transactions"`
	TotalCount   int             `json: "totalCount"`
	Page         int             `json:"page"`
	PerPage      int             `json:"perPage"`
}

type LimitedAccount struct {
}

func (db BankData) AccountExists(accountId uint64) bool {
	_, alreadyExists := DB.Accounts[accountId]
	return alreadyExists
}

func (db BankData) CreateAccount(camtAcc *camt053.Account) error {

	accountId := (*camtAcc).GetId()

	if DB.AccountExists(accountId) {
		return errors.New("trying to create account that already exists")
	}

	DB.TotalAccounts++
	DB.Accounts[accountId] = Account{
		Account: camtAcc,
	}

	return nil
}

func (db BankData) PatchAccount(accountId uint64, patch *camt053.Account) error {

	if !DB.AccountExists(accountId) {
		return errors.New("trying to patch account that does not exist")
	}

	// Do stuff

	return nil
}

func (db BankData) GetAccounts(perPage uint16, page uint64) (Accounts, error) {

	// While this may be slow while itterating over a large map of accounts
	// This is just a moc and in prod you would use and query a real db not this
	accounts := Accounts{
		Accounts: make([]*Account, 0),
	}

	// Populate slice with map values
	for _, acc := range db.Accounts {
		accounts.Accounts = append(accounts.Accounts, &acc)
	}

	// Populate accounts with correct data
	accounts.TotalCount = len(accounts.Accounts)
	accounts.PerPage = accounts.TotalCount

	return accounts, nil
}

func (db BankData) GetAccount(accountId uint64) (*Account, error) {
	if account, ok := db.Accounts[accountId]; ok {
		return &account, nil
	}
	return nil, errors.New("account not found")
}

var DB BankData = BankData{
	Accounts:     make(map[uint64]Account),
	Transactions: make(map[uint64]([]camt053.TransactionDetail)),
}

func ParseLocalCamt053(path string) (camt053.Document, error) {

	var data camt053.Document

	xmlFile, err := os.Open(path)
	if err != nil {
		return data, err
	}
	defer xmlFile.Close()

	byteData, _ := io.ReadAll(xmlFile)

	if err := xml.Unmarshal(byteData, &data); err != nil {

		return data, err
	}

	return data, err
}

func LoadCamt053(data camt053.Document) (err error) {

	jsonData, err := json.MarshalIndent(data, "", "	")
	fmt.Println(string(jsonData))

	var camtAcc camt053.Account = data.BankStatement.Statement.Account
	DB.CreateAccount(&camtAcc)

	return nil
}

func InitializeLocalMockData() (err error) {
	data, err := ParseLocalCamt053("./data/camt053.xml")
	if err != nil {
		fmt.Println("error here")
		return err
	}

	fmt.Println("Parsed local data...")

	if err = LoadCamt053(data); err != nil {
		return err
	}

	fmt.Println("Loaded local data!")

	return nil
}
