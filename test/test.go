package test

import (
	"cat/console/color"
	"encoding/json"
	"fmt"
	"testing"
)

type Asserter struct {
	t        *testing.T
	testLog  string
	cond1    string
	cond2    string
	testName string
	log      bool
}
type Logged struct {
	Condition1 string `json:"condition1"`
	Condition2 string `json:"condition2"`
	TestName   string `json:"testName"`
}

func New(t *testing.T) *Asserter {
	return &Asserter{t: t}
}

func (a *Asserter) That(cond1 string) *Asserter {
	a.cond1 = cond1
	return a
}

func (a *Asserter) ToBe(cond2 string) {
	a.cond2 = cond2
	if a.cond1 != a.cond2 {
		if a.log {
			log := &Logged{
				Condition1: a.cond1,
				Condition2: a.cond2,
				TestName:   a.testName,
			}
			l, err := json.Marshal(log)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(string(l))
		}
		expect := color.Colorize(fmt.Sprintf(" expected=%q,", a.cond1), color.Green)
		got := color.Colorize(fmt.Sprintf(" got=%q", a.cond2), color.Red)
		log := a.testLog + a.testName + expect + got
		a.t.Fatalf(log)
	}
}

func (a *Asserter) WithName(name string) *Asserter {
	a.testName = color.Colorize(name, color.Cyan)
	return a
}

func (a *Asserter) WithIndex(number int) *Asserter {
	a.testLog = color.Colorize(fmt.Sprintf("Tests[%d] - ", number), color.Yellow)
	return a
}

func (a *Asserter) WithLog() *Asserter {
	a.log = true
	return a
}
