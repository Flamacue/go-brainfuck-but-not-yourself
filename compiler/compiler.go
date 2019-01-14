package compiler

type InsType byte

const (
	ADD        InsType = '+'
	SUB        InsType = '-'
	RIGHT      InsType = '>'
	LEFT       InsType = '<'
	PUT        InsType = '.'
	READ       InsType = ','
	LOOP_BEGIN InsType = '['
	LOOP_END   InsType = ']'
)

type Instruction struct {
	Type InsType
	Arg  int // How many of the Type e.g. ++++ would be 1 Instruction with Type ADD with an Arg of 4
}

type Compiler struct {
	input    string
	length   int
	position int

	instructions []*Instruction
}

func New(input string) *Compiler {
	return &Compiler{
		input:        input,
		length:       len(input),
		instructions: []*Instruction{},
	}

}

func (c *Compiler) Compile() []*Instruction {
	loopStack := []int{}
	for c.position < c.length {
		current := c.input[c.position]

		switch current {
		case '+':
			c.compileFoldable('+', ADD)
		case '-':
			c.compileFoldable('-', SUB)
		case '>':
			c.compileFoldable('>', RIGHT)
		case '<':
			c.compileFoldable('<', LEFT)
		case '.':
			c.compileFoldable('.', PUT)
		case ',':
			c.compileFoldable(',', READ)
		case '[':
			pos := c.newInstruction(LOOP_BEGIN, 0)
			loopStack = append(loopStack, pos)
		case ']':
			beginLoopIdx := loopStack[len(loopStack)-1]
			loopStack = loopStack[:len(loopStack)-1]

			endLoopIdx := c.newInstruction(LOOP_END, beginLoopIdx)
			c.instructions[beginLoopIdx].Arg = endLoopIdx

		}

		c.position++
	}

	return c.instructions
}

func (c *Compiler) compileFoldable(char byte, insType InsType) {
	count := 1
	for c.position < c.length-1 && c.input[c.position+1] == char {
		count++
		c.position++
	}
	c.newInstruction(insType, count)
}

func (c *Compiler) newInstruction(insType InsType, arg int) int {
	ins := &Instruction{Type: insType, Arg: arg}
	c.instructions = append(c.instructions, ins)
	return len(c.instructions) - 1
}
