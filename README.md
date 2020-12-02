# xk6-zmq

This is a [k6](https://github.com/loadimpact/k6) extension using the [xk6](https://github.com/k6io/xk6) system.

| :exclamation: This is a proof of concept, isn't supported by the k6 team, and may break in the future. USE AT YOUR OWN RISK! |
|------|

## Build

To build a `k6` binary with this extension, first ensure you have the prerequisites:

- [Go toolchain](https://go101.org/article/go-toolchain.html)
- Git

Then:

1. Clone `xk6`:
  ```shell
  git clone https://github.com/k6io/xk6.git
  cd xk6
  ```

2. Build the binary:
  ```shell
  CGO_ENABLED=1 go run ./cmd/xk6/main.go build master \
    --with github.com/dgzlopes/xk6-zmq
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