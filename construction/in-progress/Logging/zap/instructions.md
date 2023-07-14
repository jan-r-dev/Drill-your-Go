1. Get `zap` from https://github.com/uber-go/zap. You'll need to use Go modules.
2. Set up a command line flag `severity` which will allow you to choose between log levels `debug`, `info`, and `error`
3. Set up TimeFieldFormat for `zap` to use the unix timestamp.
4. Depending on severity chosen, log the following:
* `error` should only log one error. Text is up to you, but **make sure to include the stacktrace.**
* `info` should add one more log with a custom string field. Text is up to you.
* `debug` should add two more logs. Add a custom int and string field to them.
5. Instead of logging to stdout, log instead to a log file `./logs/output.log`.