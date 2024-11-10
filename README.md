# walletun

## Flow

1. Registration and Authentication:
   Client -> User Service (REST)

2. Transaction Initiation (with JWT token):
   Client -> API Gateway (with JWT) -> Transaction Service -> Message Broker

3. Transaction Processing:
   Message Broker -> Transaction Service (Worker) -> ScyllaDB
                                                  -> Wallet Service (gRPC) -> PostgreSQL -> Transaction Service (Worker) -> ScyllaDB
                                                  


## Diagram
```
Authentication Flow:
--------------------

+------------+
|   Client   |
+------------+
       |
       v
+------------------+
|  User Service    |
|     (REST)       |
| (Register/Login) |
+------------------+
       |
       v
     (JWT)
       |
       v

Main Flow:
----------

+------------+          +--------------+          +----------------+          +----------+
|   Client   | -------> | API Gateway  | -------> | TransactionSvc |          | ScyllaDB |
+------------+          +--------------+          +----------------+          +----------+
                                                          |                         ^
                                                          v                         |
                                                  +-----------------+               |
                                                  |    Message      |               |
                                                  |     Broker      |               |
                                                  +-----------------+               |
                                                          |                         |
                                                          v                         |
                                                  +-----------------+               |
                                                  | TransactionSvc  | --------------/
                                                  |    (Worker)     |
                                                  +-----------------+
                                                          |
                                                          v
                                                  +-----------------+          +------------+
                                                  | Wallet Service  | -------> | PostgreSQL |
                                                  |     (gRPC)      |          +------------+
                                                  +-----------------+
                                                          |
                                                          v
                                                  +-----------------+
                                                  | Message Broker  |
                                                  |(Event Completed)|
                                                  +-----------------+
                                                          |
                                                          v
                                                  TransactionSvc updates
                                                       ScyllaDB
                                                       
```