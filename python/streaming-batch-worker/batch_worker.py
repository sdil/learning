from time import sleep
from kafka import KafkaConsumer
import threading
from queue import Queue, Empty
import signal

emails = Queue()
is_shutting_down = False

class Consumer(threading.Thread):
    def __init__(self):
        threading.Thread.__init__(self)

    def run(self):
        consumer = KafkaConsumer(
            "user_signups", bootstrap_servers=["localhost:9092"], group_id="group1"
        )

        for message in consumer:
            self.insert_to_buffer(message.value)

            if is_shutting_down:
                break

        consumer.close()

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


def exit_gracefully(*args, **kwargs):
    global is_shutting_down
    is_shutting_down = True
    process_messages()
    exit()


if __name__ == "__main__":

    signal.signal(signal.SIGINT, exit_gracefully)
    signal.signal(signal.SIGTERM, exit_gracefully)
    print("Starting batch worker")

    consumer = Consumer()
    consumer.daemon = True
    consumer.start()

    while True:
        process_messages()
        sleep(5)
