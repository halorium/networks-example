## HTTP Endpoints

### AUTHENTICATION
  * POST /authentication
  * POST /authentication/validate (temporary until authentication is in use)

### IDENTITIES
  * POST /identities
  * GET  /identities/:identity-id
  * GET  /identities/:identity-id/history
  * PUT  /identities/:identity-id/role-set

### NETWORKS
  * GET    /networks
  * DELETE /networks/:network-id
  * GET    /networks/:network-id
  * PUT    /networks/:network-id
  * GET    /networks/:network-id/history

### NETWORK ACCOUNTS
  * GET    /networks/:network-id/accounts
  * DELETE /networks/:network-id/accounts/:network-account-id
  * GET    /networks/:network-id/accounts/:network-account-id
  * PUT    /networks/:network-id/accounts/:network-account-id
  * GET    /networks/:network-id/accounts/:network-account-id/history

### NETWORK ACCOUNT PURCHASES
  * GET /networks/:network-id/accounts/:network-account-id/purchases
  * GET /networks/:network-id/accounts/:network-account-id/purchases/:purchase-id
  * PUT /networks/:network-id/accounts/:network-account-id/purchases/:purchase-id
  * GET /networks/:network-id/accounts/:network-account-id/purchases/:purchase-id/history
  * GET /networks/:network-id/accounts/:network-account-id/purchase-updates

### NETWORK ACCOUNT PURCHASES LOAD
  * POST /networks/:network-id/accounts/:network-account-id/purchases/load

### ROLES
  * GET /roles/:role-id
  * PUT /roles/:role-id
  * GET /roles/:role-id/history

### STATUS
  * GET /status
