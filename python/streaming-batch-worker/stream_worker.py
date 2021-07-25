from kafka import KafkaConsumer

def process_message(message):
    print(f"process user signup {message}")


if __name__ == "__main__":
    print("starting consumer app")
    consumer = KafkaConsumer(
        "user_signups",
        bootstrap_servers=["localhost:9092"],
        auto_offset_reset="earliest",
        enable_auto_commit=True,
    )

    for message in consumer:
        process_message(message.value)
