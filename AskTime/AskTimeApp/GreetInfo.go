//Package AskTimeApp comment
// This file war generated by tars2go 1.1
// Generated from HttpAsk.tars
package AskTimeApp

import (
	"fmt"
	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
)

//GreetInfo strcut implement
type GreetInfo struct {
	Greeting     string  `json:"greeting"`
	FeedBackInfo SayInfo `json:"feedBackInfo"`
	Time         string  `json:"time"`
}

func (st *GreetInfo) resetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *GreetInfo) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.resetDefault()

	err = _is.Read_string(&st.Greeting, 0, true)
	if err != nil {
		return err
	}

	err = st.FeedBackInfo.ReadBlock(_is, 1, false)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Time, 2, false)
	if err != nil {
		return err
	}

	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *GreetInfo) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.resetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require GreetInfo, but not exist. tag %d", tag)
		}
		return nil

	}

	st.ReadFrom(_is)

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *GreetInfo) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_string(st.Greeting, 0)
	if err != nil {
		return err
	}

	err = st.FeedBackInfo.WriteBlock(_os, 1)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Time, 2)
	if err != nil {
		return err
	}

	return nil
}

//WriteBlock encode struct
func (st *GreetInfo) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	st.WriteTo(_os)

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}
