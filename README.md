# justfredrik/bank-api
 This is a Mock RESTful bank api written in GO. It parses and loads mock data from a [camt053 file](https://www.sepaforcorporates.com/swift-for-corporates/a-practical-guide-to-the-bank-statement-camt-053-format/) and uses that for the responses. The API uses a basic API key Authentication method with bearer tokens in the header.

The API Has the ability to fetch accounts, list accounts, fetch transactions and list the transactions of a given account. Depending on the request resource it will require different role privalidges. 

# Setup
First, if you do not already have GO installed then you have to download and install GO, you can do so [here](https://go.dev/dl/).

Secondly, run `go mod download` to download all dependencies.

Thirdly, create a `.env` in the project root directory. This ENV file should contain the following:
```sh
AUTH_MODE=debug
PORT=8080
GIN_MODE=debug
PROJECT_DIR=[PATH TO PROJECT ROOT DIR]
```

Finally, to run the project run `go run cmd/main.go` in the projects root directory.


# Overview
This section goes through some asects of the project layout and details of how it works and how to interact with it.
## Authorization
A basic API key system is in place with three levels of access privalidge. These levels are: `Admin`, `Account` and `Any`.
In order to get access to the service you will need to include a valid API key with the correct access privalidge for the requested resource. 

Example header with API key attatched as a bearer token:
```json
{
    "User-Agent" : "PostmanRuntime/7.40.0",
    "Accept" : "*/*",
    "Connection" : "keep-alive",
    "host" : "localhost:8080",
    "Authorization" : "Bearer E77g8v16Au8fkLvjf1yf5f4NfLneC9EK",
}
```

### Test Keys
The API will create three predicatable API Keys with the ENV variable `AUTH_DEBUG` set to `true`.
These three have the following Authority and associated AccountIds
| Auth Role  | Accound Id      | token                            |
|:-----------|:----------------|:---------------------------------|
| Admin      | 0               | 8iqmm8vmFGHyA4ikLBBKcrn36kfggANM |
| Account    | 54400001111     | E77g8v16Au8fkLvjf1yf5f4NfLneC9EK |
| Account    | 13371337984     | kN7fgeBax424gcEEFnkFe3cqd4rfc3Mg |

The Mock data contains Account data for account `54400001111`, The API key associated with accountId `13371337984` is to test account resource access rules.

## Errors
Any error response from the API will (other than the HTTP status) have a body with JSON containing a message and error key-value pairs.

#### Example Error
```json
{
    "error": "Unauthorized", 
    "message": "Your API key is not authorized to access the requested resource"
}
```

## API Endpoints
This section lists all valid endpoints in the API along with required header fields and response examples. You will recieve a 200 OK code if you have authorization to access the resource otherwise you will get a 401 Unauthorized error. If you have access but the server is unable to find the requested resource the server will return a 404 Not Found error.


### GET /ping
Anyone with a valid API key and call the `/ping` endpoint. The server will respond with a pong if the API key is valid.

|   |   |
|---|---|
|__Required Role__| Any |

#### example resonse
```
{
    "message" : "pong"
}
```


### GET /accounts
To list accounts in the API you can call the `/accounts` endpoint. 

|   |   |
|---|---|
|__Required Role__| Admin |


### GET /accounts/:accountId
Fetching accounts can be done by specifying an account id (accountId) at the `/accounts/:accountId` endpoint. Your API key needs to have the Admin role or be associated with the requested accoundId.

|   |   |
|---|---|
|__Required Role__| Admin or Account *(with matching accountId)* |
| __accountId type__ | *uint64* |

 
### GET /accounts/:accountId/transactions
Fetching transactions for a given account can be done by specifying an account id (accountId) at the `/accounts/:accountId/transactions` endpoint. Your API key needs to have the Admin role or be associated with the requested accoundId.

|   |   |
|---|---|
|__Required Role__| Admin or Account *(with matching accountId)* |
| __accountId type__ | *uint64* |


### GET /accounts/:accountId/transactions/transactionRef
Fetching a specific transaction for a given account can be done by specifying an account id (accountId) followed by `/transactions/`, followed by a transaction reference (transactionRef) at the `/accounts/:accountId/transactions` endpoint. Your API key needs to have the Admin role or be associated with the requested accoundId.
|   |   |
|---|---|
|__Required Role__| Admin or Account *(with matching accountId)* |
| __accountId type__ | *uint64* |
| __transactionId type__ | *string* |



## Testing
The code base has partial code coverage with most focus being on that the end product, the end-points, work as expected.
  
To run all tests run 
```cli
go test ./...
```
in the project root directory.


# Future Work
There are several areas which could be further improved. This section lists some areas which I would like to improve if I had more time.

The API only supports the R (read) in the CRUD acronym. For future work, full support for creating, reading, updating and deleteing resources would increase the quality of the project.

Adding a real database such as PostgreSQL for datastorage is something that would be good as a future feature. Adding a real database would streamline implementation of proper pagination by including `LIMIT number_of_rows OFFSET offset_value` in the SQL requests.

Full code coverage of the codebase would be a good addition, as it stands local packages have very limited coverage.

Expanding and validating that camt053 are loaded properly and follows the spec would be good. An extensive test suite with varied camt053 data would be good to validate the marshaling and unmarshaling of camt053 data. 

This repo is my first lines of GO, and as such I may not have followed all idioms and best practices. When I've gained more expereince in GO it would be benefitial to go back through the code to review it and perhaps refactor to better reflect the idioms and best practices of GO.
