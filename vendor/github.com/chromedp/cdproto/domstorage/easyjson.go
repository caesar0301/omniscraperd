// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package domstorage

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

func easyjsonC5a4559bDecodeGithubComChromedpCdprotoDomstorage(in *jlexer.Lexer, out *StorageID) {
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
		case "securityOrigin":
			out.SecurityOrigin = string(in.String())
		case "isLocalStorage":
			out.IsLocalStorage = bool(in.Bool())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoDomstorage(out *jwriter.Writer, in StorageID) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"securityOrigin\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.SecurityOrigin))
	}
	{
		const prefix string = ",\"isLocalStorage\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.IsLocalStorage))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v StorageID) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoDomstorage(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v StorageID) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoDomstorage(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *StorageID) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoDomstorage(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *StorageID) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoDomstorage(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoDomstorage1(in *jlexer.Lexer, out *SetDOMStorageItemParams) {
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
		case "storageId":
			if in.IsNull() {
				in.Skip()
				out.StorageID = nil
			} else {
				if out.StorageID == nil {
					out.StorageID = new(StorageID)
				}
				(*out.StorageID).UnmarshalEasyJSON(in)
			}
		case "key":
			out.Key = string(in.String())
		case "value":
			out.Value = string(in.String())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoDomstorage1(out *jwriter.Writer, in SetDOMStorageItemParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"storageId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.StorageID == nil {
			out.RawString("null")
		} else {
			(*in.StorageID).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"key\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Key))
	}
	{
		const prefix string = ",\"value\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Value))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SetDOMStorageItemParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoDomstorage1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SetDOMStorageItemParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoDomstorage1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SetDOMStorageItemParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoDomstorage1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SetDOMStorageItemParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoDomstorage1(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoDomstorage2(in *jlexer.Lexer, out *RemoveDOMStorageItemParams) {
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
		case "storageId":
			if in.IsNull() {
				in.Skip()
				out.StorageID = nil
			} else {
				if out.StorageID == nil {
					out.StorageID = new(StorageID)
				}
				(*out.StorageID).UnmarshalEasyJSON(in)
			}
		case "key":
			out.Key = string(in.String())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoDomstorage2(out *jwriter.Writer, in RemoveDOMStorageItemParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"storageId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.StorageID == nil {
			out.RawString("null")
		} else {
			(*in.StorageID).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"key\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Key))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RemoveDOMStorageItemParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoDomstorage2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RemoveDOMStorageItemParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoDomstorage2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RemoveDOMStorageItemParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoDomstorage2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RemoveDOMStorageItemParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoDomstorage2(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoDomstorage3(in *jlexer.Lexer, out *GetDOMStorageItemsReturns) {
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
		case "entries":
			if in.IsNull() {
				in.Skip()
				out.Entries = nil
			} else {
				in.Delim('[')
				if out.Entries == nil {
					if !in.IsDelim(']') {
						out.Entries = make([]Item, 0, 2)
					} else {
						out.Entries = []Item{}
					}
				} else {
					out.Entries = (out.Entries)[:0]
				}
				for !in.IsDelim(']') {
					var v1 Item
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						in.Delim('[')
						if v1 == nil {
							if !in.IsDelim(']') {
								v1 = make(Item, 0, 4)
							} else {
								v1 = Item{}
							}
						} else {
							v1 = (v1)[:0]
						}
						for !in.IsDelim(']') {
							var v2 string
							v2 = string(in.String())
							v1 = append(v1, v2)
							in.WantComma()
						}
						in.Delim(']')
					}
					out.Entries = append(out.Entries, v1)
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoDomstorage3(out *jwriter.Writer, in GetDOMStorageItemsReturns) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Entries) != 0 {
		const prefix string = ",\"entries\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v3, v4 := range in.Entries {
				if v3 > 0 {
					out.RawByte(',')
				}
				if v4 == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
					out.RawString("null")
				} else {
					out.RawByte('[')
					for v5, v6 := range v4 {
						if v5 > 0 {
							out.RawByte(',')
						}
						out.String(string(v6))
					}
					out.RawByte(']')
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetDOMStorageItemsReturns) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoDomstorage3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetDOMStorageItemsReturns) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoDomstorage3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetDOMStorageItemsReturns) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoDomstorage3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetDOMStorageItemsReturns) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoDomstorage3(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoDomstorage4(in *jlexer.Lexer, out *GetDOMStorageItemsParams) {
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
		case "storageId":
			if in.IsNull() {
				in.Skip()
				out.StorageID = nil
			} else {
				if out.StorageID == nil {
					out.StorageID = new(StorageID)
				}
				(*out.StorageID).UnmarshalEasyJSON(in)
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoDomstorage4(out *jwriter.Writer, in GetDOMStorageItemsParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"storageId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.StorageID == nil {
			out.RawString("null")
		} else {
			(*in.StorageID).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetDOMStorageItemsParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoDomstorage4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetDOMStorageItemsParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoDomstorage4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetDOMStorageItemsParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoDomstorage4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetDOMStorageItemsParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoDomstorage4(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoDomstorage5(in *jlexer.Lexer, out *EventDomStorageItemsCleared) {
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
		case "storageId":
			if in.IsNull() {
				in.Skip()
				out.StorageID = nil
			} else {
				if out.StorageID == nil {
					out.StorageID = new(StorageID)
				}
				(*out.StorageID).UnmarshalEasyJSON(in)
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoDomstorage5(out *jwriter.Writer, in EventDomStorageItemsCleared) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"storageId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.StorageID == nil {
			out.RawString("null")
		} else {
			(*in.StorageID).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EventDomStorageItemsCleared) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoDomstorage5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EventDomStorageItemsCleared) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoDomstorage5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EventDomStorageItemsCleared) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoDomstorage5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EventDomStorageItemsCleared) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoDomstorage5(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoDomstorage6(in *jlexer.Lexer, out *EventDomStorageItemUpdated) {
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
		case "storageId":
			if in.IsNull() {
				in.Skip()
				out.StorageID = nil
			} else {
				if out.StorageID == nil {
					out.StorageID = new(StorageID)
				}
				(*out.StorageID).UnmarshalEasyJSON(in)
			}
		case "key":
			out.Key = string(in.String())
		case "oldValue":
			out.OldValue = string(in.String())
		case "newValue":
			out.NewValue = string(in.String())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoDomstorage6(out *jwriter.Writer, in EventDomStorageItemUpdated) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"storageId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.StorageID == nil {
			out.RawString("null")
		} else {
			(*in.StorageID).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"key\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Key))
	}
	{
		const prefix string = ",\"oldValue\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.OldValue))
	}
	{
		const prefix string = ",\"newValue\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.NewValue))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EventDomStorageItemUpdated) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoDomstorage6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EventDomStorageItemUpdated) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoDomstorage6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EventDomStorageItemUpdated) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoDomstorage6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EventDomStorageItemUpdated) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoDomstorage6(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoDomstorage7(in *jlexer.Lexer, out *EventDomStorageItemRemoved) {
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
		case "storageId":
			if in.IsNull() {
				in.Skip()
				out.StorageID = nil
			} else {
				if out.StorageID == nil {
					out.StorageID = new(StorageID)
				}
				(*out.StorageID).UnmarshalEasyJSON(in)
			}
		case "key":
			out.Key = string(in.String())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoDomstorage7(out *jwriter.Writer, in EventDomStorageItemRemoved) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"storageId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.StorageID == nil {
			out.RawString("null")
		} else {
			(*in.StorageID).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"key\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Key))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EventDomStorageItemRemoved) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoDomstorage7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EventDomStorageItemRemoved) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoDomstorage7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EventDomStorageItemRemoved) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoDomstorage7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EventDomStorageItemRemoved) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoDomstorage7(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoDomstorage8(in *jlexer.Lexer, out *EventDomStorageItemAdded) {
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
		case "storageId":
			if in.IsNull() {
				in.Skip()
				out.StorageID = nil
			} else {
				if out.StorageID == nil {
					out.StorageID = new(StorageID)
				}
				(*out.StorageID).UnmarshalEasyJSON(in)
			}
		case "key":
			out.Key = string(in.String())
		case "newValue":
			out.NewValue = string(in.String())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoDomstorage8(out *jwriter.Writer, in EventDomStorageItemAdded) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"storageId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.StorageID == nil {
			out.RawString("null")
		} else {
			(*in.StorageID).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"key\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Key))
	}
	{
		const prefix string = ",\"newValue\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.NewValue))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EventDomStorageItemAdded) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoDomstorage8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EventDomStorageItemAdded) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoDomstorage8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EventDomStorageItemAdded) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoDomstorage8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EventDomStorageItemAdded) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoDomstorage8(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoDomstorage9(in *jlexer.Lexer, out *EnableParams) {
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoDomstorage9(out *jwriter.Writer, in EnableParams) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EnableParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoDomstorage9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EnableParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoDomstorage9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EnableParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoDomstorage9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EnableParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoDomstorage9(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoDomstorage10(in *jlexer.Lexer, out *DisableParams) {
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoDomstorage10(out *jwriter.Writer, in DisableParams) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DisableParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoDomstorage10(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DisableParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoDomstorage10(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DisableParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoDomstorage10(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DisableParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoDomstorage10(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoDomstorage11(in *jlexer.Lexer, out *ClearParams) {
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
		case "storageId":
			if in.IsNull() {
				in.Skip()
				out.StorageID = nil
			} else {
				if out.StorageID == nil {
					out.StorageID = new(StorageID)
				}
				(*out.StorageID).UnmarshalEasyJSON(in)
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoDomstorage11(out *jwriter.Writer, in ClearParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"storageId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.StorageID == nil {
			out.RawString("null")
		} else {
			(*in.StorageID).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ClearParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoDomstorage11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ClearParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoDomstorage11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ClearParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoDomstorage11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ClearParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoDomstorage11(l, v)
}
