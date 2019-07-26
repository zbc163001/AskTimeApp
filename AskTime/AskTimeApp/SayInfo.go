//Package AskTimeApp comment
// This file war generated by tars2go 1.1
// Generated from HttpAsk.tars
package AskTimeApp

import (
	"fmt"
	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
)

//SayInfo strcut implement
type SayInfo struct {
	Name     string `json:"name"`
	Age      int32  `json:"age"`
	Address  string `json:"address"`
	Phonenum string `json:"phonenum"`
}

func (st *SayInfo) resetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *SayInfo) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.resetDefault()

	err = _is.Read_string(&st.Name, 0, true)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.Age, 1, false)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Address, 2, false)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Phonenum, 3, false)
	if err != nil {
		return err
	}

	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *SayInfo) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.resetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require SayInfo, but not exist. tag %d", tag)
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
func (st *SayInfo) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_string(st.Name, 0)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.Age, 1)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Address, 2)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Phonenum, 3)
	if err != nil {
		return err
	}

	return nil
}

//WriteBlock encode struct
func (st *SayInfo) WriteBlock(_os *codec.Buffer, tag byte) error {
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
