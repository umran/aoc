package main

func pipe(c1, c2 chan int) {
	go func() {
		for out := range c1 {
			c2 <- out
		}
	}()
}

func runAmplifier(instructions []int, phase int) (chan int, chan int) {
	c := new(Computer)

	input, output := c.runProgram(instructions)

	input <- phase
	return input, output
}

func runNetwork(instructions []int, phases []int) int {
	var lastOut chan int

	for i, phase := range phases {
		in, out := runAmplifier(instructions, phase)

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
		in, out := runAmplifier(instructions, phase)

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
