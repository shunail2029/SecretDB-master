package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/crypto"
	"gopkg.in/yaml.v2"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
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
func makeSignature(keybase keys.Keybase, name, passphrase string, msg []byte) (crypto.PubKey, []byte, error) {
	if keybase == nil {
		var err error
		keybase, err = keys.NewKeyring(sdk.KeyringServiceName(), viper.GetString(flags.FlagKeyringBackend), viper.GetString(flags.FlagHome), os.Stdin)
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
