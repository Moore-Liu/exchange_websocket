package huobi_websocket

import (
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"fmt"
	"io/ioutil"
)

func GzipDecompress(in []byte) ([]byte, error) {
	b := new(bytes.Buffer)
	binary.Write(b, binary.LittleEndian, in)
	r, err := gzip.NewReader(b)
	if err != nil {
		fmt.Println("[ParseGzip] NewReader error: ", err)
		return nil, err
	} else {
		defer r.Close()
		undatas, err := ioutil.ReadAll(r)
		if err != nil {
			fmt.Println("[ParseGzip] ioutil.ReadAll error: ", err)
			return nil, err
		}
		return undatas, nil
	}

}
