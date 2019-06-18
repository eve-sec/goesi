// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package esi

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

func easyjson46677457DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *PostUniverseIdsSystemList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(PostUniverseIdsSystemList, 0, 2)
			} else {
				*out = PostUniverseIdsSystemList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 PostUniverseIdsSystem
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson46677457EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in PostUniverseIdsSystemList) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v PostUniverseIdsSystemList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson46677457EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostUniverseIdsSystemList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson46677457EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostUniverseIdsSystemList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson46677457DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostUniverseIdsSystemList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson46677457DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson46677457DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *PostUniverseIdsSystem) {
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
		case "id":
			out.Id = int32(in.Int32())
		case "name":
			out.Name = string(in.String())
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
func easyjson46677457EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in PostUniverseIdsSystem) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Id != 0 {
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Id))
	}
	if in.Name != "" {
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PostUniverseIdsSystem) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson46677457EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostUniverseIdsSystem) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson46677457EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostUniverseIdsSystem) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson46677457DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostUniverseIdsSystem) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson46677457DecodeGithubComAntihaxGoesiEsi1(l, v)
}
