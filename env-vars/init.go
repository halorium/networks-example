package envvars

import (
	"os"

	"github.com/halorium/networks-example/halt"
)

var AffiliateWindowAPIHost string
var AffilinetAPIHost string
var AwsAccessKeyID string
var AwsContainerCredentialsRelativeURI string
var AwsEndpoint string
var AwsRegion string
var AwsSecretAccessKey string
var BookingAPIHost string
var GitBranch string
var GitSha string
var InPgPassword string
var InPgUser string
var LogColorization string
var LogDebugMessages string
var LogSerialization string
var MigrationsPath string
var NetworksHost string
var NetworksURI string
var PGDatabase string
var PGHost string
var PGPassword string
var PGPort string
var PGUser string
var PostgresPassword string
var PostgresUser string
var SFTPHost string
var SFTPKey string
var SFTPPathRakuten string
var SFTPPort string
var SFTPUser string
var StackName string

var Mocked bool
var NotMocked bool

func init() {
	if os.Getenv("IN_SET_DEBUG_ENV_VARS") == "true" {
		setDebugEnvVars()
	}

	envVarNames := []string{
		// "AWS_ACCESS_KEY_ID",
		// "AWS_CONTAINER_CREDENTIALS_RELATIVE_URI",
		// "AWS_DEFAULT_REGION",
		// "AWS_ENDPOINT",
		// "AWS_REGION",
		// "AWS_SECRET_ACCESS_KEY",
		"IN_AFFILIATE_WINDOW_API_HOST",
		"IN_AFFILINET_API_HOST",
		"IN_BOOKING_API_HOST",
		// "IN_CI",
		// "IN_LOCAL",
		"IN_GIT_BRANCH",
		"IN_GIT_SHA",
		// "IN_JWT_SIGNING_KEY",
		"IN_LOG_COLORIZATION",
		"IN_LOG_DEBUG_MESSAGES",
		"IN_LOG_SERIALIZATION",
		// "IN_MIGRATIONS_PATH",
		"IN_NETWORKS_HOST",
		"IN_NETWORKS_URI",
		"IN_PGPASSWORD",
		"IN_PGUSER",
		"IN_SFTP_HOST",
		"IN_SFTP_KEY",
		"IN_SFTP_PATH_RAKUTEN",
		"IN_SFTP_PORT",
		"IN_SFTP_USER",
		"IN_STACK_NAME",
		"PGDATABASE",
		"PGHOST",
		"PGPASSWORD",
		"PGPORT",
		"PGUSER",
		"POSTGRES_PASSWORD",
		"POSTGRES_USER",
	}

	for _, envVarName := range envVarNames {
		if os.Getenv(envVarName) == "" {
			halt.PanicWith(envVarName + " must be set")
		}
	}

	AffiliateWindowAPIHost = os.Getenv("IN_AFFILIATE_WINDOW_API_HOST")
	AffilinetAPIHost = os.Getenv("IN_AFFILINET_API_HOST")
	AwsAccessKeyID = os.Getenv("AWS_ACCESS_KEY_ID")
	AwsContainerCredentialsRelativeURI = os.Getenv("AWS_CONTAINER_CREDENTIALS_RELATIVE_URI")
	AwsEndpoint = os.Getenv("AWS_ENDPOINT")
	AwsRegion = os.Getenv("AWS_REGION")
	AwsSecretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	BookingAPIHost = os.Getenv("IN_BOOKING_API_HOST")
	GitBranch = os.Getenv("IN_GIT_BRANCH")
	GitSha = os.Getenv("IN_GIT_SHA")
	InPgPassword = os.Getenv("IN_PGPASSWORD")
	InPgUser = os.Getenv("IN_PGUSER")
	LogColorization = os.Getenv("IN_LOG_COLORIZATION")
	LogDebugMessages = os.Getenv("IN_LOG_DEBUG_MESSAGES")
	LogSerialization = os.Getenv("IN_LOG_SERIALIZATION")
	MigrationsPath = os.Getenv("IN_MIGRATIONS_PATH")
	NetworksHost = os.Getenv("IN_NETWORKS_HOST")
	NetworksURI = os.Getenv("IN_NETWORKS_URI")
	PGDatabase = os.Getenv("PGDATABASE")
	PGHost = os.Getenv("PGHOST")
	PGPassword = os.Getenv("PGPASSWORD")
	PGPort = os.Getenv("PGPORT")
	PGUser = os.Getenv("PGUSER")
	PostgresPassword = os.Getenv("POSTGRES_PASSWORD")
	PostgresUser = os.Getenv("POSTGRES_USER")
	SFTPHost = os.Getenv("IN_SFTP_HOST")
	SFTPKey = os.Getenv("IN_SFTP_KEY")
	SFTPPathRakuten = os.Getenv("IN_SFTP_PATH_RAKUTEN")
	SFTPPort = os.Getenv("IN_SFTP_PORT")
	SFTPUser = os.Getenv("IN_SFTP_USER")
	StackName = os.Getenv("IN_STACK_NAME")

	if os.Getenv("IN_LOCAL") == "true" || os.Getenv("IN_CI") == "true" {
		Mocked = true
	} else {
		NotMocked = true
	}
}
