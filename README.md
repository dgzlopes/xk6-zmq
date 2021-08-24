# xk6-zmq

This is a [k6](https://go.k6.io/k6) extension using the [xk6](https://github.com/grafana/xk6) system.

| :exclamation: This is a proof of concept, isn't supported by the k6 team, and may break in the future. USE AT YOUR OWN RISK! |
|------|

## Build

To build a `k6` binary with this extension, first ensure you have the prerequisites:

- [Go toolchain](https://go101.org/article/go-toolchain.html)
- Git

Then:

1. Install `xk6`:
  ```shell
  $ go install go.k6.io/xk6/cmd/xk6@latest
  ```

2. Build the binary:
  ```shell
  $ xk6 build --with github.com/dgzlopes/xk6-zmq@latest
  ```

## Example

```javascript
import zmq from 'k6/x/zmq';

const socket = zmq.newSocket("tcp://localhost:5555")

export default function () {
  url.Send(socket,"foo")
}

export function teardown () {
  url.closeSocket(socket)
}
```