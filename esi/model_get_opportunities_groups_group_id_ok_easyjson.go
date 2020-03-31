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

func easyjson2bed092DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetOpportunitiesGroupsGroupIdOkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetOpportunitiesGroupsGroupIdOkList, 0, 0)
			} else {
				*out = GetOpportunitiesGroupsGroupIdOkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetOpportunitiesGroupsGroupIdOk
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
func easyjson2bed092EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetOpportunitiesGroupsGroupIdOkList) {
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
func (v GetOpportunitiesGroupsGroupIdOkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson2bed092EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetOpportunitiesGroupsGroupIdOkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson2bed092EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetOpportunitiesGroupsGroupIdOkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson2bed092DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetOpportunitiesGroupsGroupIdOkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson2bed092DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson2bed092DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetOpportunitiesGroupsGroupIdOk) {
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
		case "connected_groups":
			if in.IsNull() {
				in.Skip()
				out.ConnectedGroups = nil
			} else {
				in.Delim('[')
				if out.ConnectedGroups == nil {
					if !in.IsDelim(']') {
						out.ConnectedGroups = make([]int32, 0, 16)
					} else {
						out.ConnectedGroups = []int32{}
					}
				} else {
					out.ConnectedGroups = (out.ConnectedGroups)[:0]
				}
				for !in.IsDelim(']') {
					var v4 int32
					v4 = int32(in.Int32())
					out.ConnectedGroups = append(out.ConnectedGroups, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "description":
			out.Description = string(in.String())
		case "group_id":
			out.GroupId = int32(in.Int32())
		case "name":
			out.Name = string(in.String())
		case "notification":
			out.Notification = string(in.String())
		case "required_tasks":
			if in.IsNull() {
				in.Skip()
				out.RequiredTasks = nil
			} else {
				in.Delim('[')
				if out.RequiredTasks == nil {
					if !in.IsDelim(']') {
						out.RequiredTasks = make([]int32, 0, 16)
					} else {
						out.RequiredTasks = []int32{}
					}
				} else {
					out.RequiredTasks = (out.RequiredTasks)[:0]
				}
				for !in.IsDelim(']') {
					var v5 int32
					v5 = int32(in.Int32())
					out.RequiredTasks = append(out.RequiredTasks, v5)
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
func easyjson2bed092EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetOpportunitiesGroupsGroupIdOk) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.ConnectedGroups) != 0 {
		const prefix string = ",\"connected_groups\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('[')
			for v6, v7 := range in.ConnectedGroups {
				if v6 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v7))
			}
			out.RawByte(']')
		}
	}
	if in.Description != "" {
		const prefix string = ",\"description\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Description))
	}
	if in.GroupId != 0 {
		const prefix string = ",\"group_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.GroupId))
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
	if in.Notification != "" {
		const prefix string = ",\"notification\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Notification))
	}
	if len(in.RequiredTasks) != 0 {
		const prefix string = ",\"required_tasks\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v8, v9 := range in.RequiredTasks {
				if v8 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v9))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetOpportunitiesGroupsGroupIdOk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson2bed092EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetOpportunitiesGroupsGroupIdOk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson2bed092EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetOpportunitiesGroupsGroupIdOk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson2bed092DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetOpportunitiesGroupsGroupIdOk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson2bed092DecodeGithubComAntihaxGoesiEsi1(l, v)
}
