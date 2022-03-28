# Gofood

Gofood mono repository provides some microservices and challenges

**NOTE:** This is a test project and will not be developed

## Development:
  ### Commands
    * Install dependencies: `make install`
    * Run a service: `make run [service]`
    * Format a service codes: `make fmt [service]`
    * Up infrastructure: `docker-compose -f docker-compose.infra.yml up -d`
    * Build a service image: `make build [service] [version]`