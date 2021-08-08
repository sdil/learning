from time import sleep
from kafka import KafkaConsumer
import threading
from queue import Queue, Empty

emails = Queue()


class Consumer(threading.Thread):
    def __init__(self):
        threading.Thread.__init__(self)

    def run(self):
        consumer = KafkaConsumer(
            "user_signups", bootstrap_servers=["localhost:9092"], group_id="group1"
        )

        for message in consumer:
            self.insert_to_buffer(message.value)

    def insert_to_buffer(self, message):
        print("received a message, inserting into a queue buffer")
        emails.put(message)


def process_messages():
    print("processing message in queue buffer")

    try:
        while True:
            email = emails.get_nowait()
            print(f"sending email to user {email}")
            sleep(0.5)  # pretend to do work
            print(f"updating user {email} status to Waiting Confirmation")
            sleep(0.5)  # pretend to do work
            print("finished processing message")

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
