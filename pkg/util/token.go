package util

import (
	"fmt"
	"slices"
)

type state uint8

const (
	startState state = iota
	valueState
	valueEndState
	rowEndState
	doneState
)

type TokenParser struct {
	state     state
	nextToken func(*TokenParser) Token
	nextRow   func(*TokenParser) TokenSlice
	currRow   TokenSlice
	isSep     func(*TokenParser, byte)
	curr      uint64
	sep       []byte
	rowSep    []byte
	len       uint64
	arr       []byte
}

func NewTokenParser(arr []byte) *TokenParser {
	return NewTokenParserWithSeparators(arr, []byte{' '}, []byte{'\n'})
}

func NewTokenParserWithSeparators(arr []byte, sep, rowSep []byte) *TokenParser {
	return &TokenParser{
		state:     startState,
		nextToken: (*TokenParser).stateValue,
		nextRow:   (*TokenParser).stateRow,
		isSep:     (*TokenParser).isSeparatorSetState,
		sep:       sep,
		rowSep:    rowSep,
		len:       uint64(len(arr)),
		arr:       arr,
	}
}

func (p *TokenParser) Reset(arr []byte) {
	p.state = startState
	p.currRow = p.currRow[:0]
	p.curr = 0
	p.len = uint64(len(arr))
	p.arr = arr
}

func (p *TokenParser) Next() Token {
	return p.nextToken(p)
}

func (p *TokenParser) NextRow() TokenSlice {
	return p.nextRow(p)
}

func (p *TokenParser) More() bool {
	return p.state != doneState
}

func (p *TokenParser) stateValue() Token {
	var s uint64

	for i := p.curr; i < p.len; i++ {
		if p.state == valueState {
			if i+1 == p.len {
				p.curr = p.len
				p.state = doneState
				return p.arr[s:p.len]
			}
			p.isSepSet(p.arr[i])
			if p.state == valueState {
				continue
			}
			p.curr = i + 1
			return p.arr[s:i]
		}
		if p.state != valueState {
			p.isSepSet(p.arr[i])
			if p.state == valueState {
				s = i
			}
		}
	}
	p.state = doneState
	return nil
}

func (p *TokenParser) stateRow() TokenSlice {
	p.currRow = p.currRow[:0]
	if p.state == doneState {
		//fmt.Println("done")
		return p.currRow
	}
	for {
		t := p.Next()
		p.currRow = append(p.currRow, t)
		if p.state == rowEndState || p.state == doneState {
			return p.currRow
		}
	}
}

func (p *TokenParser) isSepSet(b byte) {
	p.isSep(p, b)
}

func (p *TokenParser) isSeparatorSetState(b byte) {
	if slices.Contains(p.sep, b) {
		p.state = valueEndState
		return
	}
	if slices.Contains(p.rowSep, b) {
		p.state = rowEndState
		return
	}
	p.state = valueState
	return
}

type Token []byte

func (t Token) ParseInt64() int64 {
	return ParseInt64NoError(t)
}

func (t Token) ParseUInt64() uint64 {
	return ParseUint64NoError(t)
}

type TokenSlice []Token

func (ts TokenSlice) ToUint64() []uint64 {
	res := make([]uint64, len(ts))
	return ts.ConvertToUint64(res)
}

func (ts TokenSlice) ConvertToUint64(res []uint64) []uint64 {
	if len(res) < len(ts) {
		res = make([]uint64, len(ts))
		fmt.Println("growing res")
	}
	for i := 0; i < len(ts); i++ {
		res[i] = ts[i].ParseUInt64()
	}
	return res
}

func (ts TokenSlice) ToInt64() []int64 {
	res := make([]int64, len(ts))
	return ts.ConvertToInt64(res)
}

func (ts TokenSlice) ConvertToInt64(res []int64) []int64 {
	if len(res) < len(ts) {
		res = make([]int64, len(ts))
		fmt.Println("growing res")
	}
	for i := 0; i < len(ts); i++ {
		res[i] = ts[i].ParseInt64()
	}
	return res
}

func (ts TokenSlice) ToString() string {
	if len(ts) == 0 {
		return ""
	}
	b := []byte{'['}
	del := []byte{',', ' '}
	for i := 0; i < len(ts); i++ {
		b = append(b, ts[i]...)
		b = append(b, del...)
	}
	b[len(b)-2] = ']'
	return string(b[:len(b)-1])
}
