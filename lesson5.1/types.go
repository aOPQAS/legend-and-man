package main

import "encoding/xml"

type Recipes struct {
	XMLName xml.Name `xml:"recipes" json:"-"`
	Cakes   []Cake   `xml:"cake" json:"cake"`
}

type Cake struct {
	Name        string       `xml:"name" json:"name"`
	Stovetime   string       `xml:"stovetime" json:"time"`
	Ingredients []Ingredient `xml:"ingredients>item" json:"ingredients"`
}

type Ingredient struct {
	Itemname  string `xml:"itemname" json:"ingredient_name"`
	Itemcount string `xml:"itemcount" json:"ingredient_count"`
	Itemunit  string `xml:"itemunit" json:"ingredient_unit,omitempty"`
}
