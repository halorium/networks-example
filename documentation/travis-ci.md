## Travis CI

Travis CI handles CI and CD duties via automation scripts.

The ./crank process is designed to be extraordinarily similar on Travis CI and
Mac OS X High Sierra developer machines.

All tooling differences between the two platforms have shims in the
compatibility/ directory.

Travis runs are on a deprecated (group: deprecated-2017Q4) image that was
required to get TLS operational. It should be trivial in the future to fix
this once Certbot and/or Travis fixes issues that prevented us from using
completely current images.

Travis contains several security-critical encrypted environment variables via
their standard mechanisms to do so. This makes it unnecessary for developers
to have write-access to AWS resources during development.

    AWS_SECRET_ACCESS_KEY
    POSTGRES_PASSWORD
    IN_PGPASSWORD
    PGPASSWORD
    IN_LOGGLY_TOKEN
    IN_JWT_SIGNING_KEY
    IN_SYSDIG_TOKEN

Of these, all are self-explanatory except IN_PGPASSWORD and PGPASSWORD.

IN_PGPASSWORD is the password of the networksexample user

PGPASSWORD is the password of the postgres root user and is used only during
new cluster creation.

The rest of the Travis CI process is managed by 5 automation scripts

• travis-ci/install
• tls/manage-files
• ./crank
• travis-ci/deploy
• travis-ci/uninstall
