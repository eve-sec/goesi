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

func easyjsonEec0496dDecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdOrders200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdOrders200OkList, 0, 0)
			} else {
				*out = GetCharactersCharacterIdOrders200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdOrders200Ok
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
func easyjsonEec0496dEncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdOrders200OkList) {
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
func (v GetCharactersCharacterIdOrders200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEec0496dEncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdOrders200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEec0496dEncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdOrders200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEec0496dDecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdOrders200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEec0496dDecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjsonEec0496dDecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdOrders200Ok) {
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
		case "duration":
			out.Duration = int32(in.Int32())
		case "escrow":
			out.Escrow = float64(in.Float64())
		case "is_buy_order":
			out.IsBuyOrder = bool(in.Bool())
		case "is_corporation":
			out.IsCorporation = bool(in.Bool())
		case "issued":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Issued).UnmarshalJSON(data))
			}
		case "location_id":
			out.LocationId = int64(in.Int64())
		case "min_volume":
			out.MinVolume = int32(in.Int32())
		case "order_id":
			out.OrderId = int64(in.Int64())
		case "price":
			out.Price = float64(in.Float64())
		case "range":
			out.Range_ = string(in.String())
		case "region_id":
			out.RegionId = int32(in.Int32())
		case "type_id":
			out.TypeId = int32(in.Int32())
		case "volume_remain":
			out.VolumeRemain = int32(in.Int32())
		case "volume_total":
			out.VolumeTotal = int32(in.Int32())
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
func easyjsonEec0496dEncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdOrders200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Duration != 0 {
		const prefix string = ",\"duration\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.Duration))
	}
	if in.Escrow != 0 {
		const prefix string = ",\"escrow\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.Escrow))
	}
	if in.IsBuyOrder {
		const prefix string = ",\"is_buy_order\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.IsBuyOrder))
	}
	if in.IsCorporation {
		const prefix string = ",\"is_corporation\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.IsCorporation))
	}
	if true {
		const prefix string = ",\"issued\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Issued).MarshalJSON())
	}
	if in.LocationId != 0 {
		const prefix string = ",\"location_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.LocationId))
	}
	if in.MinVolume != 0 {
		const prefix string = ",\"min_volume\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.MinVolume))
	}
	if in.OrderId != 0 {
		const prefix string = ",\"order_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.OrderId))
	}
	if in.Price != 0 {
		const prefix string = ",\"price\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.Price))
	}
	if in.Range_ != "" {
		const prefix string = ",\"range\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Range_))
	}
	if in.RegionId != 0 {
		const prefix string = ",\"region_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.RegionId))
	}
	if in.TypeId != 0 {
		const prefix string = ",\"type_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.TypeId))
	}
	if in.VolumeRemain != 0 {
		const prefix string = ",\"volume_remain\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.VolumeRemain))
	}
	if in.VolumeTotal != 0 {
		const prefix string = ",\"volume_total\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.VolumeTotal))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdOrders200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEec0496dEncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdOrders200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEec0496dEncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdOrders200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEec0496dDecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdOrders200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEec0496dDecodeGithubComAntihaxGoesiEsi1(l, v)
}
