# Week 9: The Observability Stack

## Tasks

- Create CloudWatch/Grafana dashboards for "Response Latency" and "Payment Success".
- Set up automated alerts for 5XX errors on USSD webhooks.
- Deliverable: Real-time system health monitoring.

## Implementation

- Define the source metrics for each dashboard panel:
  - **Response Latency**: p50/p95 response time for USSD and payment endpoints.
  - **Payment Success**: successful payments vs failed payments, shown as both count and rate.
- Create CloudWatch metrics and ensure the application emits the required dimensions (service, environment, endpoint, region).
- Build Grafana dashboards backed by CloudWatch data sources, with separate panels for live status and 24-hour trends.
- Configure alert rules for USSD webhook 5XX responses, including threshold, evaluation window, and notification channel.
- Test the dashboards and alerts by generating sample traffic and verifying that failures trigger notifications and appear in the charts.
- Document dashboard links, alert recipients, and the runbook steps for acknowledging and investigating incidents.
