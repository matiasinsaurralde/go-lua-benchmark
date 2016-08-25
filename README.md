# go-lua-benchmark

Comparing the performance of LuaJIT embedded in Go.

## Handlers

Check [main.go](main.go) for details.

### NotCachedHandler

This HTTP handler will perform a standard `luaL_dofile` (file read operation) on every request, this is the slowest handler!

### CachedHandler

This handler will make use of `luaL_dostring` and the cached script data contained in `cached_script`. `cached_script` is set by `LuaCacheScript`.

### GoHandler

The fastest handler ever!

![](http://i.giphy.com/hM87DMnls5oZy.gif)

## Results

Using [boom](https://github.com/rakyll/boom) with the following parameters:

**-n 5000** (number of requests)

**-c 50** (number of concurrent requests)

On a Macbook Pro Mid 2014, Intel Core i7, 16 GB RAM.

### NotCachedHandler

```
Summary:
  Total:	0.2011 secs
  Slowest:	0.0234 secs
  Fastest:	0.0002 secs
  Average:	0.0019 secs
  Requests/sec:	24861.0858
  Total data:	75000 bytes
  Size/request:	15 bytes

Status code distribution:
  [200]	5000 responses

Response time histogram:
  0.000 [1]	|
  0.003 [4476]	|∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  0.005 [481]	|∎∎∎∎
  0.007 [21]	|
  0.009 [11]	|
  0.012 [0]	|
  0.014 [1]	|
  0.016 [4]	|
  0.019 [2]	|
  0.021 [2]	|
  0.023 [1]	|

Latency distribution:
  10% in 0.0010 secs
  25% in 0.0017 secs
  50% in 0.0019 secs
  75% in 0.0021 secs
  90% in 0.0025 secs
  95% in 0.0030 secs
  99% in 0.0046 secs
```

### CachedHandler

```
Summary:
  Total:	0.1778 secs
  Slowest:	0.0275 secs
  Fastest:	0.0001 secs
  Average:	0.0017 secs
  Requests/sec:	28123.5876
  Total data:	75000 bytes
  Size/request:	15 bytes

Status code distribution:
  [200]	5000 responses

Response time histogram:
  0.000 [1]	|
  0.003 [4551]	|∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  0.006 [365]	|∎∎∎
  0.008 [46]	|
  0.011 [8]	|
  0.014 [13]	|
  0.017 [10]	|
  0.019 [4]	|
  0.022 [1]	|
  0.025 [0]	|
  0.027 [1]	|

Latency distribution:
  10% in 0.0005 secs
  25% in 0.0009 secs
  50% in 0.0014 secs
  75% in 0.0021 secs
  90% in 0.0028 secs
  95% in 0.0036 secs
  99% in 0.0072 secs
```

### GoHandler

```
Summary:
  Total:	0.0802 secs
  Slowest:	0.0034 secs
  Fastest:	0.0001 secs
  Average:	0.0008 secs
  Requests/sec:	62368.6244
  Total data:	70000 bytes
  Size/request:	14 bytes

Status code distribution:
  [200]	5000 responses

Response time histogram:
  0.000 [1]	|
  0.000 [592]	|∎∎∎∎∎∎∎∎
  0.001 [2685]	|∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  0.001 [1004]	|∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  0.001 [352]	|∎∎∎∎∎
  0.002 [186]	|∎∎
  0.002 [102]	|∎
  0.002 [43]	|
  0.003 [19]	|
  0.003 [11]	|
  0.003 [5]	|

Latency distribution:
  10% in 0.0004 secs
  25% in 0.0005 secs
  50% in 0.0007 secs
  75% in 0.0009 secs
  90% in 0.0013 secs
  95% in 0.0016 secs
  99% in 0.0023 secs
```
