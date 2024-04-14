### Directory Descriptions

- **`/cmd/crawler`**: Contains the main application executable for crawler. The app crawls the internet from unmatched query and the query results like merchant name are push to MQ.
- **`/cmd/embedding_generator`**: Contains the main application executable for generator. This is where we receive msg from MQ and create embeddings for it and store in vector DB.
- **`/cmd/merchant_search`**: Contains the main application executable for mechant search api. The app exposes an api to get normalised merchant name. If not found pushes to unmatched pattern storage for crawling.
- **`/pkg`**: Houses all the primary logic of the application, organized into different packages:
  - **`/crawler`**: Core logic for crawling operations.
  - **`/api`**: Manages interactions with third-party APIs.
  - **`/queue`**: Handles interactions with RabbitMQ for message queuing.
  - **`/db`**: Contains all database interaction code, including connection setup and query execution.
  - **`/controllers`**: Contains request handlers for query search frontend and healthcheck.
  - **`/embeddings`**: Contains code to generate embeddings from text and push it to vector store.
- **`/config`**: Manages all configurations across the application, centralizing settings into one location.
- **`/test`**: Includes all unit and integration tests for the various components of the application.
- **`/deploy`**: Contains scripts and configuration files for deploying the application, including Docker configurations.

### Getting Started
