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

### Getting Started

Before moving forward please make sure you have docker. Follow the steps - 
1) Run `docker pull mysql:latest`
2) Run `docker pull qdrant/qdrant:latest`
3) Run `docker pull rabbitmq:latest`
4) Run `docker volume create db_data`
5) Run `docker run --rm -d --name fold-mysql -v db_data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=user123 -p 3306:3306 mysql`. Use the `tables.sql` file to create the respective DB and table.
6) Run `docker run -p 6333:6333 -p 6334:6334 -e QDRANT__SERVICE__GRPC_PORT="6334" --rm -d qdrant/qdrant`. Create collection with 
```
curl --location --request PUT 'localhost:6333 collections/merchants' --header 'Content-Type: application/json' --data '{
  "vectors": {
    "size": 64,
    "distance": "Dot"
  }
}
' 
```. 
The collection name is merchants and vector size we are using is 64. 
7) Run `docker run -d --hostname my-rabbit --name fold-rabbitmq -p 5672:5672 -e RABBITMQ_DEFAULT_USER=user -e RABBITMQ_DEFAULT_PASS=user123 --rm rabbitmq`
8) Clone the repo and move into the directory.
9) Run `docker build -t merchant_search:0.1 -f Dockerfile_merchant_search  .`
10) Run `docker build -t generator:0.1 -f Dockerfile_embedding_generator  .`
11) Run `docker build -t crawler:0.1 -f Dockerfile_crawler  .`
12) Run `docker container run --rm -p 4000:4000  --network="host" -d crawler:0.1`. This will start the container at `localhost:4000`.
Please hit `localhost:4000/healthcheck` to verify the service is running.
13) Run `docker container run --rm -p 5000:5000  --network="host" -d generator:0.1`. This start the container at `localhost:5000`.
Please hit `localhost:5000/healthcheck` to verify the service is running.
14) Run `docker container run --rm -p 6000:6000  --network="host" -d merchant_search:0.1`. This will start the container at `localhost:6000`. Please hit `localhost:6000/healthcheck` to verify the service is running.
15) Use the curl `curl --location 'localhost:6000/search?q=ajio' `to send request to merchant service. If the data is not there for query, it will insert the data into the SQL table. The crawler will pick the query and crawl google to get results. Right now the crawler is scheduled for every minute. After a couple of minutes hit the request again and hopefuly you will see the desired results. 