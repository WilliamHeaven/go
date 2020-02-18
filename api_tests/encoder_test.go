package test

import (
	"bytes"
	"encoding/json"
	"github.com/json-iterator/go"
	"github.com/stretchr/testify/require"
	"testing"
)

// Standard Encoder has trailing newline.
func TestEncoderHasTrailingNewline(t *testing.T) {
	should := require.New(t)
	var buf, stdbuf bytes.Buffer
	enc := jsoniter.ConfigCompatibleWithStandardLibrary.NewEncoder(&buf)
	enc.Encode(1)
	stdenc := json.NewEncoder(&stdbuf)
	stdenc.Encode(1)
	should.Equal(stdbuf.Bytes(), buf.Bytes())
}

func Test_set_indent(t *testing.T) {
    should := require.New(t)
    var buf1, buf2 bytes.Buffer
    enc1 := jsoniter.NewEncoder(&buf1)
    enc1.SetIndent("", " ")
    enc1.Encode([]int{1, 2})
    enc2 := json.NewEncoder(&buf2)
    enc2.SetIndent("", " ")
    enc2.Encode([]int{1, 2})
    should.Equal(buf1.String(), buf2.String())
}
