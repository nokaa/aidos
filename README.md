# Aidos
A server that makes it easy to search other websites. This is similar duckduckgo's bangs.

Bangs are stored in a file database `aidos.db` which must be in the same directory that you run `aidos` from.

A command line utility to assist in creating and adding bangs to this database can be found in the `add_bang` directory.

Any bang not in the database will be forwarded to duckduckgo, allowing users access to their vast collection.

### Build
Search requires the Go compiler to be installed on your system. It has only been tested with Go 1.6, but likely works with older versions as well.

To build:
```
git clone https://github.com/nokaa/aidos
cd aidos
go build
```

Now you can simply run the `aidos` binary that is created.
