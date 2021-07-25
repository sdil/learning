from json import dumps
from kafka import KafkaProducer

if __name__ == "__main__":
    producer = KafkaProducer(
        bootstrap_servers=["localhost:9092"],
        value_serializer=lambda x: dumps(x).encode("utf-8"),
    )

    for i in range(10):
        data = {'email': f"user{i}@gmail.com"}
        producer.send('user_signups', value=data)
