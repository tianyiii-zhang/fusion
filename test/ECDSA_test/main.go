// Use ECDSA algorithm to verify the signature. Users should use
// their own public key to sign and the public key is included in
// the DID document.

package main

import (
	"bytes"
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// Request is the type of this request. Vault_ID and Entry_ID index the data which user wants to get.
// Signature is the result of signing the above three variables.
type Proof struct {
	Request   string
	Vault_ID  string
	Entry_ID  string
	Signature string
}

func main() {
	proof := Proof{
		Request:  "request",
		Vault_ID: "1",
		Entry_ID: "1",
	}
	str := proof.Request + proof.Vault_ID + proof.Entry_ID
	msg := []byte(str)

	privateKey, err := crypto.HexToECDSA("aad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	// Convert public key to byte format
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	pubk := hexutil.Encode(publicKeyBytes)
	// Print hex of public key
	fmt.Println("The hexadecimal of the public key is:\n", pubk)

	hash := crypto.Keccak256Hash(msg)
	// Print hex of message hash
	fmt.Println("The hexadecimal of the message hash is:\n", hash.Hex()) // 0x1c8aff950685c2ed4bc3174f3472287b56d9517b9c948127319a09a7a36deac8

	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	proof.Signature = hexutil.Encode(signature)

	// Print the hexadecimal of the signature
	fmt.Println("The hexadecimal of the signature is:\n", hexutil.Encode(signature)) // 0x789a80053e4927d0a898db8e065e948f5cf086e32f9ccaa54c1908e22ac430c62621578113ddbb62d509bf6049b8fb544ab06d36f916685a2eb8e57ffadde02301

	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), signature)
	if err != nil {
		log.Fatal(err)
	}

	matches := bytes.Equal(sigPublicKey, publicKeyBytes)
	if !matches {
		log.Fatal("The public-key recover from signature has error")
	}

	signatureNoRecoverID := signature[:len(signature)-1] // remove recovery id
	verified := crypto.VerifySignature(publicKeyBytes, hash.Bytes(), signatureNoRecoverID)
	fmt.Println("The signature verification result is:", verified) // true
}
