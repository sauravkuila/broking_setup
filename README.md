# broking_setup

This is a yield engine repository for investments made by Aspire on behalf of clients.

## Getting Started

To get started with the boilerplate, follow these steps:

1. Clone this repository: `git clone git@github.com:sauravkuila/broking_setup.git`
2. Navigate to the project directory: `cd broking_setup`
3. Install dependencies: `go mod tidy`
4. Set up your environment variables by creating a `local.yaml` file based on the provided `local.yaml.example`.
5. Run the application: `make run`

### Build Code
Build the code to check if the server instance can be created. This is useful to do before a Docker build for quick fixes
```
make build
```

### Run Code
```
make run
```