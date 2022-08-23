
## Common issues
  - gitignore: missing DS_Store and others -> https://www.toptal.com/developers/gitignore/
  - missing documentation and instructions
  - missing proper error handling
  - missing tests
  - duplicated functionality (db, config, entity, mq, envString, close queues, etc.)
  - unhandled db versioning/migration
  - missing infrastructure builders (docker, ci, etc.)
  - tightly coupled layers (application layer with db and mq, main.go with mq)
  - hard to test / missing architectural patterns
  - linting issues
  - centralized login format and/or package

## Application service
  - applicationCreatedQueue value mismatch (main.go 15 needs to be the same (value or variable) as application/application.go 15)
  - empty files (Dockerfile, application/package.go)

## Application job
  - missing gracefully shutdown
  - why partners if the domain is application
  - global scoped function call -> envString()
