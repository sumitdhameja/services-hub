# SERVICES-HUB

![](https://storage.googleapis.com/bc-ops/github-gif/service-hub.png)

TODO:

- Check SQL schema generated and tune if needed

**Services Hub API** is a API build using gin/gorm. The API can be run as a command

## Structure and Assumptions

### Assumptions

1. A SQL DB is used to create relations between user/service/versions. This can be further extended to orgs where a user can belong to an org.
   For the sake of simplicity a user can create a service and a given service has various versions associated to it.
   The design only permits a user to query service created by that user only.
   AuthN/AuthZ can be added as middleware if needed.

   There are 2 API's created

   - /api/v1/users/${user_id}/services\
     This returns all services for a given user
   - /api/v1/users/${user_id}/services/${service_id}\
     This returns all service versions for a given service, provided that service belongs to that user

2. Pagination is applied to the first API so client need not query all services at once. The API accepts a list of query params to paginate the request.
   List of query params:
   - limit: integer. Limits the results returned
   - page: integer. Page of the results
   - search: string. search string to search title or description
     OffSet is calculated internally.

### Structure

```bash
|____cmd
| |____other-cmd
| |____api
| | |____server.go
| | |____migrate.go
| | |____main.go
| | |____root.go
|____go.mod
|____bin
|____config
| |____config.go
| |____config.dev.yaml
|____Makefile
|____internal
| |____dto
| | |____response.go
| | |____request.go
| |____middleware
| | |____middleware.go
| | |____pagination.go
| |____logger
| | |____logger.go
| | |____level.go
| |____utils
| | |____pagination.go
| |____models
| | |____users.go
| | |____models.go
| | |____services.go
| | |____migrations.go
| | |____service_versions.go
| |____api
| | |____v1
| | | |____app_service.go
| | | |____handlers.go
| |____errors
| | |____middleware_gin.go
| | |____errors.go
| |____services
| | |____user.go
| | |____app_service.go
| | |____app_service_version.go
|____go.sum
|____README.md
|____.gitignore
|____.git
```

## Build

Use Makefile to build the project. The binary generated would be added to /bin folder

```bash
$ VERSION=1.0.0 ENV=dev make build
```

## Getting Started

1. Commands Available

   - root command: This will provide help with all the available commands
   - **server**: server will run the API on port 8080 by default
   - **migrate**: migrate will migrate the schema needed and populate database with fake data for users/services/versions

2. See all available commands

   ```bash
   $ make root
   ```

   OR

   ```bash
   $ go run ./cmd/api
   ```

3. Run migrate command to create schema and populate data

   ```bash
   $ make migrate
   ```

   OR

   ```bash
   $ go run ./cmd/api migrate
   ```

4. Start by running the API. The configuration needed to connect to mysql is under config/config.dev.yaml

   ```bash
   $ make run
   ```

   OR

   ```bash
   $ go run ./cmd/api server
   ```
