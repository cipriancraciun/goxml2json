package xml2json

import (
	"bytes"
	"io"
)

// Convert converts the given XML document to JSON
func Convert(r io.Reader, ps ...encoderPlugin) (*bytes.Buffer, error) {
	// Decode XML document
	root := &Node{}
	err := NewDecoder(r).Decode(root)
	if err != nil {
		return nil, err
	}

	// Then encode it in JSON
	buf := new(bytes.Buffer)
	e := NewEncoder(buf)
	for _, p := range ps {
		e = p.AddTo(e)
	}
	err = e.Encode(root)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
