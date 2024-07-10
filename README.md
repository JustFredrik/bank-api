# bank-api
 This is a Mock RESTful bank api written in GOlang. It parses and loads mock data from a [camt053 file](https://www.sepaforcorporates.com/swift-for-corporates/a-practical-guide-to-the-bank-statement-camt-053-format/) and uses that for the responses. The API uses a basic API key Authentication method with bearer tokens in the header. The API has three levels of access privalidge, those being Admin, Account and Any.

## The API service
The API Has the ability to fetch accounts, list accounts, fetch transactions and list the transactions of a given account. Depending on the request resource it will require different role privalidges. 

### GET /accounts
To list accounts in the API you can call the /accounts end point.
Required Role: ADMIN
Requires 
```

## Testing
The code base has partial code coverage with most focus being on that the end product, the end-points work as expected.

To run all tests run 
```cli
go test ./...
```
in the project root directory.