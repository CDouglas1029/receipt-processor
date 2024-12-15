
# Receipt Processor

A Go-based service for processing and analyzing receipts. This project provides APIs for uploading, parsing, and managing receipt data.

## Table of Contents
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Development](#development)

---

## Features
- Upload receipt data in JSON format
- Parse and analyze receipt information
- Lightweight, scalable, and built with Go
- RESTful API using Gorilla Mux

---

## Prerequisites
Ensure you have the following installed on your system:
- [Go](https://golang.org/) (version 1.21 or higher)
- A terminal or command-line tool

---

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/CDouglas1029/receipt-processor.git
   cd receipt-processor
   ```

2. Initialize and download dependencies:
   ```bash
   go mod tidy
   ```

3. Build the project:
   ```bash
   go build
   ```

---

## Usage

### Running the Application
Start the server by running:
```bash
go run main.go
```

The server will start on `http://localhost:8080` by default.

### Testing the Application
You can use tools like [Postman](https://www.postman.com/) or `curl` to interact with the API. See the [API Endpoints](#api-endpoints) section for details.

---

## API Endpoints

### 1. **Upload Receipt**
- **POST** `/receipts/process`
- **Description**: Upload a receipt in JSON format.
- **Request Body**:
    ```json
    {
    "retailer": "Walgreens",
    "purchaseDate": "2022-01-02",
    "purchaseTime": "08:13",
    "total": "2.65",
    "items": [
        {"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
        {"shortDescription": "Dasani", "price": "1.40"}
    ]
    }
    ```
- **Response**:
    ```json
    {
        "id": "12345"
    }
    ```

### 2. **Get Receipt Details**
- **GET** `/receipts/{id}/points`
- **Description**: Retrieve receipt details by ID.
- **Response**:
    ```json
    {
    "points": 21
    }
    ```

---

## Development

### Code Structure
```
receipt-processor/
├── main.go         # Entry point of the application
├── handlers/       # API route handlers
├── models/         # Data models for receipts
├── utils/          # Helper functions
├── go.mod          # Go module file
└── go.sum          # Dependency file
```

---

## Contact
For any questions or issues, feel free to contact:
- **Name**: Christopher Lim
- **Email**: cdouglaslim@gmail.com
- **GitHub**: https://github.com/CDouglas1029
- **GitHub**: https://www.linkedin.com/in/christopher-lim-544a77334/
