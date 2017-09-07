# Bufio Package with Timouts

In some situations a blocking reader might be problematic to work with. This package provides 
wrappers for bufio classes with timeout functionality.

### Use Case

This package becomes interesting when e.g. talking to hardware or another service in a serial manner
that works on a "send request -- get answer" basis.

One would make sure to synchronize send requests and wait for the answer before sending the next
request. If for some reason the answer does not look like we expect it to look like or there is no answer
at all one would be stuck on a read -- forever.

### Handling timeout

If a timeout occurs a connection might still be open and continue to bring in data. Whatever 
function timed out will still be waiting to receive new data. The next data coming in will
end up in nirvana.

For this reason a connection underlying a reader that timed out should be reconnected.

Any suggestions on improving this behaviour are very welcome!
