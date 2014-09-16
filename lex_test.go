//Need to be written in the other computer
//tests the stack

package main

import "testing"

func testPush(t *testing.T) {
	stack := Stack("hello", nil)
	push(stack, "hi")
	if peek(stack) != "hi" {
		t.Error("Expected 'hi', got ", peek(stack))
	}
}
