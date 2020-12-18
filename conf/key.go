package conf

import (
	"bytes"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"encoding/hex"

	"github.com/bi-zone/ruwireguard-go/crypto/gost/gost3410"
)

const PrivateKeyLength = 32
const PublicKeyLength = 33

type PrivateKey [PrivateKeyLength]byte
type PublicKey [PublicKeyLength]byte

var curve = gost3410.CurveIdtc26gost34102012256paramSetA()

// PrivateKey methods

func (k *PrivateKey) String() string {
	return base64.StdEncoding.EncodeToString(k[:])
}

func (k *PrivateKey) HexString() string {
	return hex.EncodeToString(k[:])
}

func (k *PrivateKey) IsZero() bool {
	var zeros PrivateKey
	return subtle.ConstantTimeCompare(zeros[:], k[:]) == 1
}

func (k *PrivateKey) Public() *PublicKey {
	var p PublicKey
	var zeros PrivateKey

	if bytes.Equal(k[:], zeros[:]) {
		return &p
	}

	privateKey, err := gost3410.NewPrivateKey(curve, k[:])
	if err != nil {
		return nil
	}

	pubKey, err := privateKey.PublicKey()
	if err != nil {
		return nil
	}

	copy(p[:], gost3410.MarshalCompressed(curve, pubKey.X, pubKey.Y))

	return &p
}

func NewPresharedKey() (*PrivateKey, error) {
	var k PrivateKey

	_, err := rand.Read(k[:])
	if err != nil {
		return nil, err
	}

	return &k, nil
}

func NewPrivateKey() (*PrivateKey, error) {
	key, err := gost3410.GenPrivateKey(curve, rand.Reader)
	if err != nil {
		return nil, err
	}

	var k PrivateKey
	copy(k[:], key.Raw())

	return &k, nil
}

func NewPrivateKeyFromString(b64 string) (*PrivateKey, error) {
	return parsePrivateKeyBase64(b64)
}

// PublicKey methods

func (k *PublicKey) String() string {
	return base64.StdEncoding.EncodeToString(k[:])
}

func (k *PublicKey) HexString() string {
	return hex.EncodeToString(k[:])
}

func (k *PublicKey) IsZero() bool {
	var zeros PublicKey
	return subtle.ConstantTimeCompare(zeros[:], k[:]) == 1
}

func NewPublicKeyFromString(b64 string) (*PublicKey, error) {
	return parsePublicKeyBase64(b64)
}
