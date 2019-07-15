### TLS Workflow
* .travis.yml
  * install:
    * travis-ci/install
      * use virtualenv (enables multiple side-by-side installations of python)
      * this allows us to use the version of python that certbot needs
      * install certbot, awscli, etc.
      * aws/create-cli-config
  * before_script:
    * tls/manage-files
      * tls/get-files-from-s3
        * get files created-at date
        * check if files created more than 75 days ago
        * if not then copy files
      * or
      * tls/get-files-from-letsencrypt
        * use certbot client to get certs from let's encrypt
        * put txt dns as validation and then remove after
      * tls/move-and-rename-files
        * put files in tls/files directory with proper names
      * tls/put-files-in-s3
        * sync tls/files to s3
      * tls/put-files-in-acm
        * import the certificate files into acm from tls/files directory
  * script:
    * ./crank && travis-ci/deploy
      * this builds the services which includes copying tls/files onto docker images
      * docker webservers use tls certificates to ListenAndServeTLS
      * delete unused certificates from acm
  * after_script:
    * travis-ci/uninstall
      * remove tls/certbot-files
      * remove tls/files
      * remove /tmp/fresh-certificate-arn
