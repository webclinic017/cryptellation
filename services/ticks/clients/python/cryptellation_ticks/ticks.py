import threading
import queue
import iso8601
import grpc

from cryptellation_ticks.config import Config
from cryptellation_ticks.tick import Tick

import cryptellation_ticks._genproto.ticks_pb2 as ticks
import cryptellation_ticks._genproto.ticks_pb2_grpc as ticks_grpc


class Ticks(threading.Thread):
    def __init__(self, exchange, pair):
        threading.Thread.__init__(self)
        self._queue = queue.Queue(maxsize=2)
        self._config = Config()
        self._channel = grpc.insecure_channel(self._config.url)
        self._stub = ticks_grpc.TicksServiceStub(self._channel)

        req = ticks.ListenSymbolRequest(exchange=exchange, pair_symbol=pair)
        self._generator = self._stub.ListenSymbol(req)
        self.start()

    def run(self):
        for tick in self._generator:
            self._queue.put(tick)

    def get(self) -> Tick:
        e = self._queue.get()
        return Tick(iso8601.parse_date(e.time), e.exchange, e.pair_symbol, e.price)


if __name__ == "__main__":
    t = Ticks("binance", "BTC-USDT")
    while True:
        print(t.get().price)
