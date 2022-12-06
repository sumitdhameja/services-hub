# SERVICES-HUB

![](https://storage.googleapis.com/bc-ops/github-gif/service-hub.png)

**Services Hub API** is a API build using gin/gorm. The API can be run as a command. This can be used as a boilerplate template showcasing API capabilities in golang

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

## Configuration

Configuration options are available under config/config.dev.yaml. Configuration Options can also be set as environment variables.

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

## Build

Use Makefile to build the project. The binary generated would be added to /bin folder

```bash
$ VERSION=1.0.0 ENV=dev make build
```

## Considerations & improvements

I chose to use gorm library(wanted to try it out). The solution seems clean using gorm(preloads dependents in struct), however makes multiple round trips to DB.
If roundtrips to DB is an issue. Performance can be obtained by either of the below:

1. use standard golang SQL library. LEFT JOIN users with services and service_versions table
   - This provides single round trip to DB. However, will produce more rows because of multiple versions associated with each service.
2. Use SQL functions to generate JSON which can be scanned into a struct. Example query is below. SQL will return aggregated JSON as a single field.

   - This would be the most efficient solution(but not so ellegant). This could be difficult to maintain, since any change in structs in golang or SQL may break the code. Also, if you would like to switch to another relational DB, the query may need to be tested.

   ```sql
   SELECT
   JSON_ARRAYAGG(JSON_OBJECT('id', u.id, 'email', u.email, 'services', s.services))
   FROM
      users u

      LEFT JOIN (
         SELECT
               user_id,
               id,
               JSON_ARRAYAGG(JSON_OBJECT('id', id, 'title', title, 'description', description, 'service_versions', sv.json_versions)) services
         FROM services s
               LEFT JOIN (
                  SELECT
                     service_id,
                     JSON_ARRAYAGG(JSON_OBJECT('id', id, 'version', version)) json_versions
                  FROM service_versions
                  GROUP BY service_id
               ) sv ON sv.service_id = s.id
         GROUP BY user_id,id
      ) s ON s.user_id = u.id
      WHERE u.id = '0be2de59-56e0-47a7-aabc-a82ae064c85f'
   ```

3. Index is created on title field for fast lookups using FULLTEXT search
4. Other design considerations for improvements
   - Cache data in memory(redis/memcache) for a given user's services for faster lookups and apply pagination. Clear cache on POST/UPDATE for a given record
