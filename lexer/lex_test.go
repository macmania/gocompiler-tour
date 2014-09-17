//Need to be written in the other computer
//tests the stack

package lexer

import (
	"testing"
)

func TestPush(t *testing.T) {
	stack := Stack{"hello", &Stack{}}
	stack = push(stack, "hi")
	if peek(stack) != "hi" {
		t.Error("Expected 'hi', got ", peek(stack))
	}
}

func TestPop(t *testing.T) {
	stack := Stack{"hello", &Stack{}}
	stack = pop(stack)
	if peek(stack) != "" {
		t.Error("Expected true", peek(stack))
	}
	stack = push(stack, "hi")
	stack = push(stack, "hoola")
	stack = pop(stack)
	if peek(stack) != "hi" {
		t.Error("Expected hi", peek(stack))
	}
}
