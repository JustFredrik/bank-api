# bank-api
 This is a Mock RESTful bank api written in GOlang. It parses and loads mock data from a [camt053 file](https://www.sepaforcorporates.com/swift-for-corporates/a-practical-guide-to-the-bank-statement-camt-053-format/) and uses that for the responses. The API uses a basic API key Authentication method with bearer tokens in the header.

The API Has the ability to fetch accounts, list accounts, fetch transactions and list the transactions of a given account. Depending on the request resource it will require different role privalidges. 

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

## API Endpoints
This section lists all valid endpoints in the API along with required header fields and response examples.
### GET /accounts
To list accounts in the API you can call the /accounts end point. You will recieve a 200 OK if you have authorization to access the resource otherwise you will get a 401 Unauthorized error. If you for some reason have access but the server for some reason is u

Required Role: ADMIN


```
```

## Testing
The code base has partial code coverage with most focus being on that the end product, the end-points work as expected.

To run all tests run 
```cli
go test ./...
```
in the project root directory.