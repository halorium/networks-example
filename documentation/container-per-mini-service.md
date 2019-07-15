## Container per mini-service

Mini-services are executable binaries wrapped in docker containers that provide stand-alone functionality for an entire top-level class of entities.

In practice, that can scale down to a single HTTP endpoint, and up to multiple endpoints for the first path segment of the entities URI, including multiple verbs, and subordinate entities. Such mini-services are always identified by the term "api".

Asynchronous message processing workers are and should always be handled by a separate worker mini-service in order to avoid destabilizing the relatively simple HTTP endpoints should problems arise in the relatively more complex backends. Such mini-services are always identified by the term "worker".

Mini-services that do not serve API endpoints and do not process queued messages, such as cycler, are always identified by the term "service".

Docker is used extensively in this system.

Miniservices are built into Docker containers and managed locally by docker-compose.

A complete "snowglobe" environment is built, including local mocks for datastores (PostgreSQL), AWS services (SNS and SQS), an SFTP server, and 3rd party vendor APIs.

This allows 100% control over the testing environment.

Most the the docker images are built from SCRATCH and contain only a few files each. This keeps them small, lightweight and trivially managable.
