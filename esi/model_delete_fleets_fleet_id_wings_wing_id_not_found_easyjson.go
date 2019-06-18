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

func easyjson5de13898DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *DeleteFleetsFleetIdWingsWingIdNotFoundList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(DeleteFleetsFleetIdWingsWingIdNotFoundList, 0, 4)
			} else {
				*out = DeleteFleetsFleetIdWingsWingIdNotFoundList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 DeleteFleetsFleetIdWingsWingIdNotFound
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
func easyjson5de13898EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in DeleteFleetsFleetIdWingsWingIdNotFoundList) {
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
func (v DeleteFleetsFleetIdWingsWingIdNotFoundList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5de13898EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DeleteFleetsFleetIdWingsWingIdNotFoundList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5de13898EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DeleteFleetsFleetIdWingsWingIdNotFoundList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5de13898DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DeleteFleetsFleetIdWingsWingIdNotFoundList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5de13898DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson5de13898DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *DeleteFleetsFleetIdWingsWingIdNotFound) {
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
		case "error":
			out.Error_ = string(in.String())
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
func easyjson5de13898EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in DeleteFleetsFleetIdWingsWingIdNotFound) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Error_ != "" {
		const prefix string = ",\"error\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Error_))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DeleteFleetsFleetIdWingsWingIdNotFound) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5de13898EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DeleteFleetsFleetIdWingsWingIdNotFound) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5de13898EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DeleteFleetsFleetIdWingsWingIdNotFound) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5de13898DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DeleteFleetsFleetIdWingsWingIdNotFound) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5de13898DecodeGithubComAntihaxGoesiEsi1(l, v)
}
