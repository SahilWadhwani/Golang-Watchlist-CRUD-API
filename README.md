# Golang Watchlist CRUD API

A Golang API that allows users to manage watchlists through basic CRUD (Create, Read, Update, Delete) operations. Users can create watchlists with a unique ID, update them, retrieve them, and delete them. This guide walks through how to set up and run the API and how to use it with Postman.

## Features

- Create a watchlist with a unique ID, name, and a list of scripts (e.g., stock symbols).
- Retrieve a specific watchlist by its ID.
- Update a watchlist by its ID.
- Delete a watchlist by its ID.

## Getting Started

### Prerequisites

- Golang installed on your machine.
- Postman or any other API testing tool.
- A tool like `curl` or `httpie` can also be used to test the API.

### Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/SahilWadhwani/Golang-Watchlist-CRUD-API.git
   cd Golang-Watchlist-CRUD-API
Install dependencies: No external dependencies are required beyond the standard Go library.

```markdown
## Run the API:

```bash
go run main.go
```

The API will start running locally at `http://localhost:7777`.

## API Endpoints

### 1. Create a Watchlist (POST)

- **Endpoint:** `/watchlists`
- **Method:** `POST`
- **Request Body:**
  - Content-Type: `application/json`
  - Example:
    ```json
    {
      "id": "list1",
      "watchListName": "My First List",
      "scripts": ["AAPL", "GOOGL"]
    }
    ```
- **Response:**
  - `201 Created` if the watchlist is successfully created.
  - `400 Bad Request` if the request body is invalid.

#### Steps to Create a Watchlist Using Postman:

1. Open Postman and create a new request.
2. Set the method to `POST` and the URL to `http://localhost:7777/watchlists`.
3. Go to the **Body** tab, select **raw**, and choose **JSON**.
4. In the body, provide a watchlist in JSON format, like this:
    ```json
    {
      "id": "list1",
      "watchListName": "My First List",
      "scripts": ["AAPL", "GOOGL"]
    }
    ```
5. Click **Send** to create the watchlist. You should receive a `201 Created` response if the watchlist is successfully created.

---

### 2. Retrieve a Watchlist (GET)

- **Endpoint:** `/watchlists/{id}`
- **Method:** `GET`
- **Response:**
  - `200 OK` with the watchlist details if found.
  - `404 Not Found` if the watchlist with the provided ID does not exist.

#### Steps to Retrieve a Watchlist Using Postman:

1. Set the method to `GET` and the URL to `http://localhost:7777/watchlists/{id}`, where `{id}` is the unique ID of the watchlist you want to retrieve. For example:
    ```bash
    http://localhost:7777/watchlists/list1
    ```
2. Click **Send**. You will receive the watchlist in the response if it exists.

---

### 3. Update a Watchlist (PUT)

- **Endpoint:** `/watchlists/{id}`
- **Method:** `PUT`
- **Request Body:**
  - Content-Type: `application/json`
  - Example:
    ```json
    {
      "watchListName": "Updated List Name",
      "scripts": ["MSFT", "TSLA"]
    }
    ```
- **Response:**
  - `200 OK` if the watchlist is successfully updated.
  - `404 Not Found` if the watchlist with the provided ID does not exist.

#### Steps to Update a Watchlist Using Postman:

1. Set the method to `PUT` and the URL to `http://localhost:7777/watchlists/{id}`, where `{id}` is the unique ID of the watchlist you want to update. For example:
    ```bash
    http://localhost:7777/watchlists/list1
    ```
2. Go to the **Body** tab, select **raw**, and choose **JSON**.
3. In the body, provide the updated watchlist details, like this:
    ```json
    {
      "watchListName": "Updated List Name",
      "scripts": ["MSFT", "TSLA"]
    }
    ```
4. Click **Send**. You should receive a `200 OK` response if the watchlist is successfully updated.

---

### 4. Delete a Watchlist (DELETE)

- **Endpoint:** `/watchlists/{id}`
- **Method:** `DELETE`
- **Response:**
  - `200 OK` if the watchlist is successfully deleted.
  - `404 Not Found` if the watchlist with the provided ID does not exist.

#### Steps to Delete a Watchlist Using Postman:

1. Set the method to `DELETE` and the URL to `http://localhost:7777/watchlists/{id}`, where `{id}` is the unique ID of the watchlist you want to delete. For example:
    ```bash
    http://localhost:7777/watchlists/list1
    ```
2. Click **Send**. You should receive a `200 OK` response if the watchlist is successfully deleted.

---

## Testing with Postman

You can use Postman to test all the above operations. Here's a quick summary of the requests you can make:

- **POST** `/watchlists` to create a watchlist.
- **GET** `/watchlists/{id}` to retrieve a watchlist by its ID.
- **PUT** `/watchlists/{id}` to update a watchlist by its ID.
- **DELETE** `/watchlists/{id}` to delete a watchlist by its ID.
```

