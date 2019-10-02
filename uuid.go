package goIdentity

import (
	"crypto/ecdsa"
	"crypto/sha1"
	"encoding/base64"
	"crypto/x509"
	"encoding/pem"
)

func CalculateUUID(privateKey *ecdsa.PrivateKey) string {
	/*
	 *  The fingerprint is what the TS3 Client shows you as "Unique ID".
	 *  fingerprint := base64encode(sha1(PUBLICKEY))
	 */
	pub, _ := x509.MarshalPKIXPublicKey(privateKey.Public())
	pubBytes := pem.EncodeToMemory(&pem.Block{
		Type: "ECDSA PUBLIC KEY",
		Bytes: pub,
	})

	h := sha1.New()
	h.Write([]byte(string(pubBytes)))
	str := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return  str;
}
