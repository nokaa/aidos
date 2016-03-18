This is a simple utility for adding new bangs to the database. You can easily modify this to work via the web,
I prefer the command line.

### Build
Search requires the Go compiler to be installed on your system. It has only been tested with Go 1.6, but likely works with older versions as well.

To build:
```
go build
```

Now you can simply run the `add_db` binary that is created.

I recommend moving this binary to the same directory as the aidos binary. This way you will not have to move the database file that is created.

You can safely modify the database while `aidos` is running.
