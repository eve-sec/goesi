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

func easyjsonD1506562DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCorporationCorporationIdMiningExtractions200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCorporationCorporationIdMiningExtractions200OkList, 0, 1)
			} else {
				*out = GetCorporationCorporationIdMiningExtractions200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCorporationCorporationIdMiningExtractions200Ok
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
func easyjsonD1506562EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCorporationCorporationIdMiningExtractions200OkList) {
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
func (v GetCorporationCorporationIdMiningExtractions200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD1506562EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCorporationCorporationIdMiningExtractions200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD1506562EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCorporationCorporationIdMiningExtractions200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD1506562DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCorporationCorporationIdMiningExtractions200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD1506562DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjsonD1506562DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCorporationCorporationIdMiningExtractions200Ok) {
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
		case "chunk_arrival_time":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ChunkArrivalTime).UnmarshalJSON(data))
			}
		case "extraction_start_time":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ExtractionStartTime).UnmarshalJSON(data))
			}
		case "moon_id":
			out.MoonId = int32(in.Int32())
		case "natural_decay_time":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.NaturalDecayTime).UnmarshalJSON(data))
			}
		case "structure_id":
			out.StructureId = int64(in.Int64())
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
func easyjsonD1506562EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCorporationCorporationIdMiningExtractions200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if true {
		const prefix string = ",\"chunk_arrival_time\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.ChunkArrivalTime).MarshalJSON())
	}
	if true {
		const prefix string = ",\"extraction_start_time\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.ExtractionStartTime).MarshalJSON())
	}
	if in.MoonId != 0 {
		const prefix string = ",\"moon_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.MoonId))
	}
	if true {
		const prefix string = ",\"natural_decay_time\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.NaturalDecayTime).MarshalJSON())
	}
	if in.StructureId != 0 {
		const prefix string = ",\"structure_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.StructureId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCorporationCorporationIdMiningExtractions200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD1506562EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCorporationCorporationIdMiningExtractions200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD1506562EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCorporationCorporationIdMiningExtractions200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD1506562DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCorporationCorporationIdMiningExtractions200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD1506562DecodeGithubComAntihaxGoesiEsi1(l, v)
}
