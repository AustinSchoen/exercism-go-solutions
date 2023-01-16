package robotname

import (
	"math/rand"
	"time"
	"strconv"
)

type Robot struct {
	Designation string
}

func (r *Robot) Reset() {
	r.SetDesignation()
}

func (r *Robot) Name() string {
	if r.Designation == "" {
		r.SetDesignation()
	}
	return r.Designation
}

func (r *Robot) SetDesignation() {
	r.Designation = randCapitalLetter()+randCapitalLetter()
	r.Designation = r.Designation + strconv.Itoa(rand3Digits())
}

func randCapitalLetter() string {
	rand.Seed(time.Now().UTC().UnixNano())
	i := 0

	for i < 65 || i > 90 {
		i = rand.Intn(100)
	}

	return string(i)
}

func rand3Digits() int {
	rand.Seed(time.Now().UTC().UnixNano())

	i := 0

	for i < 100 || i > 999 {
		i = rand.Intn(1000)
	}

	return i
}