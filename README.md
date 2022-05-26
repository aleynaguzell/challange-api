## Challange BackEnd Api

### Environment

##### Development
```json
// Verify that you've installed Go
$ go version

// Clone the project
$ git clone github.com/aleynaguzell/challange-api
$ cd challange-api

// Run test
$ go test -v ./tests -run Test

// Run project
$ go run main.go
```
### API Document



**HealthCheck**

- HealthCheck Endpoint  
  API ping endpoint barely checks that the API is running and is accessible.

```json
EXAMPLE URL

Prod:
https://challange-api.herokuapp.com/in-memory?key=active-tabs

Dev:
http://localhost:8081/
```
```json
EXAMPLE SUCCESS RESPONSE

Code : 200

Body: Ok
```

**Memory**

- Get Endpoint  
  Get data from an in-memory database.

```json
URI
/in-memory
```
```json
METHOD
GET
```
```json
EXAMPLE URL
Prod:
https://challange-api.herokuapp.com/in-memory?key=active-tabs

Dev:
http://localhost:8081/in-memory?key=active-tabs

```
```json
EXAMPLE SUCCESS RESPONSE

Code : 200

Payload:
        
{
    "key": "active-tabs",
    "value": "getir"
}
```


- Set Endpoint  
  Set data to an in-memory database.

```json
URI
/in-memory/
```
```json
EXAMPLE REQUEST

Payload:

{
  "key": "active-tabs",
  "value": "getir"
}
```
```json
EXAMPLE URL

Prod:
https://challange-api.herokuapp.com/in-memory/

Dev:
http://localhost:8081/in-memory/
```
```json
EXAMPLE SUCCESS RESPONSE

Code : 201

Body:
        
{
  "key": "active-tabs",
  "value": "getir"
}
```


**Record**

- GetRecords Endpoint  
  Fetch data from records collection
```json
URI
/records
```
```json
METHOD
POST
```
```json
EXAMPLE URL

Prod:
https://challange-api.herokuapp.com/records

Dev:
http://localhost:8081/records
```
```json
EXAMPLE REQUEST 

Payload:

{
  "startDate": "2016-01-26",
  "endDate": "2018-02-02",
  "minCount": 2700,
  "maxCount": 3000
}
```
```json
EXAMPLE SUCCESS RESPONSE

Code : 200

Body:

{
  "code":0,
  "msg":"Success",
  "records": [
    {
      "key":"TAKwGc6Jr4i8Z487",
      "createdAt":"2017-01-28T01:22:14.398Z",
      "totalCount":2800
    },
    {
      "key":"NAeQ8eX7e5TEg7oH",
      "createdAt":"2017-01-27T08:19:14.135Z",
      "totalCount":2900
    }
  ]
}
```


