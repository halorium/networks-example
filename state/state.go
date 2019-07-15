package state

import (
	"io"
	"regexp"

	"github.com/halorium/networks-example/authentication/authorization"
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/identities/identity"
	"github.com/halorium/networks-example/identities/identity-role-set"
	"github.com/halorium/networks-example/json-web-token"
	"github.com/halorium/networks-example/networks/network"
	"github.com/halorium/networks-example/networks/network-account"
	"github.com/halorium/networks-example/networks/network-account-purchase"
	"github.com/halorium/networks-example/networks/network-account-purchase-load"
	"github.com/halorium/networks-example/notifier"
	"github.com/halorium/networks-example/request"
	"github.com/halorium/networks-example/responses"
	"github.com/halorium/networks-example/roles/role"
	"github.com/halorium/networks-example/stamps"
)

type State struct {
	// Context specific objects
	ID        string            `json:"id"`
	CreatedAt *stamps.Timestamp `json:"created-at"`
	Nickname  string            `json:"nickname"`
	FlawError *flaw.Error       `json:"flaw-error"`

	Authorization *authorization.Entity `json:"authorization,omitempty"`

	Identity   *identity.Entity   `json:"identity,omitempty"`
	Identities []*identity.Entity `json:"identities,omitempty"`

	IdentityRoleSet *identityroleset.Entity `json:"identity-role-set,omitempty"`

	JSONWebToken *jsonwebtoken.Token `json:"json-web-token,omitempty"`

	Network  *network.Entity   `json:"network,omitempty"`
	Networks []*network.Entity `json:"networks,omitempty"`

	NetworkAccount  *networkaccount.Entity   `json:"network-account,omitempty"`
	NetworkAccounts []*networkaccount.Entity `json:"network-accounts,omitempty"`

	NetworkAccountPurchase  *networkaccountpurchase.Entity   `json:"network-account-purchase,omitempty"`
	NetworkAccountPurchases []*networkaccountpurchase.Entity `json:"network-account-purchases,omitempty"`

	NetworkAccountPurchasesLoad *networkaccountpurchaseload.Entity `json:"network-account-purchases-load,omitempty"`

	Role  *role.Entity   `json:"role,omitempty"`
	Roles []*role.Entity `json:"roles,omitempty"`

	Message              *notifier.Message `json:"-"`
	Regexp               *regexp.Regexp    `json:"-"`
	VariantReadCloser    io.ReadCloser     `json:"-"`
	VariantReadCloserTwo io.ReadCloser     `json:"-"`
	VariantResponse      interface{}       `json:"-"`

	// HTTP
	Request  *request.Request    `json:"-"`
	Response *responses.Response `json:"-"`
}
