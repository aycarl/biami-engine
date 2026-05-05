# Week 10: Full CI/CD & Final Architecture

## Tasks

- Build a GitHub Actions pipeline with Nessus or Snyk scans.
- Finalize a System Architecture Diagram for the portfolio.
- Deliverable: A fully automated "Zero-Downtime" deployment workflow.

## Implementation

- Create a GitHub Actions workflow that runs on every push and pull request.
- Add build, test, and deployment stages so changes are validated before release.
- Integrate either Nessus or Snyk into the pipeline to scan dependencies and fail the workflow on critical findings.
- Finalize the System Architecture Diagram to reflect the deployed services, CI/CD flow, security scans, and rollback path.
- Configure deployment to use a zero-downtime strategy, such as rolling or blue-green deployment, so traffic is not interrupted during releases.
- Validate the workflow end to end by triggering a deployment and confirming the application remains available throughout the release process.
