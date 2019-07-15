package awsclient

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/halorium/networks-example/aws/http-signing-client"
	"github.com/halorium/networks-example/env-vars"
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/halt"
	"github.com/halorium/networks-example/http-client"
	"github.com/halorium/networks-example/logger"
)

type Client struct {
	Creds      *awshttpsigningclient.Credentials
	HTTPClient *httpclient.Client
	SNS        *sns.SNS
	SQS        *sqs.SQS
}

func new() *Client {
	logger.Debug("awsclient-new", nil)

	var client *Client

	if envvars.Mocked {
		client = &Client{
			HTTPClient: httpclient.New(),
			SNS:        localSNSService(),
			SQS:        localSQSService(),
		}
	} else {
		httpClient, creds := awshttpsigningclient.New()

		client = &Client{
			Creds:      creds,
			HTTPClient: httpClient,
			SNS:        awsSNSService(),
			SQS:        awsSQSService(),
		}
	}

	return client
}

func localSNSService() *sns.SNS {
	return sns.New(localSession())
}

func awsSNSService() *sns.SNS {
	return sns.New(awsSession())
}

func localSQSService() *sqs.SQS {
	return sqs.New(localSession())
}

func awsSQSService() *sqs.SQS {
	return sqs.New(awsSession())
}

func localSession() *session.Session {
	creds := credentials.NewEnvCredentials()

	configs := &aws.Config{
		Credentials: creds,
		Region:      aws.String(envvars.AwsRegion),
		DisableSSL:  aws.Bool(envvars.Mocked),
		Endpoint:    aws.String(envvars.AwsEndpoint),
	}

	sess, err := session.NewSession(configs)

	if err != nil {
		logger.Panic("awsclient-local-session", flaw.From(err))
	}

	return sess
}

func awsSession() *session.Session {
	sess, err := session.NewSession()

	if err != nil {
		logger.Panic("awsclient-aws-session", flaw.From(err))
	}

	_, err = sess.Config.Credentials.Get()

	if err != nil {
		halt.Panic(flaw.From(err).Wrap("cannot awsSession"))
	}

	return sess
}
