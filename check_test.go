package check

import (
	"fmt"
	"log"
	"testing"
)

func MyFunc1() {
	log.Panic("MyFunc1")
}

func MyFunc2(do func()) {
	do()
}

func TestCase4(t *testing.T) {
	c1 := New(nil)
	MyFunc2(func() {
		err := fmt.Errorf("MyFunc2 Check")
		c1.CheckReport(err, "TestCase4", 0)
	})
}

func TestCase3(t *testing.T) {
	MyFunc2(func() {
		err := fmt.Errorf("MyFunc2 Check")
		CheckSkip(err, 2)
	})
}

func TestCase1(t *testing.T) {
	c1 := New(nil)
	MyFunc2(func() {
		err := fmt.Errorf("MyFunc2 Check")
		c1.Check(err)
	})

}

func TestCase2(t *testing.T) {
	MyFunc2(func() {
		err := fmt.Errorf("MyFunc2 Check")
		Check(err)
	})
}
