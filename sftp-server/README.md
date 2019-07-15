# Production sftp-server SSH private key update procedure...

Create file ```sftp-server/production.key``` with a valid private key.

Install the travis gem:
```
gem install travis
```

Run the following script, which may require you to authenticate to GitHub for Travis access.
```
sftp-server/encrypt-production-key
```

Add the encrypted private key to git.
```
git add sftp-server/production.key.enc
git commit -m 'update production SFTP server private key'
```
