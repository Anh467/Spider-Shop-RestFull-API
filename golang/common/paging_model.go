package common

import "strconv"

const LIMIT_DEFAULT = 5
const OFFSET_DEFAULT = 0

type Paging struct {
	limit  int
	offset int
}

func (p *Paging) GetLimit() int {
	return p.limit
}

func (p *Paging) GetOffset() int {
	return p.offset
}

func (p *Paging) SetDefaultValue() {
	p.limit = LIMIT_DEFAULT
	p.offset = OFFSET_DEFAULT
}

func (p *Paging) Process(limitStr, offsetStr string) {
	var limit, offset int
	var err error
	if limit, err = strconv.Atoi(limitStr); err != nil {
		p.SetDefaultValue()
		return
	}
	if offset, err = strconv.Atoi(offsetStr); err != nil {
		p.SetDefaultValue()
		return
	}
	if limit < 0 || offset < 0 {
		p.SetDefaultValue()
		return
	}
	p.limit = limit
	p.offset = offset
}
