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

func easyjson9d70c3a8DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCorporationsCorporationIdCustomsOffices200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCorporationsCorporationIdCustomsOffices200OkList, 0, 1)
			} else {
				*out = GetCorporationsCorporationIdCustomsOffices200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCorporationsCorporationIdCustomsOffices200Ok
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
func easyjson9d70c3a8EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCorporationsCorporationIdCustomsOffices200OkList) {
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
func (v GetCorporationsCorporationIdCustomsOffices200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9d70c3a8EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCorporationsCorporationIdCustomsOffices200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9d70c3a8EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCorporationsCorporationIdCustomsOffices200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9d70c3a8DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCorporationsCorporationIdCustomsOffices200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9d70c3a8DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson9d70c3a8DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCorporationsCorporationIdCustomsOffices200Ok) {
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
		case "alliance_tax_rate":
			out.AllianceTaxRate = float32(in.Float32())
		case "allow_access_with_standings":
			out.AllowAccessWithStandings = bool(in.Bool())
		case "allow_alliance_access":
			out.AllowAllianceAccess = bool(in.Bool())
		case "bad_standing_tax_rate":
			out.BadStandingTaxRate = float32(in.Float32())
		case "corporation_tax_rate":
			out.CorporationTaxRate = float32(in.Float32())
		case "excellent_standing_tax_rate":
			out.ExcellentStandingTaxRate = float32(in.Float32())
		case "good_standing_tax_rate":
			out.GoodStandingTaxRate = float32(in.Float32())
		case "neutral_standing_tax_rate":
			out.NeutralStandingTaxRate = float32(in.Float32())
		case "office_id":
			out.OfficeId = int64(in.Int64())
		case "reinforce_exit_end":
			out.ReinforceExitEnd = int32(in.Int32())
		case "reinforce_exit_start":
			out.ReinforceExitStart = int32(in.Int32())
		case "standing_level":
			out.StandingLevel = string(in.String())
		case "system_id":
			out.SystemId = int32(in.Int32())
		case "terrible_standing_tax_rate":
			out.TerribleStandingTaxRate = float32(in.Float32())
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
func easyjson9d70c3a8EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCorporationsCorporationIdCustomsOffices200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if in.AllianceTaxRate != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"alliance_tax_rate\":")
		out.Float32(float32(in.AllianceTaxRate))
	}
	if in.AllowAccessWithStandings {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"allow_access_with_standings\":")
		out.Bool(bool(in.AllowAccessWithStandings))
	}
	if in.AllowAllianceAccess {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"allow_alliance_access\":")
		out.Bool(bool(in.AllowAllianceAccess))
	}
	if in.BadStandingTaxRate != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"bad_standing_tax_rate\":")
		out.Float32(float32(in.BadStandingTaxRate))
	}
	if in.CorporationTaxRate != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"corporation_tax_rate\":")
		out.Float32(float32(in.CorporationTaxRate))
	}
	if in.ExcellentStandingTaxRate != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"excellent_standing_tax_rate\":")
		out.Float32(float32(in.ExcellentStandingTaxRate))
	}
	if in.GoodStandingTaxRate != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"good_standing_tax_rate\":")
		out.Float32(float32(in.GoodStandingTaxRate))
	}
	if in.NeutralStandingTaxRate != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"neutral_standing_tax_rate\":")
		out.Float32(float32(in.NeutralStandingTaxRate))
	}
	if in.OfficeId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"office_id\":")
		out.Int64(int64(in.OfficeId))
	}
	if in.ReinforceExitEnd != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"reinforce_exit_end\":")
		out.Int32(int32(in.ReinforceExitEnd))
	}
	if in.ReinforceExitStart != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"reinforce_exit_start\":")
		out.Int32(int32(in.ReinforceExitStart))
	}
	if in.StandingLevel != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"standing_level\":")
		out.String(string(in.StandingLevel))
	}
	if in.SystemId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"system_id\":")
		out.Int32(int32(in.SystemId))
	}
	if in.TerribleStandingTaxRate != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"terrible_standing_tax_rate\":")
		out.Float32(float32(in.TerribleStandingTaxRate))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCorporationsCorporationIdCustomsOffices200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9d70c3a8EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCorporationsCorporationIdCustomsOffices200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9d70c3a8EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCorporationsCorporationIdCustomsOffices200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9d70c3a8DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCorporationsCorporationIdCustomsOffices200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9d70c3a8DecodeGithubComAntihaxGoesiEsi1(l, v)
}
