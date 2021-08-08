from time import sleep
from kafka import KafkaConsumer
import threading
from queue import Queue, Empty

users = Queue()


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
            self.insert_to_buffer(message.value)

        consumer.close()

    def insert_to_buffer(self, message):
        users.put(message)


def process_messages():
    try:
        while True:
            user = users.get_nowait()
            print(f"sending email to user {user}")
            sleep(2) # Pretend that we're doing IO work here
            print(f"updating database for user {user}")

    except Empty:
        pass


if __name__ == "__main__":
    print("Starting batch worker")

    consumer = Consumer()
    consumer.daemon = True
    consumer.start()

    while True:
        process_messages()
        sleep(5)
