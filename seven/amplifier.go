package main

func pipe(c1, c2 chan int) {
	go func() {
		for out := range c1 {
			c2 <- out
		}
	}()
}

func (c *Computer) runAmplifier(instructions []int, phase int) (chan int, chan int) {
	input := make(chan int)
	output := make(chan int)

	go func() {
		c.runProgram(instructions, input, output)
	}()

	input <- phase
	return input, output
}

func runNetwork(instructions []int, phases []int) int {
	var lastOut chan int

	for i, phase := range phases {
		c := new(Computer)
		in, out := c.runAmplifier(instructions, phase)

		switch i {
		case 0:
			in <- 0
		default:
			pipe(lastOut, in)
		}

		lastOut = out
	}

	log := make([]int, 0)
	for value := range lastOut {
		log = append(log, value)
	}

	return log[len(log)-1]
}

func runFeedbackNetwork(instructions []int, phases []int) int {
	var (
		firstIn chan int
		lastOut chan int
	)

	for i, phase := range phases {
		c := new(Computer)
		in, out := c.runAmplifier(instructions, phase)

		switch i {
		case 0:
			in <- 0
			firstIn = in
		default:
			pipe(lastOut, in)
		}

		lastOut = out
	}

	log := make([]int, 0)
	for value := range lastOut {
		log = append(log, value)
		firstIn <- value
	}

	return log[len(log)-1]
}
