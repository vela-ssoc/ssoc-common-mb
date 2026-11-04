package mgodsl

import (
	"reflect"
	"strings"
	"time"

	"github.com/vela-ssoc/ssoc-common-mb/datalayer/repository"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type KeyInfo struct {
	Key         string `json:"key"         jsonschema:"字段名"`
	Type        string `json:"type"        jsonschema:"字段类型（string/number/boolean/datetime/oid）"`
	Description string `json:"description" jsonschema:"字段说明"`
}

type KeyInfos []KeyInfo

func (ks KeyInfos) Keys(filter func(KeyInfo) bool) []string {
	keys := make([]string, 0, len(ks))
	for _, k := range ks {
		if filter(k) {
			keys = append(keys, k.Key)
		}
	}

	return keys
}

type CollectionInfo struct {
	Raw         repository.CollectionInfer `json:"-"           jsonschema:"-"`
	Name        string                     `json:"name"        jsonschema:"集合名称"`
	Description string                     `json:"description" jsonschema:"集合备注说明"`
	Keys        KeyInfos                   `json:"keys"        jsonschema:"集合内的字段列表"`
}

type ParserOptions struct {
}

func ForCollection(zero repository.CollectionInfer) CollectionInfo {
	name, comment := zero.CollectionInfo()

	rtyp := reflect.TypeOf(zero)
	keys := parseStruct(rtyp, "")

	return CollectionInfo{
		Raw:         zero,
		Name:        name,
		Description: comment,
		Keys:        keys,
	}
}

func parseStructField(sf reflect.StructField, prefix string) KeyInfos {
	if !sf.IsExported() {
		return nil
	}

	stag := sf.Tag.Get("jsonschema")
	if stag == "-" {
		return nil
	}
	btag := sf.Tag.Get("bson")
	if btag == "-" {
		return nil
	}
	name, _, _ := strings.Cut(btag, ",")
	if name == "" {
		name = strings.ToLower(sf.Name)
	}
	if prefix != "" {
		name = prefix + "." + name
	}

	key := KeyInfo{Key: name, Description: stag}

	sft := sf.Type
	sfkind := sf.Type.Kind()
	if sfkind == reflect.Array && reflect.TypeFor[bson.ObjectID]() == sft {
		key.Type = "oid"
		return KeyInfos{key}
	}
	if sfkind == reflect.Slice || sfkind == reflect.Array {
		sft = sft.Elem()
	}
	if sft.Kind() == reflect.Pointer {
		sft = sft.Elem()
	}

	switch sft.Kind() {
	case reflect.String:
		key.Type = "string"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		key.Type = "number"
	case reflect.Bool:
		key.Type = "boolean"
	case reflect.Struct:
		if reflect.TypeFor[time.Time]() == sft {
			key.Type = "datetime"
		}
	default:
	}

	if key.Type != "" {
		return KeyInfos{key}
	}

	return parseStruct(sft, name)
}

func parseStruct(t reflect.Type, prefix string) KeyInfos {
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil
	}

	keys := make(KeyInfos, 0, 10)
	for field := range t.Fields() {
		subs := parseStructField(field, prefix)
		keys = append(keys, subs...)
	}

	return keys
}
