# Thai Address API

This is a simple REST API for retrieving Thai address information (provinces, districts, subdistricts, and postal codes). It is built using the Go Echo framework and provides endpoints to fetch address data in JSON format. The database is stored in SQLite.

## Getting Started

### Prerequisites

Make sure you have the following installed:

1. [Go](https://go.dev/) (version 1.24 or later)
2. [golangci-lint](https://golangci-lint.run/)
3. [Docker](https://www.docker.com/) (optional, for deployment)
4. [SQLite3](https://sqlite.org/) (optional, for local database management)

Run `go mod tidy` to install the required dependencies.

### Make Commands

1. You can start the development server using:

   ```bash
   make run
   ```

   By default, the server will run on `http://localhost:8080`.

2. You can run the linters using:

   ```bash
   make lint
   ```

## Use Case

### Common Parameters

For all endpoints, you can use the following common parameters:

- `page`: The page number for pagination.
- `page_size`: The number of items per page (default is all items).
- `search`: The keyword to search for in the Thai or English names of the specific address components.
- `sort_order`: The order of sorting; default is ascending (0). Use 1 for descending order.

**Note:** All endpoints support `sort_by`, but the available options depend on the specific endpoint.

### Response Format

All responses are returned in JSON format with the following fields:

```json
{
  "meta": {
    // Record counts
    "total_records": 100,
    "display_records": 10,
    // Pagination details
    "current_page": 1,
    "records_per_page": 10,
    "total_pages": 10
  },
  "data": {},
  "message": "Normal or error message"
}
```

Details of the fields:

- `data`: Varies depending on the endpoint.
- `meta`: Contains pagination details. Included only if both the page and page_size parameters are provided.
- **Error responses:** Do not include meta or data. They contain only a message field describing the error.

### API Endpoints

All endpoints under the prefix `/api/v1` return data in JSON format and use GET requests. Below is a list of available endpoints:

| Endpoint                     | Description                                 | Extra Query Parameters   | Remarks                                                                                                       |
| ---------------------------- | ------------------------------------------- | ------------------------ | ------------------------------------------------------------------------------------------------------------- |
| `/provinces`                 | Retrieve a list of provinces in Thailand    | `sort_by`                | `sort_by` must be 0 (province_id), 1 (name_th), or 2 (name_en)                                                |
| `/districts`                 | Retrieve a list of districts in Thailand    | `province_id`, `sort_by` | `province_id` is optional; `sort_by` must be 0 (district_id), 1 (name_th), or 2 (name_en)                     |
| `/subdistricts`              | Retrieve a list of subdistricts in Thailand | `district_id`, `sort_by` | `district_id` is optional; `sort_by` must be 0 (subdistrict_id), 1 (name_th), 2 (name_en), or 3 (postal_code) |
| `/subdistricts/:postal_code` | Retrieve subdistricts by postal code        | `sort_by`                | `sort_by` must be 0 (subdistrict_id), 1 (name_th), 2 (name_en), or 3 (postal_code)                            |

**Note:** The API provides a health check endpoint at the base route (`/`) that returns "Service is running" if everything is functioning properly.

## Deployment

This API can be deployed using Docker on Google Cloud Run. You can build the Docker image and run it as follows:

```bash
export GCP_PROJECT_ID=<your-gcp-project-id>
make deploy
```

**Note:** The above command assumes you have set up Google Cloud SDK and authenticated your account. It will automatically build the Docker image and push it to Google Artifact Registry before deploying it to Cloud Run.

## Contributing

Feel free to contribute to this project by submitting issues or pull requests. Please ensure that your code adheres to the existing style and includes appropriate tests. The existing CI pipeline will run automatically to check for linting before manually triggering the deployment.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
