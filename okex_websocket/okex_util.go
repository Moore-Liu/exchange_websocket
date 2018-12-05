package okex_websocket

import (
	"bytes"
	"compress/flate"
	"io/ioutil"
)

func GzipDecode(in []byte) ([]byte, error) {
	reader := flate.NewReader(bytes.NewReader(in))
	defer reader.Close()

	return ioutil.ReadAll(reader)

}
