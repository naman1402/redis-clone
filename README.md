Here’s the revised **README.md** file, following your instructions:

---

# Redis Clone

## Overview

Redis Clone is a lightweight, in-memory key-value store developed in **Go**, inspired by the core features of Redis. This project aims to provide a foundational understanding of key-value store architecture, concurrent client handling, and efficient data storage mechanisms.  

The project currently supports a set of essential commands for interacting with strings, hashes, and lists, making it ideal for learning or experimenting with Redis-like functionality in Go.

---

## Features Checklist

### Implemented Commands
- **SET** `key value`
  Store a key-value pair in the memory store.
- **GET** `key`  
  Retrieve the value associated with a key.
- **HSET** `key field value`  
  Store a field-value pair in a hash.
- **HGET** `key field`  
  Retrieve the value of a specific field in a hash.
- **HGETALL** `key`  
  Retrieve all field-value pairs of a hash.

### Unimplemented Commands
- **EXPIRE** `key seconds`  
  Set a timeout on a key.
- **DEL** `key`  
  Delete a key from the memory store.
- **EXISTS** `key`  
  Check if a key exists in the memory store.

---

## How to Run the Project

Follow these steps to run and interact with the Redis Clone project:

### 1. Clone the Repository
First, clone this repository to your local system and navigate to the project directory:
```bash
git clone https://github.com/naman1402/redis-clone.git
cd redis-clone
```

### 2. Build and Start the Server
Use Docker to build and run the project:
```bash
docker build -t redis-clone .
docker run -p 6379:6379 redis-clone
```

### 3. Install and Use `redis-cli`
To interact with the Redis Clone server, you need `redis-cli`. Follow these steps:

#### Install Redis on Windows (via Ubuntu WSL)
Follow the official guide to set up Redis on Windows using Ubuntu WSL:  
[**Install Redis on Windows (via Ubuntu WSL)**](https://redis.io/docs/latest/operate/oss_and_stack/install/install-redis/install-redis-on-windows/)

#### Install Redis on Ubuntu (inside WSL)
Run the following commands in your WSL terminal to install Redis:
```bash
sudo apt update
sudo apt install redis-server
```

#### Start Redis Server
Even though you're using the Redis Clone, you need to start the `redis-server` service for the CLI to function:
```bash
sudo service redis-server start
```

#### Connect with `redis-cli`
Once the server is running, connect to the Redis Clone with:
```bash
redis-cli -h localhost -p 6379
```

### 4. Manage Docker Containers
To manage Docker containers, use these commands:

- **List running containers:**
  ```bash
  docker ps
  ```
- **Stop a running container:**
  ```bash
  docker stop <container_id>
  ```
- **Remove a stopped container:**
  ```bash
  docker rm <container_id>
  ```
- **Delete the Docker image:**
  ```bash
  docker rmi redis-clone
  ```


## Acknowledgements

This project was inspired by Redis, one of the most popular in-memory key-value databases. Redis Clone is built as a learning tool to deepen understanding of how Redis functions under the hood, focusing on architecture, concurrency, and protocol handling in Go.  

Feel free to explore, experiment, and contribute! 🚀

