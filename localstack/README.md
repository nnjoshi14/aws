##Installation
pip install localstack
localstack start
SERVICES=kinesis,lambda,sns,sqs,dynamodb localstack start
