package awshttpsigningclient

// Credentials stores the information necessary to authorize with AWS and it
// is from this information that requests are signed.
type Credentials struct {
	AccessKeyID     string `json:"AccessKeyId"`
	Code            string `json:"Code"`
	Expiration      string `json:"Expiration"`
	RoleArn         string `json:"RoleArn"`
	SecretAccessKey string `json:"SecretAccessKey"`
	Token           string `json:"Token"`
	Type            string `json:"Type"`
}
