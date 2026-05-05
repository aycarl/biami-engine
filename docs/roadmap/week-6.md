# Week 6: Event-Driven Serverless Logic

## Tasks

- Offload Quality Scoring to AWS Lambda triggered by Amazon SQS.
- Move scoring logic out of the core API to save compute costs.
- Deliverable: An asynchronous, decoupled scoring service.

## Implementation

- Create an Amazon SQS queue dedicated to quality-scoring jobs and define the message payload shape the API will publish.
- Update the core API so that new scoring requests are enqueued to SQS instead of being processed synchronously in the request path.
- Implement an AWS Lambda consumer that is triggered by SQS, validates each message, and runs the extracted scoring logic.
- Persist or publish scoring results from Lambda and add retry/dead-letter handling so failed messages can be inspected without blocking the API.
- Verify the end-to-end flow in a non-production environment, including queue delivery, Lambda execution, error handling, and expected latency improvements.
