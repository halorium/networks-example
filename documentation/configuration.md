## Configuration

Configuration is done via environment variables.
These environment variables are located in the .env file or overridden in the docker-compose.yaml file for local development.
For deployment to AWS the environment variables are located in aws/stack-templates/environment.gotemplate file as well as the .travis.yml file which contains encrypted environment variables such as AWS creds, DB creds, etc.

Many of the services in this repo use the unified env-vars package to access various environment variables.

The network and network-account entities also contain some configuration as part of their dataset.
