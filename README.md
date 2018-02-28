# Mock Style Log Data Writer for Honeycomb.io

## Purpose:

This repo is written to be used to write data to Honeycomb.io for testing & experimentation purposes.

## How-to:

Setup the following with your keys & pertinent information. See the [Go lib] and check out the initialization section for more specifics. The specific code snippet that needs the write key looks like this.

### Step 1

```go
libhoney.Init(libhoney.Config{
  WriteKey: "77f502feb3ccfcdaa5b792ff49c872b0",
  Dataset: "honeycomb-golang-example",
})
defer libhoney.Close() // Flush any pending calls to Honeycomb
```

**NOTE:** The code is setup however to use an environment variable I've called "THRASHER_HONEYCOMBKEY". If the environment variable is set it will use that key, preventing the need to place the key inside the repository. However if you'd like, you could always just change the WriteKey out to your write key, but just be sure not to commit it back into your repository.

### Step 2

To change the number of threads writing to Honeycomb you can change the Go Routines started in this snippet of code. Just change the count, which is 10 in this example, and that many thread will kick off via Go Routines.

```go
for i := 1; i < goRoutineCount; i++ {
    go loadHoneycombData(perRoutineCount, "load_"+strconv.Itoa(i))
}
```

### Step 3

Run `go build` and then run the executable `data-writer-honeycombio`.
