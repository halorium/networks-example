## INSTALLATION
### Install Brew (first time only)
```
/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
```

### Install Go (1.9.2), jq and OpenSSL
```
brew update
brew install go jq openssl
```

### Set GOPATH and GOBIN
```
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$PATH
```

### Install Docker
* Download and install [Docker for Mac](https://www.docker.com/products/docker#/mac)

### Clone this repo into $GOPATH
```
mkdir -p $GOPATH/src/github.com/halorium
cd $GOPATH/src/github.com/halorium
git clone https://github.com/halorium/networks-example
```
### Validate, Build and Test
```
./crank
```
### If all is green you're good-to-go

## DOCUMENTATION

[Overview](documentation/README.md)

## OPERATION

The AWS ECS AMI-ID needs to be refreshed periodically.

The current version details are available at:

    https://aws.amazon.com/marketplace/fulfillment?productId=52d5fd7f-92c7-4d60-a830-41a596f4d8f3&region=us-east-1 (ami-13401669, 2017.09.d)
