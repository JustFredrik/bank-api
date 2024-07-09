package db

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/justfredrik/bank-api/internal/camt053"
)

type BankData struct {
	Accounts           map[uint64]*Account
	TotalAccounts      uint64
	LoadedTransactions map[string]bool
}

type IDataBase interface {
	AccountsExists(accountId uint64) bool
	GetAccounts(perPage uint16, page uint64) (AccountsResponse, error)
	GetAccount(accountId uint64) (*Account, error)
	CreateAccount(camtAcc *camt053.Account) (*Account, error)
	PatchAccount(accountId uint64, patch *camt053.Account) (*Account, error)
	GetAccountTransactions(accountId uint64) (TransactionsResponse, error)
}

type Account struct {
	Account      camt053.Account   `json:"account"`
	Balances     []camt053.Balance `json:"balances"`
	Transactions []camt053.Entry   `json:"-"`
}

type AccountsResponse struct {
	Accounts   []*Account `json:"accounts"`
	TotalCount int        `json:"totalCount"`
	Page       int        `json:"page"`
	PerPage    int        `json:"perPage"`
}

type TransactionsResponse struct {
	Transactions []*camt053.Entry `json:"transactions"`
	TotalCount   int              `json:"totalCount"`
	Page         int              `json:"page"`
	PerPage      int              `json:"perPage"`
}

type LimitedAccount struct {
}

func (db BankData) AccountExists(accountId uint64) bool {
	_, alreadyExists := DB.Accounts[accountId]
	return alreadyExists
}

func (db BankData) CreateAccount(camtAcc *camt053.Account) (*Account, error) {

	accountId := (*camtAcc).GetId()

	if DB.AccountExists(accountId) {
		return nil, errors.New("trying to create account that already exists")
	}

	DB.TotalAccounts++

	acc := Account{
		Account:      *camtAcc,
		Transactions: make([]camt053.Entry, 0),
	}

	DB.Accounts[accountId] = &acc

	return &acc, nil
}

func (db BankData) PatchAccount(accountId uint64, patch *camt053.Account) error {

	if !DB.AccountExists(accountId) {
		return errors.New("trying to patch account that does not exist")
	}

	// Do stuff

	return nil
}

func (db BankData) GetAccounts(perPage uint16, page uint64) (AccountsResponse, error) {

	// While this may be slow while itterating over a large map of accounts
	// This is just a moc and in prod you would use and query a real db not this
	accounts := AccountsResponse{
		Accounts: make([]*Account, 0),
	}

	// Populate slice with map values
	for _, acc := range db.Accounts {
		accounts.Accounts = append(accounts.Accounts, acc)
	}

	// Populate accounts with correct data
	accounts.TotalCount = len(accounts.Accounts)
	accounts.PerPage = accounts.TotalCount

	return accounts, nil
}

func (db BankData) GetAccount(accountId uint64) (*Account, error) {
	if account, ok := db.Accounts[accountId]; ok {
		return account, nil
	}
	return nil, errors.New("account not found")
}

var DB BankData = BankData{
	Accounts:           make(map[uint64]*Account),
	LoadedTransactions: make(map[string]bool),
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

func LoadCamt053(data camt053.Document) error {

	// jsonData, err := json.MarshalIndent(data, "", "	")
	// fmt.Println(string(jsonData))

	var camtAcc camt053.Account = data.BankStatement.Statement.Account
	_, err := DB.CreateAccount(&camtAcc)
	if err != nil {
		return err
	}

	for _, entry := range *(data.BankStatement.Statement.Entries) {
		if _, ok := DB.LoadedTransactions[*entry.Reference]; !ok {

		}
	}

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
