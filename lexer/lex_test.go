//Need to be written in the other computer
//tests the stack

package lexer

import (
	"fmt"
	"testing"
)

func TestPush(t *testing.T) {
	stack := Stack{"hello", &Stack{}}
	push(&stack, "hi")
	if peek(&stack) != "hi" {
		t.Error("Expected 'hi', got ", peek(&stack))
	}
}

func TestPop(t *testing.T) {
	stack := Stack{"hello", &Stack{"", nil}}
	pop(&stack)
	if peek(&stack) != "" {
		t.Error("Expected true", peek(&stack))
	}
	push(&stack, "hi")
	push(&stack, "hoola")
	pop(&stack)

	if peek(&stack) != "hi" {
		t.Error("Expected hi", peek(&stack))
	}
}
