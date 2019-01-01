// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package testdata

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

func easyjsonA96ca39cDecodeGithubComCorestoreioPkgSqlDmlgenTestdata(in *jlexer.Lexer, out *DmlgenTypesCollection) {
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
		case "data":
			if in.IsNull() {
				in.Skip()
				out.Data = nil
			} else {
				in.Delim('[')
				if out.Data == nil {
					if !in.IsDelim(']') {
						out.Data = make([]*DmlgenTypes, 0, 8)
					} else {
						out.Data = []*DmlgenTypes{}
					}
				} else {
					out.Data = (out.Data)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *DmlgenTypes
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(DmlgenTypes)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Data = append(out.Data, v1)
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
func easyjsonA96ca39cEncodeGithubComCorestoreioPkgSqlDmlgenTestdata(out *jwriter.Writer, in DmlgenTypesCollection) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Data) != 0 {
		const prefix string = ",\"data\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v2, v3 := range in.Data {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					(*v3).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DmlgenTypesCollection) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA96ca39cEncodeGithubComCorestoreioPkgSqlDmlgenTestdata(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DmlgenTypesCollection) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA96ca39cEncodeGithubComCorestoreioPkgSqlDmlgenTestdata(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DmlgenTypesCollection) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA96ca39cDecodeGithubComCorestoreioPkgSqlDmlgenTestdata(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DmlgenTypesCollection) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA96ca39cDecodeGithubComCorestoreioPkgSqlDmlgenTestdata(l, v)
}
func easyjsonA96ca39cDecodeGithubComCorestoreioPkgSqlDmlgenTestdata1(in *jlexer.Lexer, out *DmlgenTypes) {
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
			out.ID = int64(in.Int64())
		case "colBigint1":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ColBigint1).UnmarshalJSON(data))
			}
		case "colBigint2":
			out.ColBigint2 = int64(in.Int64())
		case "colBigint3":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ColBigint3).UnmarshalJSON(data))
			}
		case "colBigint4":
			out.ColBigint4 = uint64(in.Uint64())
		case "colBlob":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ColBlob).UnmarshalJSON(data))
			}
		case "colDate1":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ColDate1).UnmarshalJSON(data))
			}
		case "colDate2":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ColDate2).UnmarshalJSON(data))
			}
		case "colDatetime1":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ColDatetime1).UnmarshalJSON(data))
			}
		case "colDatetime2":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ColDatetime2).UnmarshalJSON(data))
			}
		case "colDecimal100":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ColDecimal100).UnmarshalJSON(data))
			}
		case "colDecimal124":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ColDecimal124).UnmarshalJSON(data))
			}
		case "price124a":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Price124a).UnmarshalJSON(data))
			}
		case "price124b":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Price124b).UnmarshalJSON(data))
			}
		case "colDecimal123":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ColDecimal123).UnmarshalJSON(data))
			}
		case "colDecimal206":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ColDecimal206).UnmarshalJSON(data))
			}
		case "colDecimal2412":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ColDecimal2412).UnmarshalJSON(data))
			}
		case "colFloat":
			out.ColFloat = float64(in.Float64())
		case "colInt1":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ColInt1).UnmarshalJSON(data))
			}
		case "colInt2":
			out.ColInt2 = int64(in.Int64())
		case "colInt3":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ColInt3).UnmarshalJSON(data))
			}
		case "colInt4":
			out.ColInt4 = uint64(in.Uint64())
		case "colLongtext1":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ColLongtext1).UnmarshalJSON(data))
			}
		case "colLongtext2":
			out.ColLongtext2 = string(in.String())
		case "colMediumblob":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ColMediumblob).UnmarshalJSON(data))
			}
		case "colMediumtext1":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ColMediumtext1).UnmarshalJSON(data))
			}
		case "colMediumtext2":
			out.ColMediumtext2 = string(in.String())
		case "colSmallint1":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ColSmallint1).UnmarshalJSON(data))
			}
		case "colSmallint2":
			out.ColSmallint2 = int64(in.Int64())
		case "colSmallint3":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ColSmallint3).UnmarshalJSON(data))
			}
		case "colSmallint4":
			out.ColSmallint4 = uint64(in.Uint64())
		case "hasSmallint5":
			out.HasSmallint5 = bool(in.Bool())
		case "isSmallint5":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.IsSmallint5).UnmarshalJSON(data))
			}
		case "colText":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ColText).UnmarshalJSON(data))
			}
		case "colTimestamp1":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ColTimestamp1).UnmarshalJSON(data))
			}
		case "colTimestamp2":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ColTimestamp2).UnmarshalJSON(data))
			}
		case "colTinyint1":
			out.ColTinyint1 = int64(in.Int64())
		case "colVarchar1":
			out.ColVarchar1 = string(in.String())
		case "colVarchar100":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ColVarchar100).UnmarshalJSON(data))
			}
		case "colVarchar16":
			out.ColVarchar16 = string(in.String())
		case "colChar1":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ColChar1).UnmarshalJSON(data))
			}
		case "colChar2":
			out.ColChar2 = string(in.String())
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
func easyjsonA96ca39cEncodeGithubComCorestoreioPkgSqlDmlgenTestdata1(out *jwriter.Writer, in DmlgenTypes) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != 0 {
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ID))
	}
	if true {
		const prefix string = ",\"colBigint1\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.ColBigint1).MarshalJSON())
	}
	if in.ColBigint2 != 0 {
		const prefix string = ",\"colBigint2\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ColBigint2))
	}
	if true {
		const prefix string = ",\"colBigint3\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.ColBigint3).MarshalJSON())
	}
	if in.ColBigint4 != 0 {
		const prefix string = ",\"colBigint4\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.ColBigint4))
	}
	if true {
		const prefix string = ",\"colBlob\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.ColBlob).MarshalJSON())
	}
	if true {
		const prefix string = ",\"colDate1\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.ColDate1).MarshalJSON())
	}
	if true {
		const prefix string = ",\"colDate2\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.ColDate2).MarshalJSON())
	}
	if true {
		const prefix string = ",\"colDatetime1\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.ColDatetime1).MarshalJSON())
	}
	if true {
		const prefix string = ",\"colDatetime2\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.ColDatetime2).MarshalJSON())
	}
	if true {
		const prefix string = ",\"colDecimal100\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.ColDecimal100).MarshalJSON())
	}
	if true {
		const prefix string = ",\"colDecimal124\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.ColDecimal124).MarshalJSON())
	}
	if true {
		const prefix string = ",\"price124a\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Price124a).MarshalJSON())
	}
	if true {
		const prefix string = ",\"price124b\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Price124b).MarshalJSON())
	}
	if true {
		const prefix string = ",\"colDecimal123\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.ColDecimal123).MarshalJSON())
	}
	if true {
		const prefix string = ",\"colDecimal206\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.ColDecimal206).MarshalJSON())
	}
	if true {
		const prefix string = ",\"colDecimal2412\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.ColDecimal2412).MarshalJSON())
	}
	if in.ColFloat != 0 {
		const prefix string = ",\"colFloat\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.ColFloat))
	}
	if true {
		const prefix string = ",\"colInt1\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.ColInt1).MarshalJSON())
	}
	if in.ColInt2 != 0 {
		const prefix string = ",\"colInt2\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ColInt2))
	}
	if true {
		const prefix string = ",\"colInt3\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.ColInt3).MarshalJSON())
	}
	if in.ColInt4 != 0 {
		const prefix string = ",\"colInt4\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.ColInt4))
	}
	if true {
		const prefix string = ",\"colLongtext1\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.ColLongtext1).MarshalJSON())
	}
	if in.ColLongtext2 != "" {
		const prefix string = ",\"colLongtext2\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ColLongtext2))
	}
	if true {
		const prefix string = ",\"colMediumblob\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.ColMediumblob).MarshalJSON())
	}
	if true {
		const prefix string = ",\"colMediumtext1\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.ColMediumtext1).MarshalJSON())
	}
	if in.ColMediumtext2 != "" {
		const prefix string = ",\"colMediumtext2\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ColMediumtext2))
	}
	if true {
		const prefix string = ",\"colSmallint1\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.ColSmallint1).MarshalJSON())
	}
	if in.ColSmallint2 != 0 {
		const prefix string = ",\"colSmallint2\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ColSmallint2))
	}
	if true {
		const prefix string = ",\"colSmallint3\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.ColSmallint3).MarshalJSON())
	}
	if in.ColSmallint4 != 0 {
		const prefix string = ",\"colSmallint4\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.ColSmallint4))
	}
	if in.HasSmallint5 {
		const prefix string = ",\"hasSmallint5\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.HasSmallint5))
	}
	if true {
		const prefix string = ",\"isSmallint5\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.IsSmallint5).MarshalJSON())
	}
	if true {
		const prefix string = ",\"colText\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.ColText).MarshalJSON())
	}
	if true {
		const prefix string = ",\"colTimestamp1\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.ColTimestamp1).MarshalJSON())
	}
	if true {
		const prefix string = ",\"colTimestamp2\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.ColTimestamp2).MarshalJSON())
	}
	if in.ColTinyint1 != 0 {
		const prefix string = ",\"colTinyint1\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ColTinyint1))
	}
	if in.ColVarchar1 != "" {
		const prefix string = ",\"colVarchar1\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ColVarchar1))
	}
	if true {
		const prefix string = ",\"colVarchar100\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.ColVarchar100).MarshalJSON())
	}
	if in.ColVarchar16 != "" {
		const prefix string = ",\"colVarchar16\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ColVarchar16))
	}
	if true {
		const prefix string = ",\"colChar1\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.ColChar1).MarshalJSON())
	}
	if in.ColChar2 != "" {
		const prefix string = ",\"colChar2\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ColChar2))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DmlgenTypes) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA96ca39cEncodeGithubComCorestoreioPkgSqlDmlgenTestdata1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DmlgenTypes) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA96ca39cEncodeGithubComCorestoreioPkgSqlDmlgenTestdata1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DmlgenTypes) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA96ca39cDecodeGithubComCorestoreioPkgSqlDmlgenTestdata1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DmlgenTypes) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA96ca39cDecodeGithubComCorestoreioPkgSqlDmlgenTestdata1(l, v)
}
func easyjsonA96ca39cDecodeGithubComCorestoreioPkgSqlDmlgenTestdata2(in *jlexer.Lexer, out *CoreConfigDataCollection) {
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
		case "data":
			if in.IsNull() {
				in.Skip()
				out.Data = nil
			} else {
				in.Delim('[')
				if out.Data == nil {
					if !in.IsDelim(']') {
						out.Data = make([]*CoreConfigData, 0, 8)
					} else {
						out.Data = []*CoreConfigData{}
					}
				} else {
					out.Data = (out.Data)[:0]
				}
				for !in.IsDelim(']') {
					var v4 *CoreConfigData
					if in.IsNull() {
						in.Skip()
						v4 = nil
					} else {
						if v4 == nil {
							v4 = new(CoreConfigData)
						}
						(*v4).UnmarshalEasyJSON(in)
					}
					out.Data = append(out.Data, v4)
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
func easyjsonA96ca39cEncodeGithubComCorestoreioPkgSqlDmlgenTestdata2(out *jwriter.Writer, in CoreConfigDataCollection) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Data) != 0 {
		const prefix string = ",\"data\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v5, v6 := range in.Data {
				if v5 > 0 {
					out.RawByte(',')
				}
				if v6 == nil {
					out.RawString("null")
				} else {
					(*v6).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CoreConfigDataCollection) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA96ca39cEncodeGithubComCorestoreioPkgSqlDmlgenTestdata2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CoreConfigDataCollection) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA96ca39cEncodeGithubComCorestoreioPkgSqlDmlgenTestdata2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CoreConfigDataCollection) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA96ca39cDecodeGithubComCorestoreioPkgSqlDmlgenTestdata2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CoreConfigDataCollection) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA96ca39cDecodeGithubComCorestoreioPkgSqlDmlgenTestdata2(l, v)
}
func easyjsonA96ca39cDecodeGithubComCorestoreioPkgSqlDmlgenTestdata3(in *jlexer.Lexer, out *CoreConfigData) {
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
		case "config_id":
			out.ConfigID = uint64(in.Uint64())
		case "scope":
			out.Scope = string(in.String())
		case "scope_id":
			out.ScopeID = int64(in.Int64())
		case "expires":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Expires).UnmarshalJSON(data))
			}
		case "x_path":
			out.Path = string(in.String())
		case "value":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Value).UnmarshalJSON(data))
			}
		case "version_ts":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.VersionTs).UnmarshalJSON(data))
			}
		case "version_te":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.VersionTe).UnmarshalJSON(data))
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
func easyjsonA96ca39cEncodeGithubComCorestoreioPkgSqlDmlgenTestdata3(out *jwriter.Writer, in CoreConfigData) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ConfigID != 0 {
		const prefix string = ",\"config_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.ConfigID))
	}
	if in.Scope != "" {
		const prefix string = ",\"scope\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Scope))
	}
	if in.ScopeID != 0 {
		const prefix string = ",\"scope_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ScopeID))
	}
	if true {
		const prefix string = ",\"expires\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Expires).MarshalJSON())
	}
	if in.Path != "" {
		const prefix string = ",\"x_path\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Path))
	}
	if true {
		const prefix string = ",\"value\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Value).MarshalJSON())
	}
	if true {
		const prefix string = ",\"version_ts\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.VersionTs).MarshalJSON())
	}
	if true {
		const prefix string = ",\"version_te\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.VersionTe).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CoreConfigData) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA96ca39cEncodeGithubComCorestoreioPkgSqlDmlgenTestdata3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CoreConfigData) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA96ca39cEncodeGithubComCorestoreioPkgSqlDmlgenTestdata3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CoreConfigData) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA96ca39cDecodeGithubComCorestoreioPkgSqlDmlgenTestdata3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CoreConfigData) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA96ca39cDecodeGithubComCorestoreioPkgSqlDmlgenTestdata3(l, v)
}