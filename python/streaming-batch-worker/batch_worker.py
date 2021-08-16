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

    temp_emails = []

    try:
        while True:
            temp_emails.append(emails.get_nowait())

    except Empty:
        pass

    # Combine all emails in 1 call
    # This is the beauty of batch worker
    print(f"sending email to user " + str(temp_emails))
    sleep(0.5)  # pretend to do work
    print(f"updating status to Waiting Confirmation for users " + str(temp_emails))
    sleep(0.5)  # pretend to do work
    print("finished processing messages")


def exit_gracefully(*args, **kwargs):
    global is_shutting_down
    is_shutting_down = True
    process_messages()
    exit()


if __name__ == "__main__":

    signal.signal(signal.SIGINT, exit_gracefully)
    signal.signal(signal.SIGTERM, exit_gracefully)

    print("starting batch consumer worker")

    consumer = Consumer()
    consumer.daemon = True
    consumer.start()

    while True:
        process_messages()
        sleep(5)
