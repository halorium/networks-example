## TLS

TLS certificates are provided by LetsEncrypt via CertBot.

The certificates are cached on S3 for 75 days of their 90 day life to avoid
creating more certificates for a domain than are allowed by LetsEncrypt.

Certificates are also stored in the ACM service for use by ALB/ELBv2.

Certificates are also stored on the API containers so that TLS is used from
client to application. It's not quite end-to-end as the ALB/ELBv2 decrypts
and re-encrypts but is highly preferable to the common practice of terminating
TLS at the ALB/ELBv2 and using non-TLS connection to the application.

### Implementation Details
The application flow for TLS is as follows:
.travis.yml initiates the travis-ci/install script
This script installs certbot.
Then tls/manage-files script is called. This is where all the work is done to either get current certs or new ones.
The cert files are stored on S3 in (s3://halorium.networks-example/tls/files)
If they are current and valid then they are copied to the services and used otherwise we fetch new certs from LetsEncrypt and put them in S3 and then into ACM and on the containers.
We use wildcard certs to limit the number of certs we need to generate as well as the number we upload into the ACM as this is limited.
