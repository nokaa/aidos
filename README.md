# Search
A server that makes it easy to search other websites. This is similar duckduckgo's bangs.

TODO(nokaa): Create a config file of some sort to make adding new bangs easier.

Current bangs:
- `!w` searchs English Wikipedia
- `!wt` searches English Wiktionary
- `!gh` searches GitHub
- `!g` searches Encrypted Google

Any other bang provided will be forwarded to duckduckgo, allowing users access to their vast collection.

### Build
Search requires the Go compiler to be installed on your system. It has only been tested with Go 1.6, but likely works with older versions as well.

To build:
```
git clone https://github.com/nokaa/search
cd search
go build
```

Now you can simply run the `search` binary that is created.
