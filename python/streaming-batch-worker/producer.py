from kafka import KafkaProducer, KafkaAdminClient
from kafka.admin import NewTopic
from time import sleep

if __name__ == "__main__":

    topic_name = "user_signups"
    producer = KafkaProducer(bootstrap_servers=["localhost:9092"])

    try:
        # Create Kafka topic
        topic = NewTopic(name=topic_name,
                             num_partitions=1,
                             replication_factor=1)
        admin = KafkaAdminClient(bootstrap_servers="localhost:9092")
        admin.create_topics([topic])
    except Exception:
        print(f"Topic {topic_name} is already created")

    for i in range(10):
        email = f"user{i}@gmail.com"
        producer.send(topic_name, email.encode())
        sleep(0.1)
        print(f"Published message to message broker. User email: {email}")

