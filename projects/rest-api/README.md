# REST API - Books Management 

### Different Operations

- GET request
```bash

$ curl -s localhost:8080/books

[
    {
        "id": "1",
        "title": "Harry Potter: Deathly Hollows P1",
        "author": "JK Rowling",
        "quantity": 2
    },
    {
        "id": "2",
        "title": "The Great Gatsby",
        "author": "F. Scott Fitzgerald",
        "quantity": 5
    },
...
```
- POST request
```bash
$ curl localhost:8080/books -H 'Content-Type: application/json' -d @body.json -X POST    

{
    "id": "4",
    "title": "Harry Potter: Philosopher's Stone",
    "author": "JK Rowling",
    "quantity": 2
}
```
- GET request (By ID)
```bash

$ curl -s localhost:8080/books/2

{
    "id": "2",
    "title": "The Great Gatsby",
    "author": "F. Scott Fitzgerald",
    "quantity": 5
}

```

### Random TidBits

- What is a Payload in an API?
The payload of an API is the data you are interested in transporting to the server when you make an API request. Simply put, it is the body of your HTTP request and response message.