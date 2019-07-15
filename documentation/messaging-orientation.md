## Messaging Orientation

Messaging based architecture is where the application is built to distribute and process messages when actions take place within the system.

For instance, given the request:

* PUT /networks/1

There are several messages that could be emitted and worked by that operation.

* emit: put-network-request
* work: put-network-request
* emit: put-network-in-postgres
* work: put-network-in-postgres
* emit: put-network-in-postgres-success (or -failure)
* emit: put-network-in-elasticsearch
* work: put-network-in-elasticsearch
* emit: put-network-in-elasticsearch-success (or -failure)
* emit: put-network-request-success (or -failure)

Notice that not all messages have workers. All messages are logged for the purposes of debugging and telemetry, so the entire point of some messages is to get logged for the purposes of telemetry.

Messages in this system have strict JSON representations to avoid code complexity in the workers.

Messages to workers are sent through a pub-sub system. This makes it trivial to create multiple independent workers to process the "same" message such as storing an entity in both postgres and elasticsearch. This allows new functionality to be added to the system with no need to alter any existing code -- which is one of the most significant advantages of this architecture.

This system relies upon AWS SNS for pub/sub and AWS SQS for message queuing.

### Messaging in use
The following endpoints sends an SNS messages which in turn creates an SQS messages:

POST /networks/:id/accounts/:id/purchases/load
This sends the message "<branch-name>-post-network-account-purchases-load" with state.
A worker is fetching messages from the queue "<branch-name>-post-network-account-purchases-load" and runs the code to fetch transactions from the network.

PUT /networks/:id/accounts/:id/purchases/:id
This sends the message "<branch-name>-put-network-account-purchase" with state.
A worker is fetching messages from the queue "<branch-name>-put-purchase-in-postgres" and runs the code to store the entity in the Postgres DB.
