## Challange BackEnd Api

### Environment

##### Development
```
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

```
EXAMPLE URL

Prod:
https://challange-api.herokuapp.com/

Dev:
http://localhost:8080/
```
```
EXAMPLE SUCCESS RESPONSE

Code : 200

Body: Ok 
```

**Memory**

- Get Endpoint  
  Get data from an in-memory database.

```
URI
/in-memory
```
```
METHOD
GET
```
```
EXAMPLE URL
Prod:
https://challange-api.herokuapp.com/in-memory?key=active-tabs

Dev:
http://localhost:8080/in-memory?key=active-tabs

```
```
EXAMPLE SUCCESS RESPONSE

Code : 200

Payload:
 ```
```json       
{
    "key": "active-tabs",
    "value": "getir"
}
```


- Set Endpoint  
  Set data to an in-memory database.

```
URI
/in-memory/
```
METHOD
POST
```

EXAMPLE REQUEST


Payload:
```
```json
{
  "key": "active-tabs",
  "value": "getir"
}
```
```
EXAMPLE URL

Prod:
https://challange-api.herokuapp.com/in-memory/

Dev:
http://localhost:8080/in-memory/
```
```
EXAMPLE SUCCESS RESPONSE

Code : 201

Body:
```
 ```json       
{
  "key": "active-tabs",
  "value": "getir"
}
```


**Record**

- GetRecords Endpoint  
  Fetch data from records collection
```
URI
/records

METHOD
POST

EXAMPLE URL

Prod:
https://challange-api.herokuapp.com/records

Dev:
http://localhost:8080/records


EXAMPLE REQUEST 

Payload:
```
```json
{
  "startDate": "2016-01-26",
  "endDate": "2018-02-02",
  "minCount": 2700,
  "maxCount": 3000
}

```
```
EXAMPLE SUCCESS RESPONSE

Code : 200

Body:
```
```json
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


