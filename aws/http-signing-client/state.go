package awshttpsigningclient

import (
	"net/http"
	"regexp"
	"time"

	"github.com/halorium/networks-example/env-vars"
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/logger"
)

type State struct {
	Creds                      *Credentials  `json:"-"`
	FlawError                  *flaw.Error   `json:"flaw-error"`
	CanonicalHeaders           string        `json:"canonical-headers"`
	CanonicalQuery             string        `json:"canonical-query"`
	CanonicalRequest           string        `json:"canonical-request"`
	CanonicalRequestSHA256Hex  string        `json:"canonical-request-sha256-hex"`
	CanonicalURI               string        `json:"canonical-uri"`
	Credential                 string        `json:"credential"`
	Region                     string        `json:"region"`
	Request                    http.Request  `json:"-"`
	RequestBodySHA256Hex       string        `json:"request-body-sha256-hex"`
	RequestDate                string        `json:"request-date"`
	RequestTime                time.Time     `json:"request-time"`
	Scope                      string        `json:"scope"`
	Service                    string        `json:"service"`
	Signature                  string        `json:"signature"`
	SignedHeaders              string        `json:"signed-headers"`
	SigningKey                 []byte        `json:"signing-key"`
	SortedLowercaseHeaderNames []string      `json:"sorted-lowercase-header-names"`
	SpaceRunRegexp             regexp.Regexp `json:"-"`
	StringToSign               string        `json:"string-to-sign"`
	XAmzDate                   string        `json:"x-amz-date"`
	RefreshAvailableAt         time.Time     `json:"refresh-available-at"`
	CredsFromSTS               bool          `json:"creds-from-sts"`
}

func NewState() *State {
	logger.Debug("awshttpsigningclient-new-state", nil)

	return &State{
		Creds: &Credentials{
			AccessKeyID:     envvars.AwsAccessKeyID,
			SecretAccessKey: envvars.AwsSecretAccessKey,
		},
		RefreshAvailableAt: time.Now(),
		SpaceRunRegexp:     *regexp.MustCompile(" +"),
	}
}
