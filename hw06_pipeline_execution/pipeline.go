package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := in

	for _, stage := range stages {
		currentOut := make(Bi)
		input := stage(out)
		go func() {
			defer close(currentOut)
			for {
				select {
				case <-done:
					return
				case val, ok := <-input:
					if !ok {
						return
					}
					currentOut <- val
				}
			}
		}()
		out = currentOut
	}

	return out
}
