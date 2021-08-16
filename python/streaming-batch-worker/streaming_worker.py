from kafka import KafkaConsumer
from time import sleep
import signal

is_shutting_down = False


def process_message(email):
    print(f"sending email to {email}")
    sleep(0.5)  # pretend to do work
    print(f"updating user {email} status to Waiting Confirmation")
    sleep(0.5)  # pretend to do work
    print("finished processing message")


def graceful_exit(*args, **kwargs):
    global is_shutting_down
    is_shutting_down = True


if __name__ == "__main__":

    signal.signal(signal.SIGINT, graceful_exit)
    signal.signal(signal.SIGTERM, graceful_exit)

    print("starting streaming consumer worker")

    consumer = KafkaConsumer(
        "user_signups", bootstrap_servers=["localhost:9092"], group_id="group1"
    )

    for message in consumer:
        process_message(message.value)

        if is_shutting_down:
            break

    print("End of the program. I was killed gracefully")
    consumer.close()
    exit()
