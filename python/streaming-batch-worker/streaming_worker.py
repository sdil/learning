from kafka import KafkaConsumer
from time import sleep

def process_message(email):
    print(f"sending email to {email}")
    #sleep(1) # pretend to do work
    print(f"updating user {email} status to Waiting Confirmation")
    #sleep(1) # pretend to do work
    print("finished processing message")

if __name__ == "__main__":

    print("starting streaming consumer app")

    consumer = KafkaConsumer(
        'user_signups',
        bootstrap_servers=["localhost:9092"],
        group_id="group1"
    )

    for message in consumer:
        process_message(message.value)

        if killer.kill_now:
            print("End of the program. I was killed gracefully :)")
            consumer.close()
            exit()


