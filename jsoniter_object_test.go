package jsoniter

import (
	"encoding/json"
	"fmt"
	"testing"
	"github.com/json-iterator/go/require"
	"bytes"
)

func Test_empty_object(t *testing.T) {
	iter := ParseString(`{}`)
	field := iter.ReadObject()
	if field != "" {
		t.Fatal(field)
	}
}

func Test_one_field(t *testing.T) {
	iter := ParseString(`{"a": "b"}`)
	field := iter.ReadObject()
	if field != "a" {
		fmt.Println(iter.Error)
		t.Fatal(field)
	}
	value := iter.ReadString()
	if value != "b" {
		t.Fatal(field)
	}
	field = iter.ReadObject()
	if field != "" {
		t.Fatal(field)
	}
}

func Test_two_field(t *testing.T) {
	iter := ParseString(`{ "a": "b" , "c": "d" }`)
	field := iter.ReadObject()
	if field != "a" {
		t.Fatal(field)
	}
	value := iter.ReadString()
	if value != "b" {
		t.Fatal(field)
	}
	field = iter.ReadObject()
	if field != "c" {
		t.Fatal(field)
	}
	value = iter.ReadString()
	if value != "d" {
		t.Fatal(field)
	}
	field = iter.ReadObject()
	if field != "" {
		t.Fatal(field)
	}
	iter = ParseString(`{"field1": "1", "field2": 2}`)
	for field := iter.ReadObject(); field != ""; field = iter.ReadObject() {
		switch field {
		case "field1":
			iter.ReadString()
		case "field2":
			iter.ReadInt64()
		default:
			iter.reportError("bind object", "unexpected field")
		}
	}
}

func Test_write_object(t *testing.T) {
	should := require.New(t)
	buf := &bytes.Buffer{}
	stream := NewStream(buf, 4096)
	stream.IndentionStep = 2
	stream.WriteObjectStart()
	stream.WriteObjectField("hello")
	stream.WriteInt(1)
	stream.WriteMore()
	stream.WriteObjectField("world")
	stream.WriteInt(2)
	stream.WriteObjectEnd()
	stream.Flush()
	should.Nil(stream.Error)
	should.Equal("{\n  \"hello\":1,\n  \"world\":2\n}", buf.String())
}

type TestObj struct {
	Field1 string
	Field2 uint64
}

func Benchmark_jsoniter_object(b *testing.B) {
	for n := 0; n < b.N; n++ {
		iter := ParseString(`{"field1": "1", "field2": 2}`)
		obj := TestObj{}
		for field := iter.ReadObject(); field != ""; field = iter.ReadObject() {
			switch field {
			case "field1":
				obj.Field1 = iter.ReadString()
			case "field2":
				obj.Field2 = iter.ReadUint64()
			default:
				iter.reportError("bind object", "unexpected field")
			}
		}
	}
}

func Benchmark_json_object(b *testing.B) {
	for n := 0; n < b.N; n++ {
		result := TestObj{}
		json.Unmarshal([]byte(`{"field1": "1", "field2": 2}`), &result)
	}
}
