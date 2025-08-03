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

## Use Cases & API Reference

This API provides Thai address data (provinces, districts, subdistricts) with pagination, search, and sorting support. See the **[Full API Documentation](docs/api_reference.md)** for more details.

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
