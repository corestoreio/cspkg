// AUTOGENERATED FILE: easyjson marshaller/unmarshallers.

package problem

import (
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

func easyjson1f631ffDecodeGithubComCorestoreioCsfwNetProblem(in *jlexer.Lexer, out *Detail) {
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
		case "type":
			out.Type = in.String()
		case "title":
			out.Title = in.String()
		case "status":
			out.Status = in.Int()
		case "detail":
			out.Detail = in.String()
		case "instance":
			out.Instance = in.String()
		case "cause":
			if in.IsNull() {
				in.Skip()
				out.Cause = nil
			} else {
				if out.Cause == nil {
					out.Cause = new(Detail)
				}
				(*out.Cause).UnmarshalEasyJSON(in)
			}
		case "extension":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Extension = make([]string, 0, 4)
				} else {
					out.Extension = nil
				}
				for !in.IsDelim('}') {
					key := in.String()
					in.WantColon()
					v1 := in.String()
					out.Extension = append(out.Extension, key, v1)
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

func easyjson1f631ffEncodeGithubComCorestoreioCsfwNetProblem(out *jwriter.Writer, in Detail) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Type != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"type\":")
		out.String(in.Type)
	}
	if in.Title != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"title\":")
		out.String(in.Title)
	}
	if in.Status != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"status\":")
		out.Int(in.Status)
	}
	if in.Detail != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"detail\":")
		out.String(in.Detail)
	}
	if in.Instance != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"instance\":")
		out.String(in.Instance)
	}
	if in.Cause != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"cause\":")
		if in.Cause == nil {
			out.RawString("null")
		} else {
			(*in.Cause).MarshalEasyJSON(out)
		}
	}
	if len(in.Extension) != 0 {
		if !first {
			out.RawByte(',')
		}
		out.RawString("\"extension\":")
		if in.Extension == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('{')
			le := len(in.Extension)
			for i := 0; i < le; i = i + 2 {
				j := i + 1
				out.String(in.Extension[i])
				out.RawByte(':')
				out.String(in.Extension[j])
				if j < le-1 {
					out.RawByte(',')
				}
			}
			out.RawByte('}')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Detail) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1f631ffEncodeGithubComCorestoreioCsfwNetProblem(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Detail) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1f631ffEncodeGithubComCorestoreioCsfwNetProblem(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Detail) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1f631ffDecodeGithubComCorestoreioCsfwNetProblem(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Detail) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1f631ffDecodeGithubComCorestoreioCsfwNetProblem(l, v)
}
