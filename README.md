# Weather Information System

This Code uses **[OpenWeather API](api.openweathermap.org)** to fetch and the store weather data to a MySQL database which can be accessed by an REST API Endpoint.

### REST APIs

|No. | API Endpoint | HTTP Method | Description |
| --- | --- | --- | --- |
|1.| **`/weather`** |**POST**| Fetch Data from OpenWeather API |
|2.| **`/weather/`*`{cityName}`*** |**GET**| Get Data from the MySQL DB |

### Steps to Run

Paste the following Commands on your terminal:

#### Step 1: Create a Database

```bash
mysql -u root -p
```

```bash
CREATE DATABASE new_database;
```

#### Step 2: Import the  Database Schema

```bash
mysql -u username -p new_database < data-dump.sql
```

#### Step 3: Import the Go Modules/Dependencies

```bash
go get -v ./...
```

#### Step 4: Run the Code

```bash
go run main.go
```

#### Step 5: Access the Endpoints

You can use **[Postman](https://www.postman.com/)** to access the REST API Endpoints mentioned above.

(The service runs on "http://localhost:9001")

### Folder Structure:

```bash
.
├── cmd
│   ├── database
│   │   └── connection.go
│   ├── models
│   │   └── models.go
│   └── pkg
│       ├── endpoints
│       │   └── endpoints.go
│       ├── handlers
│       │   └── handlers.go
│       ├── repositories
│       │   └── repositories.go
│       └── service
│           ├── cities.go
│           └── service.go
├── go.mod
├── go.sum
├── LICENSE
├── main.go
├── README.md
└── weather-information-system.sql

8 directories, 13 files
```
