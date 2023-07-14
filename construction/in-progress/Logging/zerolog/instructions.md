1. Get `zerolog` from https://github.com/rs/zerolog. You'll need to use Go modules. (https://pkg.go.dev/github.com/rs/zerolog#section-readme)
2. Set up a command line flag `severity` which will allow you to choose between log levels `debug`, `info`, and `error`
3. Set up TimeFieldFormat for `zerolog` to use the unix timestamp.
4. Depending on severity chosen, log the following:
* `error` should only log one error. Text is up to you, but **make sure to include the stacktrace.**
* `info` should add one more log with a custom string field. Text is up to you.
* `debug` should add two more logs. Add a custom int and string field to them.
5. Instead of logging to stdout, log instead to a log file `./logs/output.log`.