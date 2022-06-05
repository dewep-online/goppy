// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package http

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

func easyjson124e5e3DecodeGithubComDewepOnlineGoppyPluginsHttp(in *jlexer.Lexer, out *ErrMessage) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "code":
			out.InternalCode = string(in.String())
		case "msg":
			out.Message = string(in.String())
		case "ctx":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Ctx = make(map[string]interface{})
				} else {
					out.Ctx = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v1 interface{}
					if m, ok := v1.(easyjson.Unmarshaler); ok {
						m.UnmarshalEasyJSON(in)
					} else if m, ok := v1.(json.Unmarshaler); ok {
						_ = m.UnmarshalJSON(in.Raw())
					} else {
						v1 = in.Interface()
					}
					(out.Ctx)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
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
func easyjson124e5e3EncodeGithubComDewepOnlineGoppyPluginsHttp(out *jwriter.Writer, in ErrMessage) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"code\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.InternalCode))
	}
	{
		const prefix string = ",\"msg\":"
		out.RawString(prefix)
		out.String(string(in.Message))
	}
	if len(in.Ctx) != 0 {
		const prefix string = ",\"ctx\":"
		out.RawString(prefix)
		{
			out.RawByte('{')
			v2First := true
			for v2Name, v2Value := range in.Ctx {
				if v2First {
					v2First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v2Name))
				out.RawByte(':')
				if m, ok := v2Value.(easyjson.Marshaler); ok {
					m.MarshalEasyJSON(out)
				} else if m, ok := v2Value.(json.Marshaler); ok {
					out.Raw(m.MarshalJSON())
				} else {
					out.Raw(json.Marshal(v2Value))
				}
			}
			out.RawByte('}')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ErrMessage) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson124e5e3EncodeGithubComDewepOnlineGoppyPluginsHttp(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ErrMessage) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson124e5e3EncodeGithubComDewepOnlineGoppyPluginsHttp(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ErrMessage) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson124e5e3DecodeGithubComDewepOnlineGoppyPluginsHttp(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ErrMessage) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson124e5e3DecodeGithubComDewepOnlineGoppyPluginsHttp(l, v)
}