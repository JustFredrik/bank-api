// package db is a local mock database.
package db

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/justfredrik/bank-api/internal/camt053"
)

// IDataBase represents the Mock Database
type IDataBase interface {
	AccountsExists(accountId uint64) bool
	GetAccounts(perPage uint16, page uint64) (AccountsResponse, error)
	GetAccount(accountId uint64) (*Account, error)
	CreateAccount(camtAcc *camt053.Account) (*Account, error)
	GetAccountTransactions(accountId uint64) (TransactionsResponse, error)
	GetAccountTransaction(accountId uint64, transactionRef string) (TransactionsResponse, error)
}

// BankData is used as the root of the database. Implements IDataBase
type BankData struct {
	Accounts           map[uint64]*Account
	TotalAccounts      uint64
	LoadedTransactions map[string]bool
}

// Account stores an account along with it's balances and transactions.
type Account struct {
	Account      camt053.Account          `json:"account"`
	Balances     []camt053.Balance        `json:"balances"`
	Transactions map[string]camt053.Entry `json:"-"`
}

// AccountResponse is the format for /accounts request responses.
type AccountsResponse struct {
	Accounts   []*Account `json:"accounts"`
	TotalCount int        `json:"totalCount"`
	Page       int        `json:"page"`
	PerPage    int        `json:"perPage"`
}

// TransactionsResponse is the format for /transactions request responses.
type TransactionsResponse struct {
	Transactions []*camt053.Entry `json:"transactions"`
	TotalCount   int              `json:"totalCount"`
	Page         int              `json:"page"`
	PerPage      int              `json:"perPage"`
}

// AccountExists checks if an account exists in the database.
func (db BankData) AccountExists(accountId uint64) bool {
	_, alreadyExists := DB.Accounts[accountId]
	return alreadyExists
}

// CreateAccount creates an account in the database.
func (db BankData) CreateAccount(camtAcc *camt053.Account) (*Account, error) {

	accountId := (*camtAcc).GetId()

	if DB.AccountExists(accountId) {
		return nil, errors.New("trying to create account that already exists")
	}

	DB.TotalAccounts++

	acc := Account{
		Account:      *camtAcc,
		Transactions: make(map[string]camt053.Entry, 0),
	}

	DB.Accounts[accountId] = &acc

	return &acc, nil
}

// GetAccounts gets the list of accounts in the database. (pagination is not implemented)
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

// GetAccount gets a specific account from the database.
func (db BankData) GetAccount(accountId uint64) (*Account, error) {
	if account, ok := db.Accounts[accountId]; ok {
		return account, nil
	}
	return nil, errors.New("account not found")
}

// GetAccountTransactions gets a list of an accounts transactions from the database.
func (db BankData) GetAccountTransactions(accountId uint64) (*TransactionsResponse, error) {
	transactions := []*camt053.Entry{}

	// Fetch Account
	account, err := db.GetAccount(accountId)
	if err != nil {
		return nil, nil
	}

	// Convert Map data to slice since we don't use a real DB
	for _, transaction := range account.Transactions {
		transactions = append(transactions, &transaction)
	}

	totalCount := len(transactions)

	return &TransactionsResponse{
		Transactions: transactions,
		TotalCount:   totalCount,
		Page:         1,
		PerPage:      totalCount,
	}, nil
}

// GetAccountTransaction gets a specific transaction for an ccount from the database.
func (db BankData) GetAccountTransaction(accountId uint64, transactionRef string) (*camt053.Entry, error) {

	// Fetch Account
	account, err := db.GetAccount(accountId)
	if err != nil {
		return nil, errors.New("unable to fetch account data")
	}

	transaction, ok := account.Transactions[transactionRef]
	if !ok {
		return nil, errors.New("transaction not found")
	}

	return &transaction, nil
}

// Instance of the BankData Database used as the database in the project.
var DB BankData = BankData{
	Accounts:           make(map[uint64]*Account),
	LoadedTransactions: make(map[string]bool),
}

// ParseLocalCamt053 opens and unmarshals a camt053 document.
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

// Loads unmarshaled camt053 into the database.
func LoadCamt053(data camt053.Document) error {

	// Load Account data and Create Account
	var camtAcc camt053.Account = data.BankStatement.Statement.Account
	account, err := DB.CreateAccount(&camtAcc)
	if err != nil {
		return err
	}

	// Load Transactions (Entries) into Account data struct
	for _, entry := range *(data.BankStatement.Statement.Entries) {

		// Convert Ref to URL friendly string
		{
			URLSafeRef := convertEntryRef(*entry.Reference)
			entry.URLReference = &URLSafeRef
		}

		// Add transaction if no duplicate exists
		if _, ok := DB.LoadedTransactions[*entry.URLReference]; !ok {
			account.Transactions[*entry.URLReference] = entry
		}
	}

	return nil
}

// InitializeLocalMockData loads the /data/camt053.xml file for testing.
func InitializeLocalMockData() (err error) {
	if localMockIsInitialized == true {
		return nil
	}
	data, err := ParseLocalCamt053(os.Getenv("PROJECT_DIR") + "/data/camt053.xml")
	if err != nil {
		return err
	}

	//fmt.Println("Parsed local data...")

	if err = LoadCamt053(data); err != nil {
		return err
	}

	fmt.Println("Loaded local data!")
	localMockIsInitialized = true

	return nil
}

// Converts references to URL friendly references.
func convertEntryRef(rawRef string) string {
	// These strings are not good to have in a resource name/Id/Ref in an URL
	unwantedCharacters := []string{";", "/", "?", ":", "@", "=", "&", "\"",
		"<", ">", "#", "%", "{", "}", "|", "\\", "^", "~", "[", "]", "`", " "}

	// Tread spaces as -
	resRef := strings.ReplaceAll(rawRef, " ", "-")

	// Remove all unwanted Characters
	for _, char := range unwantedCharacters {
		resRef = strings.ReplaceAll(rawRef, char, "")
	}

	return resRef
}

// Stores if mock data has already been initialized or not
var localMockIsInitialized bool = false
