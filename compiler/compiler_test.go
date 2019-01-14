package compiler

import (
	"testing"
)

func TestCompiler(t *testing.T) {
	input := `aa++><--.,testingdawg[->+<]as  d`

	c := New(input)
	c.Compile()

	tests := []struct {
		expectedType InsType
		expectedArg  int
	}{
		{ADD, 2},
		{RIGHT, 1},
		{LEFT, 1},
		{SUB, 2},
		{PUT, 1},
		{READ, 1},
		{LOOP_BEGIN, 11},
		{SUB, 1},
		{RIGHT, 1},
		{ADD, 1},
		{LEFT, 1},
		{LOOP_END, 6},
	}

	for i, tt := range tests {
		if c.instructions[i].Type != tt.expectedType {
			t.Fatalf("tests[%d] - InstructionType wrong. expected=%q, got=%q", i, tt.expectedType, c.instructions[i].Type)
		}

		if c.instructions[i].Arg != tt.expectedArg {
			t.Fatalf("tests[%d] - Arg wrong. expected=%d, got=%d", i, tt.expectedArg, c.instructions[i].Arg)
		}

	}
}
