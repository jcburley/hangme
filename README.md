# HangMe -- Go 1.14.x Program That Intermittently Hangs on Mac OS X (Darwin)

After doing `go build` here, do something like this:

```
{ grep '^\s*_' hangme.go; while true; do timeout -s QUIT 30s ./hangme 1000; date; done } 2>&1
```

That'll repeatedly invoke `hangme`, telling it to perform 1,000 iterations of running itself (passing `0` as the only argument, meaning do nothing and exit immediately), under the auspices
of a `timeout` invocation that'll cause a `SIGQUIT` to be raised if `hangme` takes longer than 30 seconds to perform the task. `SIGQUIT` causes the Go runtime to output a stack dump to
`stderr`.

If this doesn't `SIGQUIT` after (say) 10 minutes, perhaps the bug is not being triggered or has been fixed.

Commenting out the `import` of `plugin` seems to work around the problem. Perhaps there's some strange interaction between whatever (pre-`main()`) initialization code is run by `plugin` and the
"new" timer code switched on around 2019-04-11 via commit `6becb033341602f2df9d7c55cc23e64b925bbee2` in the Golang source tree and/or the `os/exec` `Command.Start()` receiver.
