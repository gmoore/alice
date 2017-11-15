Alice

My hacky take on [run-length encoding](https://en.wikipedia.org/wiki/Run-length_encoding). Inspired by this [old computer magazine](https://archive.org/stream/creativecomputing-1982-02/Creative_Computing_v08_n02_1982_February#page/n121/mode/2up).

# Building

You'll need Go and have your GOPATH set.

```
cd $GOPATH/src
git clone git@github.com:gmoore/alice.git
cd alice
make
```

# Running

```
./bin/drink --file littlegidding.txt
./bin/eat --file littlegidding.alice --out littlegidding-decompressed.txt
```

Note that most plain text files get compressed to a larger size than their uncompressed counterparts.
