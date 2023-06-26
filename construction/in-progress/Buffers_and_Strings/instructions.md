1. Open the attached `movements.log` file. It is a rudimentary record of financial transacations.
2. Create:
* Some method of keeping track of the total `balance`, which will be in integers.
* File `outgoing.log` and `incoming.log`
3. Using `bufio` from Go's standard library, scan each line from this file. For each line:
* If the line is marked OUT, copy its contents to file `outgoing.log` and subtract its value from `balance`.
* If the line is marked IN, copy its contents to file `incoming.log` and add its value to `balance`.
* If the line is in error, copy the unaltered line to `inError.log`. Line is in error if:
    * One of the three fields is missing.
    * Second field is neither `OUT` nor `IN`.
    * The third field cannot be parsed to an integer.
    * The third field is lower than zero.
4. Print `balance` and linecount to stdout.