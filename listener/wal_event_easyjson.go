// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package listener

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson39bd694bDecodeGithubComIhippikWalListenerListener(in *jlexer.Lexer, out *WalEvent) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "nextlsn":
			out.NextLSN = string(in.String())
		case "change":
			if in.IsNull() {
				in.Skip()
				out.Change = nil
			} else {
				in.Delim('[')
				if out.Change == nil {
					if !in.IsDelim(']') {
						out.Change = make([]ChangeItem, 0, 1)
					} else {
						out.Change = []ChangeItem{}
					}
				} else {
					out.Change = (out.Change)[:0]
				}
				for !in.IsDelim(']') {
					var v1 ChangeItem
					easyjson39bd694bDecodeGithubComIhippikWalListenerListener1(in, &v1)
					out.Change = append(out.Change, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson39bd694bEncodeGithubComIhippikWalListenerListener(out *jwriter.Writer, in WalEvent) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"nextlsn\":"
		out.RawString(prefix[1:])
		out.String(string(in.NextLSN))
	}
	{
		const prefix string = ",\"change\":"
		out.RawString(prefix)
		if in.Change == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Change {
				if v2 > 0 {
					out.RawByte(',')
				}
				easyjson39bd694bEncodeGithubComIhippikWalListenerListener1(out, v3)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v WalEvent) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson39bd694bEncodeGithubComIhippikWalListenerListener(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v WalEvent) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson39bd694bEncodeGithubComIhippikWalListenerListener(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *WalEvent) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson39bd694bDecodeGithubComIhippikWalListenerListener(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *WalEvent) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson39bd694bDecodeGithubComIhippikWalListenerListener(l, v)
}
func easyjson39bd694bDecodeGithubComIhippikWalListenerListener1(in *jlexer.Lexer, out *ChangeItem) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "kind":
			out.Kind = string(in.String())
		case "schema":
			out.Schema = string(in.String())
		case "table":
			out.Table = string(in.String())
		case "columnnames":
			if in.IsNull() {
				in.Skip()
				out.ColumnNames = nil
			} else {
				in.Delim('[')
				if out.ColumnNames == nil {
					if !in.IsDelim(']') {
						out.ColumnNames = make([]string, 0, 4)
					} else {
						out.ColumnNames = []string{}
					}
				} else {
					out.ColumnNames = (out.ColumnNames)[:0]
				}
				for !in.IsDelim(']') {
					var v4 string
					v4 = string(in.String())
					out.ColumnNames = append(out.ColumnNames, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "columntypes":
			if in.IsNull() {
				in.Skip()
				out.ColumnTypes = nil
			} else {
				in.Delim('[')
				if out.ColumnTypes == nil {
					if !in.IsDelim(']') {
						out.ColumnTypes = make([]string, 0, 4)
					} else {
						out.ColumnTypes = []string{}
					}
				} else {
					out.ColumnTypes = (out.ColumnTypes)[:0]
				}
				for !in.IsDelim(']') {
					var v5 string
					v5 = string(in.String())
					out.ColumnTypes = append(out.ColumnTypes, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "columnvalues":
			if in.IsNull() {
				in.Skip()
				out.ColumnValues = nil
			} else {
				in.Delim('[')
				if out.ColumnValues == nil {
					if !in.IsDelim(']') {
						out.ColumnValues = make([]interface{}, 0, 4)
					} else {
						out.ColumnValues = []interface{}{}
					}
				} else {
					out.ColumnValues = (out.ColumnValues)[:0]
				}
				for !in.IsDelim(']') {
					var v6 interface{}
					if m, ok := v6.(easyjson.Unmarshaler); ok {
						m.UnmarshalEasyJSON(in)
					} else if m, ok := v6.(json.Unmarshaler); ok {
						_ = m.UnmarshalJSON(in.Raw())
					} else {
						v6 = in.Interface()
					}
					out.ColumnValues = append(out.ColumnValues, v6)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "oldkeys":
			easyjson39bd694bDecode(in, &out.OldKeys)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson39bd694bEncodeGithubComIhippikWalListenerListener1(out *jwriter.Writer, in ChangeItem) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"kind\":"
		out.RawString(prefix[1:])
		out.String(string(in.Kind))
	}
	{
		const prefix string = ",\"schema\":"
		out.RawString(prefix)
		out.String(string(in.Schema))
	}
	{
		const prefix string = ",\"table\":"
		out.RawString(prefix)
		out.String(string(in.Table))
	}
	{
		const prefix string = ",\"columnnames\":"
		out.RawString(prefix)
		if in.ColumnNames == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v7, v8 := range in.ColumnNames {
				if v7 > 0 {
					out.RawByte(',')
				}
				out.String(string(v8))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"columntypes\":"
		out.RawString(prefix)
		if in.ColumnTypes == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v9, v10 := range in.ColumnTypes {
				if v9 > 0 {
					out.RawByte(',')
				}
				out.String(string(v10))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"columnvalues\":"
		out.RawString(prefix)
		if in.ColumnValues == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v11, v12 := range in.ColumnValues {
				if v11 > 0 {
					out.RawByte(',')
				}
				if m, ok := v12.(easyjson.Marshaler); ok {
					m.MarshalEasyJSON(out)
				} else if m, ok := v12.(json.Marshaler); ok {
					out.Raw(m.MarshalJSON())
				} else {
					out.Raw(json.Marshal(v12))
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"oldkeys\":"
		out.RawString(prefix)
		easyjson39bd694bEncode(out, in.OldKeys)
	}
	out.RawByte('}')
}
func easyjson39bd694bDecode(in *jlexer.Lexer, out *struct {
	KeyNames  []string      `json:"keynames"`
	KeyTypes  []string      `json:"keytypes"`
	KeyValues []interface{} `json:"keyvalues"`
}) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "keynames":
			if in.IsNull() {
				in.Skip()
				out.KeyNames = nil
			} else {
				in.Delim('[')
				if out.KeyNames == nil {
					if !in.IsDelim(']') {
						out.KeyNames = make([]string, 0, 4)
					} else {
						out.KeyNames = []string{}
					}
				} else {
					out.KeyNames = (out.KeyNames)[:0]
				}
				for !in.IsDelim(']') {
					var v13 string
					v13 = string(in.String())
					out.KeyNames = append(out.KeyNames, v13)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "keytypes":
			if in.IsNull() {
				in.Skip()
				out.KeyTypes = nil
			} else {
				in.Delim('[')
				if out.KeyTypes == nil {
					if !in.IsDelim(']') {
						out.KeyTypes = make([]string, 0, 4)
					} else {
						out.KeyTypes = []string{}
					}
				} else {
					out.KeyTypes = (out.KeyTypes)[:0]
				}
				for !in.IsDelim(']') {
					var v14 string
					v14 = string(in.String())
					out.KeyTypes = append(out.KeyTypes, v14)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "keyvalues":
			if in.IsNull() {
				in.Skip()
				out.KeyValues = nil
			} else {
				in.Delim('[')
				if out.KeyValues == nil {
					if !in.IsDelim(']') {
						out.KeyValues = make([]interface{}, 0, 4)
					} else {
						out.KeyValues = []interface{}{}
					}
				} else {
					out.KeyValues = (out.KeyValues)[:0]
				}
				for !in.IsDelim(']') {
					var v15 interface{}
					if m, ok := v15.(easyjson.Unmarshaler); ok {
						m.UnmarshalEasyJSON(in)
					} else if m, ok := v15.(json.Unmarshaler); ok {
						_ = m.UnmarshalJSON(in.Raw())
					} else {
						v15 = in.Interface()
					}
					out.KeyValues = append(out.KeyValues, v15)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson39bd694bEncode(out *jwriter.Writer, in struct {
	KeyNames  []string      `json:"keynames"`
	KeyTypes  []string      `json:"keytypes"`
	KeyValues []interface{} `json:"keyvalues"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"keynames\":"
		out.RawString(prefix[1:])
		if in.KeyNames == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v16, v17 := range in.KeyNames {
				if v16 > 0 {
					out.RawByte(',')
				}
				out.String(string(v17))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"keytypes\":"
		out.RawString(prefix)
		if in.KeyTypes == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v18, v19 := range in.KeyTypes {
				if v18 > 0 {
					out.RawByte(',')
				}
				out.String(string(v19))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"keyvalues\":"
		out.RawString(prefix)
		if in.KeyValues == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v20, v21 := range in.KeyValues {
				if v20 > 0 {
					out.RawByte(',')
				}
				if m, ok := v21.(easyjson.Marshaler); ok {
					m.MarshalEasyJSON(out)
				} else if m, ok := v21.(json.Marshaler); ok {
					out.Raw(m.MarshalJSON())
				} else {
					out.Raw(json.Marshal(v21))
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
