package awshttpsigningclient

import "fmt"

func authorizationHeader(state *State) {
	hmacSHA256Binary := hmacSHA256Binary(state.SigningKey, state.StringToSign)
	hmacSHA256Hex := fmt.Sprintf("%x", hmacSHA256Binary)

	authorization := ALGORITHM + " " +
		"Credential=" + state.Credential + ", " +
		"SignedHeaders=" + state.SignedHeaders + ", " +
		"Signature=" + hmacSHA256Hex

	state.Request.Header.Set(
		"Authorization",
		authorization,
	)
}
