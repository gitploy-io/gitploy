package license

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"strings"

	"github.com/gitploy-io/gitploy/model/extent"
)

const (
	publicPem = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAspbW4NEaD9w6PW/hfvoa
cZ8yR+hJe9iGQ6T3NTc7dUORL3BpJL+Cr+kO7TyL8l9bu4i66AlD7SqoJ0TShZqU
oieonoSozoI61+oKNYUJxs3I755Ubp5ZNYWd0NR7miFBH8FBUr2on0S1PE6CKJx+
6Gydtx/301xe29x9/lKHIm3CB5twsaXi4HJiL5p8EAd93szXP2pkbBjanE+ZSkfS
+ZqGE9WHBnG+6BdSytgjBJypCyX8VXoxnfnZotZTnd4dHWTdMkax4dLNecw7xDMU
Xl+6e0hW2ZWSf1abjkzDjM3loSgB2rmRbKTIbAovBCG7nwuMmlBk+Mjqftr0Sjvc
4wIDAQAB
-----END PUBLIC KEY-----`
)

// Decode verifies the key and
// returns the signing data if it is verified.
func Decode(key string) (*extent.SigningData, error) {
	_, encodedSigning, encodedSignature, err := parseKey(key)
	if err != nil {
		return nil, err
	}

	ok, err := verifySignature(encodedSigning, encodedSignature)
	if !ok {
		return nil, err
	}

	return decodeSigningData(encodedSigning)
}

func verifySignature(encodedSigning, encodedSignature string) (bool, error) {
	block, _ := pem.Decode([]byte(publicPem))
	if block == nil {
		return false, fmt.Errorf("failed to parse PEM block containing the public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return false, fmt.Errorf("failed to parse DER encoded public key: %s", err)
	}

	hashed := sha256.Sum256([]byte(fmt.Sprintf("key/%s", encodedSigning)))

	// Decode signature.
	// The signature is encoded by base64
	decodedSignature, err := base64.URLEncoding.DecodeString(encodedSignature)
	if err != nil {
		return false, fmt.Errorf("failed to decode the signature: %s", err)
	}

	err = rsa.VerifyPKCS1v15(pub.(*rsa.PublicKey), crypto.SHA256, hashed[:], decodedSignature)
	if err != nil {
		return false, fmt.Errorf("failed to verify the signature: %s", err)
	}

	return true, nil
}

// parseKey returns the signining prefix, the encoded key, and the encoded signiture.
// the key and the signiture are encoded by base64.
func parseKey(key string) (string, string, string, error) {
	signingSignature := strings.Split(key, ".")
	if len(signingSignature) != 2 {
		return "", "", "", fmt.Errorf("couldn't split into the signing and the signature")
	}

	prefixKey := strings.Split(signingSignature[0], "/")
	if len(prefixKey) != 2 {
		return "", "", "", fmt.Errorf("couldn't split into the prefix and the key")
	}

	return prefixKey[0], prefixKey[1], signingSignature[1], nil
}

func decodeSigningData(encodedSigning string) (*extent.SigningData, error) {
	decoded, err := base64.URLEncoding.DecodeString(encodedSigning)
	if err != nil {
		return nil, fmt.Errorf("failed to decode the signing data: %s", err)
	}

	signing := &extent.SigningData{}
	if err := json.Unmarshal(decoded, signing); err != nil {
		return nil, fmt.Errorf("failed to unmarshal the signing data: %s", err)
	}

	return signing, nil
}
