package graph

import (
)

type value interface {}

type Element struct {
	Value value
	Parent *Element
}

func (self *Element) Union(e2 Element) {
	self.Parent = e2.Parent
}

func (self Element) Find() Element {
	if *self.Parent == self {
		return self
	} else {
		self.Parent = self.Parent.Parent
		return self.Find()
	}
}