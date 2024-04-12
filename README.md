Project layout

/crawler-project
|-- /cmd
|   |-- /crawler
|       |-- main.go           # Entry point for the crawler
|
|-- /pkg
|   |-- /crawler             # Core crawling logic
|   |   |-- crawler.go       # Main crawler functionalities
|   |   |-- setup.go         # Setup configurations for the crawler
|   |
|   |-- /api                 # Third-party API interactions
|   |   |-- client.go        # API client setup and methods
|   |   |-- types.go         # Data types for API responses
|   |
|   |-- /queue               # RabbitMQ integration
|   |   |-- producer.go      # Message producing functionalities
|   |   |-- consumer.go      # Message consuming functionalities
|   |
|   |-- /db                  # Database interactions
|   |   |-- conn.go          # Database setup and queries
|   |   |-- model.go         # Database models
|   |   |-- config.go        # Database connection
|   |   |-- query.go         # GetUnmatched pattern
|   |
|   |-- /config              # Configuration management
|       |-- config.go        # Load and handle configuration settings
|
|-- /internal
|   |-- /util                # Utility functions and helpers
|       |-- helper.go        # Helper functions
|
|-- /test                    # Test files
|   |-- crawler_test.go      # Tests for crawler functionalities
|   |-- api_test.go          # Tests for API client
|   |-- db_test.go           # Tests for database interactions
|
|-- /deploy                  # Deployment scripts and configuration files
|   |-- docker-compose.yml   # Docker-compose for local deployment/testing
|
|-- .gitignore               # Specifies intentionally untracked files to ignore
|-- README.md                # Project overview and setup instructions
|-- go.mod                   # Go module definitions
|-- go.sum                   # Go module checksums
