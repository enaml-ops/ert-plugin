package config

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

func GetFingerPrint(pubKey string) (string, error) {
	parts := strings.Fields(pubKey)
	if len(parts) < 2 {
		return "", fmt.Errorf("bad key")
	}

	k, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return "", errors.Wrap(err, "DecodeString failed")
	}

	fp := md5.Sum([]byte(k))
	var res string

	for i, b := range fp {
		res += fmt.Sprintf("%02x", b)

		if i < len(fp)-1 {
			res += fmt.Sprint(":")
		}
	}
	return res, nil
}
