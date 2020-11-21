import zmq from 'k6/x/zmq';

const socket = zmq.newSocket("tcp://localhost:5555")

export default function () {
  zmq.send(socket,"foo")
}

export function teardown () {
  zmq.closeSocket(socket)
}