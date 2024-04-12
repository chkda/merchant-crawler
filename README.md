### Directory Descriptions

- **`/cmd/crawler`**: Contains the main application executable. This is where the application is initialized and run from the command line.
- **`/pkg`**: Houses all the primary logic of the application, organized into different packages:
  - **`/crawler`**: Core logic for crawling operations.
  - **`/api`**: Manages interactions with third-party APIs.
  - **`/queue`**: Handles interactions with RabbitMQ for message queuing.
  - **`/db`**: Contains all database interaction code, including connection setup and query execution.
  - **`/config`**: Manages all configurations across the application, centralizing settings into one location.
- **`/internal/util`**: Provides utility functions that are internal to the application and not meant to be used by external packages.
- **`/test`**: Includes all unit and integration tests for the various components of the application.
- **`/deploy`**: Contains scripts and configuration files for deploying the application, including Docker configurations.

### Getting Started
