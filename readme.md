# Kafka with Terraform and Golang App

This project demonstrates how to configure Kafka and Zookeeper using Terraform and Docker, then use a simple Golang application to consume messages from Kafka.

## Prerequisites

- Terraform installed on your machine
- Docker installed and running
- Go (Golang) installed

## Project Structure

- `main.tf`: Terraform file that defines Docker containers for Kafka and Zookeeper.
- `main.go`: Main Go program that connects to Kafka.
- `kafka/startkafka.go`: Contains Kafka consumer logic that listens for messages on a specific topic.

## Getting Started

### Step 1: Initialize Terraform

Initialize your Terraform environment. This will download necessary providers and set up the environment.

```bash
terraform init
```

### Step 2: Apply Terraform Configuration

This will create a custom Docker network and start Kafka and Zookeeper containers.

```bash
terraform apply
```

- You will be prompted to confirm the execution plan. Type `yes` to proceed.

### Step 3: Run the Golang Application

Once Kafka and Zookeeper are running, you can run the Go application. The app will connect to Kafka and listen for messages on `topic1`.

```bash
go run main.go
```

The Go app will:
- Connect to the Kafka broker at `localhost:9092`.
- Subscribe to `topic1`.
- Listen for and print messages.

### Step 4: Send a Message to Kafka

To send a message to Kafka, follow these steps:

1. **Exec into the Kafka container**:

```bash
docker exec -it kafka /bin/bash
```

2. **Create a topic** (`topic1` if it doesn't already exist):

```bash
kafka-topics --create --topic topic1 --bootstrap-server localhost:9092 --partitions 1 --replication-factor 1
```

3. **Send a message to `topic1`**:

```bash
kafka-console-producer --topic topic1 --bootstrap-server localhost:9092
```

- After running the command, type a message and hit `Enter`.

### Step 5: Read Messages from Kafka

You can view the messages sent to the Kafka topic using the Go application or directly from the Kafka container.

1. **Exec into the Kafka container** if you haven’t already:

```bash
docker exec -it kafka /bin/bash
```

2. **Consume messages from `topic1`**:

```bash
kafka-console-consumer --topic topic1 --bootstrap-server localhost:9092 --from-beginning
```

This will display the messages that were sent to `topic1`.

### How it Works

1. **Terraform** sets up Kafka and Zookeeper in Docker containers.
2. **Zookeeper** manages Kafka's cluster state and broker coordination.
3. The **Golang app** connects to Kafka and consumes messages from `topic1`.
4. You can send and read messages using Kafka’s producer and consumer CLI commands inside the Docker container.

## Useful Commands

- **Terraform commands**:
  - `terraform init`: Initializes Terraform and downloads the required providers.
  - `terraform apply`: Provisions the Docker containers for Kafka and Zookeeper.
  - `terraform destroy`: Removes the containers and network created by Terraform.

- **Docker commands**:
  - `docker exec -it kafka /bin/bash`: Opens a shell inside the Kafka container.
  - `kafka-console-producer`: Sends messages to a Kafka topic.
  - `kafka-console-consumer`: Reads messages from a Kafka topic.

- **Go command**:
  - `go run main.go`: Runs the Golang app that connects to Kafka and listens for messages.

### Clean Up

To remove the infrastructure created by Terraform, run:

```bash
terraform destroy
```

This will stop and remove the Docker containers for Kafka and Zookeeper.

---

Enjoy your Kafka setup and happy coding!"
