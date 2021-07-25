from time import sleep
from kafka import KafkaConsumer
import threading
from queue import Queue, Empty

items = Queue()


class Consumer(threading.Thread):
    def __init__(self):
        threading.Thread.__init__(self)

    def run(self):
        consumer = KafkaConsumer(
            "user_signups",
            bootstrap_servers=["localhost:9092"],
            auto_offset_reset="earliest",
            enable_auto_commit=True,
        )
        for message in consumer:
            self.process_message(message.value)

        consumer.close()

    def process_message(self, message):
        items.put(message)


def process_messages():
    try:
        while True:
            i = items.get_nowait()
            print(f"process user signup {i}")

    except Empty:
        pass


if __name__ == "__main__":
    print("Starting batching worker")

    consumer = Consumer()
    consumer.daemon = True
    consumer.start()

    while True:
        process_messages()
        sleep(5)
