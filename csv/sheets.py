import pika
import json
import csv

def process_response(ch, method, properties, body):
    response_data = body.decode('utf-8')
    print(f"Received response data: {response_data}")

    process_response_data(response_data)

    ch.basic_ack(delivery_tag=method.delivery_tag)

def process_response_data(response_data):

    data = json.loads(response_data)

    with open('responses.csv', 'a', newline='') as csvfile:
        writer = csv.writer(csvfile)

        if csvfile.tell() == 0:
            headers = ["form_id", "form_name", "form_desc", "question_id", "question", "answer_id", "answer"]
            writer.writerow(headers)

        form_id = data["form_id"]
        form_name = data["form_name"]
        form_desc = data["form_desc"]
        question_id = data["question_id"]
        question = data["question"]
        answer_id = data["answer_id"]
        answer = data["answer"]
        row = [form_id, form_name, form_desc, question_id, question, answer_id, answer]
        writer.writerow(row)

    print("Response added to CSV successfully!")


def main():
    connection = pika.BlockingConnection(pika.ConnectionParameters('localhost'))
    channel = connection.channel()

    queue_name = 'form_responses'
    channel.queue_declare(queue=queue_name)


    channel.basic_consume(queue=queue_name, on_message_callback=process_response)

    print("Waiting for messages...")
    channel.start_consuming()

if __name__ == "__main__":
    main()

