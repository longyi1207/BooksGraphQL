# BooksGraphQL

## Description
BooksGraphQL is a practice project that showcases creating a GraphQL service in Go. This service allows users to easily query information about authors and books.

This project was inspired by an example provided by Web Dev Simplified on [YouTube](https://www.youtube.com/watch?v=ZQL7tL2S0oQ).

## Usage

1. Start the service:
   ```bash
   go run *.go
    ```
2. Open your browser and navigate to:
    ```bash
   http://localhost:8080/graphql
    ```
3. Enter queries to fetch data. For example:
* Fetch all authors with their ID, name, and the names of the books they wrote:
    ```bash
    {
      authors 
       {
        id
        name
        books 
         {
          name
         }
      }
    }

    ```
* If you're a fan of "The Two Towers", you can find out its author and discover other books by the same author:
    ```bash
    {
    bookByName(name: "The Two Towers")
      {
        authorId
      }
    }
    ```
* The result might be an ID, for example, 2. You can then query:
    ```bash
    {
    authorById(id: 2) 
      {
        books {
          name
        }
      }
    }
    ```