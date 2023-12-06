package util

import (
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
		//fmt.Println("loop", string(p.arr[i]))
		if p.state == valueState {
			//fmt.Println("next", string(p.arr[i]))
			if i+1 == p.len {
				//fmt.Println("last", string(p.arr[i]))
				p.state = doneState
				return p.arr[s:i]
			}
			p.isSepSet(p.arr[i])
			if p.state == valueState {
				//fmt.Println("more", string(p.arr[i]))
				continue
			}
			//fmt.Println("value", string(p.arr[i]))
			p.curr = i + 1
			return p.arr[s:i]
		}
		if p.state != valueState {
			p.isSepSet(p.arr[i])
			if p.state == valueState {
				s = i
				//fmt.Println("start", string(p.arr[i]))
			}

			p.state = valueState
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

func (ts TokenSlice) ToString() string {
	if len(ts) == 0 {
		return ""
	}
	b := []byte{'['}
	del := []byte{' ', ','}
	for i := 0; i < len(ts); i++ {
		b = append(b, ts[i]...)
		b = append(b, del...)
	}
	b[len(b)-2] = ']'
	return string(b[:len(b)-1])
}
