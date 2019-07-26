// +build go1.11,!go1.12

package stdlib

// Code generated by 'goexports crypto'. DO NOT EDIT.

import (
	"crypto"
	"io"
	"reflect"
)

func init() {
	Symbols["crypto"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"BLAKE2b_256":  reflect.ValueOf(crypto.BLAKE2b_256),
		"BLAKE2b_384":  reflect.ValueOf(crypto.BLAKE2b_384),
		"BLAKE2b_512":  reflect.ValueOf(crypto.BLAKE2b_512),
		"BLAKE2s_256":  reflect.ValueOf(crypto.BLAKE2s_256),
		"MD4":          reflect.ValueOf(crypto.MD4),
		"MD5":          reflect.ValueOf(crypto.MD5),
		"MD5SHA1":      reflect.ValueOf(crypto.MD5SHA1),
		"RIPEMD160":    reflect.ValueOf(crypto.RIPEMD160),
		"RegisterHash": reflect.ValueOf(crypto.RegisterHash),
		"SHA1":         reflect.ValueOf(crypto.SHA1),
		"SHA224":       reflect.ValueOf(crypto.SHA224),
		"SHA256":       reflect.ValueOf(crypto.SHA256),
		"SHA384":       reflect.ValueOf(crypto.SHA384),
		"SHA3_224":     reflect.ValueOf(crypto.SHA3_224),
		"SHA3_256":     reflect.ValueOf(crypto.SHA3_256),
		"SHA3_384":     reflect.ValueOf(crypto.SHA3_384),
		"SHA3_512":     reflect.ValueOf(crypto.SHA3_512),
		"SHA512":       reflect.ValueOf(crypto.SHA512),
		"SHA512_224":   reflect.ValueOf(crypto.SHA512_224),
		"SHA512_256":   reflect.ValueOf(crypto.SHA512_256),

		// type definitions
		"Decrypter":     reflect.ValueOf((*crypto.Decrypter)(nil)),
		"DecrypterOpts": reflect.ValueOf((*crypto.DecrypterOpts)(nil)),
		"Hash":          reflect.ValueOf((*crypto.Hash)(nil)),
		"PrivateKey":    reflect.ValueOf((*crypto.PrivateKey)(nil)),
		"PublicKey":     reflect.ValueOf((*crypto.PublicKey)(nil)),
		"Signer":        reflect.ValueOf((*crypto.Signer)(nil)),
		"SignerOpts":    reflect.ValueOf((*crypto.SignerOpts)(nil)),

		// interface wrapper definitions
		"_Decrypter":     reflect.ValueOf((*_crypto_Decrypter)(nil)),
		"_DecrypterOpts": reflect.ValueOf((*_crypto_DecrypterOpts)(nil)),
		"_PrivateKey":    reflect.ValueOf((*_crypto_PrivateKey)(nil)),
		"_PublicKey":     reflect.ValueOf((*_crypto_PublicKey)(nil)),
		"_Signer":        reflect.ValueOf((*_crypto_Signer)(nil)),
		"_SignerOpts":    reflect.ValueOf((*_crypto_SignerOpts)(nil)),
	}
}

// _crypto_Decrypter is an interface wrapper for Decrypter type
type _crypto_Decrypter struct {
	WDecrypt func(rand io.Reader, msg []byte, opts crypto.DecrypterOpts) (plaintext []byte, err error)
	WPublic  func() crypto.PublicKey
}

func (W _crypto_Decrypter) Decrypt(rand io.Reader, msg []byte, opts crypto.DecrypterOpts) (plaintext []byte, err error) {
	return W.WDecrypt(rand, msg, opts)
}
func (W _crypto_Decrypter) Public() crypto.PublicKey { return W.WPublic() }

// _crypto_DecrypterOpts is an interface wrapper for DecrypterOpts type
type _crypto_DecrypterOpts struct {
}

// _crypto_PrivateKey is an interface wrapper for PrivateKey type
type _crypto_PrivateKey struct {
}

// _crypto_PublicKey is an interface wrapper for PublicKey type
type _crypto_PublicKey struct {
}

// _crypto_Signer is an interface wrapper for Signer type
type _crypto_Signer struct {
	WPublic func() crypto.PublicKey
	WSign   func(rand io.Reader, digest []byte, opts crypto.SignerOpts) (signature []byte, err error)
}

func (W _crypto_Signer) Public() crypto.PublicKey { return W.WPublic() }
func (W _crypto_Signer) Sign(rand io.Reader, digest []byte, opts crypto.SignerOpts) (signature []byte, err error) {
	return W.WSign(rand, digest, opts)
}

// _crypto_SignerOpts is an interface wrapper for SignerOpts type
type _crypto_SignerOpts struct {
	WHashFunc func() crypto.Hash
}

func (W _crypto_SignerOpts) HashFunc() crypto.Hash { return W.WHashFunc() }
