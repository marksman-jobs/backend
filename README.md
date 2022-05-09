# Marksman.com - Backend

## Setup instructions

1. Ensure you have the following installed on your system and added to your PATH. Mac installation instructions are given below
    a. `go 1.17`
      ```bash
      go install golang.org/dl/go1.17.2@latest
      go1.17.2 download
      ```
    b. `sqlc`
      ```bash
      brew install sqlc
      ```
    c. `Golang migrate`
      ```bash
      brew install golang-migrate
      ```
    d. `docker` [Installation instructions](https://docs.docker.com/desktop/mac/install/)
    e. `GNU make`
      ```bash
      xcode-select --install
      ```

2. After verifying that all the above are installed, navigate to the root of the project in your terminal and type
    ```bash
    source .env
    make setup
    make postgresup
    make createdb
    make migrateup
    go1.17.2 run main.go
    ```

3. To shutdown the processes, exit the go process and type
    ```bash
    make migratedown
    make dropdb
    make postgresdown
    ```
