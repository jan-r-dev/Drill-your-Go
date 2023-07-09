1. Create a function that does the following (for all of the tasks, print simple text to stdout to indicate what is happening):
* Accepts a channel. You can name it `quit`.
* Uses a `select` statement inside of an infinite loop. This select statement should check if the `quit` channel has received a signal.
* If no signal is received, continue looping.
* If signal has been received, run a cleanup function and exit the loop. Acknowledgement should be sent across the `quit` channel from the function as the last action.
2. Run this function in a separate go routine.
3. After a set amount of time, send a signal on the `quit` channel to trigger cleanup and exit of the goroutine.
4. Wait until you receive the return signal from the goroutine and then exit.
