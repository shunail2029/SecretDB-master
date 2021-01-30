package cli

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/url"
	"os"

	gethSecp256k1 "github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/shunail2029/SecretDB-master/x/secretdb/types"
	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	"golang.org/x/crypto/sha3"
	"gopkg.in/yaml.v2"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/codec"
	cryptokeys "github.com/cosmos/cosmos-sdk/crypto/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func findItemBeginnings(data []byte) []int {
	var idxs []int
	cnt := 0
	for i, d := range data {
		if cnt == 0 {
			idxs = append(idxs, i)
		}
		if d == '{' {
			cnt++
		} else if d == '}' {
			cnt--
		}
	}
	return idxs
}

func printOutput(toPrint interface{}, outputFormat string, indent bool) error {
	var (
		out []byte
		err error
	)

	switch outputFormat {
	case "text":
		out, err = yaml.Marshal(&toPrint)

	case "json":
		if indent {
			out, err = json.MarshalIndent(toPrint, "", "  ")
		} else {
			out, err = json.Marshal(toPrint)
		}
	}

	if err != nil {
		return err
	}

	fmt.Println(string(out))
	return nil
}

// makeSignature builds a signature of (msg byte[])
func makeSignature(keybase cryptokeys.Keybase, name, passphrase string, msg []byte) (crypto.PubKey, []byte, error) {
	if keybase == nil {
		var err error
		keybase, err = cryptokeys.NewKeyring(sdk.KeyringServiceName(), viper.GetString(flags.FlagKeyringBackend), viper.GetString(flags.FlagHome), os.Stdin)
		if err != nil {
			return nil, nil, err
		}
	}

	sigBytes, pubkey, err := keybase.Sign(name, passphrase, msg)
	if err != nil {
		return nil, nil, err
	}
	return pubkey, sigBytes, nil
}

// GenerateSharedKey generates shared key with given pukey and privkey in specified keyring
func GenerateSharedKey(pubkey secp256k1.PubKeySecp256k1, kb cryptokeys.Keybase, name, pass string, cdc *codec.Codec) (types.SharedKey, error) {
	var err error
	if kb == nil {
		kb, err = cryptokeys.NewKeyring(sdk.KeyringServiceName(), types.KeyringBackend, types.CLIHome, os.Stdin)
		if err != nil {
			return types.SharedKey{}, err
		}
	}
	pk, err := kb.ExportPrivateKeyObject(name, pass)
	if err != nil {
		return types.SharedKey{}, err
	}
	var privkey secp256k1.PrivKeySecp256k1
	err = cdc.UnmarshalBinaryBare(pk.Bytes(), &privkey)
	if err != nil {
		return types.SharedKey{}, err
	}

	x, y := gethSecp256k1.DecompressPubkey(pubkey[:])
	if x == nil {
		return types.SharedKey{}, errors.New("failed to decompress pubkey")
	}
	key, _ := gethSecp256k1.S256().ScalarMult(x, y, privkey[:])
	if key == nil {
		return types.SharedKey{}, errors.New("failed to multiply pubkey by privkey")
	}
	return sha3.Sum256(key.Bytes()), nil
}

// EncryptWithKey encrypts data with given shared key
func EncryptWithKey(plainText []byte, key types.SharedKey) (cipherText []byte, err error) {
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return
	}

	cipherText = make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		return
	}

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)
	return
}

// DecryptWithKey decrypts data with given shared key
func DecryptWithKey(cipherText []byte, key types.SharedKey) (plainText []byte, err error) {
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return
	}

	plainText = make([]byte, len(cipherText)-aes.BlockSize)
	iv := cipherText[:aes.BlockSize]
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(plainText, cipherText[aes.BlockSize:])
	return
}

func encryptMsg(plainMsg []byte, cliCtx context.CLIContext, kb cryptokeys.Keybase, cdc *codec.Codec) (cipherMsg []byte, err error) {
	// get operator's pubkey
	res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", types.ModuleName, types.QueryGetOperatorPubkey), nil)
	if err != nil {
		return
	}
	var pubkeyStr string
	cdc.MustUnmarshalJSON(res, &pubkeyStr)
	pubkeyStr, err = url.PathUnescape(pubkeyStr)
	if err != nil {
		return
	}
	var pubkey secp256k1.PubKeySecp256k1
	err = cdc.UnmarshalBinaryBare([]byte(pubkeyStr), &pubkey)
	if err != nil {
		return
	}

	key, err := GenerateSharedKey(pubkey, kb, cliCtx.GetFromName(), keys.DefaultKeyPass, cdc)
	if err != nil {
		return
	}
	cipherMsg, err = EncryptWithKey(plainMsg, key)
	return
}
