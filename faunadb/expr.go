package faunadb

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

/*
Expr is the base type for FaunaDB query language expressions.

Expressions are created by using the query language functions in query.go. Query functions are designed to compose with each other, as well as with
custom data structures. For example:

	type User struct {
		Name string
	}

	_, _ := client.Query(
		Create(
			Collection("users"),
			Obj{"data": User{"John"}},
		),
	)

*/
type Expr interface {
	String() string
	expr() // Make sure only internal structures can be marked as valid expressions
}

type unescapedObj map[string]Expr
type unescapedArr []Expr
type invalidExpr struct{ err error }

func (obj unescapedObj) expr() {}
func (obj unescapedObj) String() string {
	if len(obj) == 1 && obj["object"] != nil {
		return fmt.Sprintf("%s", obj["object"])
	}
	strs := []string{}
	for k, v := range obj {
		strs = append(strs, fmt.Sprintf("%s: %s", strconv.Quote(k), v))
	}
	return fmt.Sprintf("Obj{%s}", strings.Join(strs, ", "))
}

func (arr unescapedArr) expr() {}
func (arr unescapedArr) String() string {
	strs := []string{}
	for _, v := range arr {
		strs = append(strs, fmt.Sprintf("%s", v))
	}
	return fmt.Sprintf("Arr{%s}", strings.Join(strs, ", "))
}

func (inv invalidExpr) expr()          {}
func (inv invalidExpr) String() string { return "invalidExpr" }

func (inv invalidExpr) MarshalJSON() ([]byte, error) {
	return nil, inv.err
}

// Obj is a expression shortcut to represent any valid JSON object
type Obj map[string]interface{}

func (obj Obj) expr() {}

func (obj Obj) String() string {
	if len(obj) == 1 && obj["object"] != nil {
		return fmt.Sprintf("%s", obj["object"])
	}
	strs := []string{}
	for k, v := range obj {
		strs = append(strs, fmt.Sprintf("%s: %s", k, v))
	}
	return fmt.Sprintf("Obj{%s}", strings.Join(strs, ", "))
}

// Arr is a expression shortcut to represent any valid JSON array
type Arr []interface{}

func (arr Arr) expr() {}

func (arr Arr) String() string {
	strs := []string{}
	for _, v := range arr {
		strs = append(strs, fmt.Sprintf("%s", v))
	}
	return fmt.Sprintf("Arr{%s}", strings.Join(strs, ", "))
}

// MarshalJSON implements json.Marshaler for Obj expression
func (obj Obj) MarshalJSON() ([]byte, error) { return json.Marshal(wrap(obj)) }

// MarshalJSON implements json.Marshaler for Arr expression
func (arr Arr) MarshalJSON() ([]byte, error) { return json.Marshal(wrap(arr)) }

// OptionalParameter describes optional parameters for query language functions
type OptionalParameter func(optionalKeysMapping)

type optionalKeysMapping map[string]*Expr

func applyOptionals(mappings optionalKeysMapping, options []OptionalParameter) {
	for _, option := range options {
		option(mappings)
	}
}
