# Execution guide


### Run and process:

- Open the Terminal o CMD
- In the current path of this README file, run the commands:
````
go get	
````
- After, run the follow command
````
go run server.go
````

<br>

#### GraphQL request:

- Go to localhost:8080 to execute the follow queries

<br>

> Insert/Create:

````
mutation createBook {
  createBook (
    input: { title:"Libro 4", name:"Nombre 4", userId:"4"}
  ) {
    author {
      id
      name
    }
    id
    title
  }
}
````
<br>

> Find/Select
````
query findBooks {
  books {
    id
    title
    author {
      id
      name
    }
  }
}
````
