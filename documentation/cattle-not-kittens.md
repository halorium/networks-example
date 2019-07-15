## Cattle, Not Kittens

Cattle, not kittens is a philosophy that servers should be be short-lived and have no personal connnection, no unique non-automated characteristics and/or configuration.

Which is to say they're to be raised and terminated at will and not treated as long-lived pets.

With this system being container based, that removes us one level from the viewpoint of servers, but at the time of writing this we still manage servers to run the EC2 images. AWS Fargate promises the end of all server configuration in the future but is a brand new product and would require some changes to the logging configuration that are not worth pursuing until Fargate has matured a bit, IMHO.

This system maintains it's server via auto-scaling groups (which are not presently set to auto-scale) via Cloudformation. It's trivial to change both the server size and AMI via a tiny commit and merge to master.

All transition is entirely managed by a combination of scripts, Cloudformation, auto-scale groups, ECS and ALB/ELBv2 for zero downtime deployments.
