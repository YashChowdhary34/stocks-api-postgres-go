# Stocks API with PostgreSQL and Go

This repository contains a Go application that interfaces with a PostgreSQL database to manage stock data. The application provides RESTful API endpoints to perform Create, Read, Update, and Delete (CRUD) operations on stock entries.

## Features

- **Create Stock:** Add new stock entries to the database.
- **Retrieve Stocks:** Fetch details of existing stocks.
- **Update Stock:** Modify information of existing stocks.
- **Delete Stock:** Remove stock entries from the database.

## Project Structure

The project is organized into the following directories:

- `models/`: Contains the data models representing the stock entities.
- `router/`: Defines the routes/endpoints for the API.
- `middleware/`: Includes middleware functions for tasks such as logging and authentication.

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/YashChowdhary34/stocks-api-postgres-go.git
cd stocks-api-postgres-go
```

### 2. Install Dependencies

Ensure you have Go installed. Then, install the required Go modules:

```bash
go mod tidy
```

### 3. Set Up PostgreSQL Database

- Install PostgreSQL and create a new database named `stocks_db`.
- Create a table for storing stock information:

```sql
CREATE TABLE stocks (
    stockid SERIAL PRIMARY KEY,
    name VARCHAR(100),
    price NUMERIC,
    company VARCHAR(100)
);
```

### 4. Configure Environment Variables

Create a `.env` file in the project root with the following content:

```
POSTGRES_URL=your_postgres_url
```

Replace `your_postgres_url` with your actual PostgreSQL connection string.

### 5. Run the Application

Start the Go application:

```bash
go run main.go
```

## API Endpoints

- **Create Stock:** `POST /api/new-stock`
- **Get All Stocks:** `GET /api/stock`
- **Get Stock by ID:** `GET /api/stock/{id}`
- **Update Stock by ID:** `PUT /api/stock/{id}`
- **Delete Stock by ID:** `DELETE /api/delete-stock/{id}`

## Contributing

Contributions are welcome. Feel free to submit a pull request or open an issue to contribute to the project. Please include a detailed description of the changes and any relevant context.

## Contact

For any questions or inquiries, please contact the maintainer at [YashChowdhary34](https://github.com/YashChowdhary34).
