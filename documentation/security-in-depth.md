## Security In Depth

This system was built to be extremely secure. The security is layered in a
practice known as security-in-depth.

  • Docker containers built from SCRATCH have a single executable with no
    OS tooling including no shells. Each runs a single process.

  • Docker containers have minimal port assignments.

  • Docker containers are executed in production with Docker supplied SELinux
    profiles.

  • TLS encryption extends from initiating client to AWS ELBv2 to individual
    Docker containers.

  • AWS VPC security groups strictly control inbound IP traffic. Just ports 80
    and 443 are open to external traffic, and only to the AWS ELBv2.

  • AWS ECS containers hosts are in a separate security group that only allows
	  inbound connectivity from the AWS ELBv2 on unprivileged (high) port numbers.

  • AWS ECS tasks use strictly limited IAM task profiles for accessing AWS
    services and have no AWS credentials on them.

  • Developers have zero need for more than read-only access to AWS services
    to build, test, update, and monitor the system. All deployment related
    activity is automated via the CI/CD system (Travis CI) via scripts included
    in this repository.
