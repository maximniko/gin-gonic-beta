package models

import "strconv"

type Product struct {
	Id int
}

func (p *Product) String() string {
	return strconv.Itoa(p.Id)
}
