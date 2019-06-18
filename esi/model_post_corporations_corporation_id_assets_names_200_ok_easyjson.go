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

func easyjson5cc83670DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *PostCorporationsCorporationIdAssetsNames200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(PostCorporationsCorporationIdAssetsNames200OkList, 0, 2)
			} else {
				*out = PostCorporationsCorporationIdAssetsNames200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 PostCorporationsCorporationIdAssetsNames200Ok
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
func easyjson5cc83670EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in PostCorporationsCorporationIdAssetsNames200OkList) {
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
func (v PostCorporationsCorporationIdAssetsNames200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5cc83670EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostCorporationsCorporationIdAssetsNames200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5cc83670EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostCorporationsCorporationIdAssetsNames200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5cc83670DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostCorporationsCorporationIdAssetsNames200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5cc83670DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson5cc83670DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *PostCorporationsCorporationIdAssetsNames200Ok) {
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
		case "item_id":
			out.ItemId = int64(in.Int64())
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
func easyjson5cc83670EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in PostCorporationsCorporationIdAssetsNames200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ItemId != 0 {
		const prefix string = ",\"item_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ItemId))
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
func (v PostCorporationsCorporationIdAssetsNames200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5cc83670EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostCorporationsCorporationIdAssetsNames200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5cc83670EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostCorporationsCorporationIdAssetsNames200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5cc83670DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostCorporationsCorporationIdAssetsNames200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5cc83670DecodeGithubComAntihaxGoesiEsi1(l, v)
}
