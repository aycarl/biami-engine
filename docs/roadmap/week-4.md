# Week 4: Managed Data & Containerization

## Tasks

- Provision RDS PostgreSQL and ElastiCache Redis via Terraform.
- Create multi-stage Dockerfiles for the Go API and Celery/Go workers.
- Deliverable: Optimized Docker images and secure managed databases.

## Implementation

- Define Terraform modules/resources for RDS PostgreSQL and ElastiCache Redis, including subnet groups, security groups, backups, and encryption settings.
- Store database and cache connection details in environment-specific configuration or secret management, and update the application to read these values at runtime.
- Create multi-stage Dockerfiles that compile/build artifacts in a builder stage and copy only the minimal runtime assets into the final image.
- Configure the Go API and Celery/Go workers to connect to the managed PostgreSQL and Redis instances using least-privilege credentials.
- Validate the setup by running Terraform plan/apply in the target environment, building the images, and confirming database/cache connectivity from the containers.
- Document image size, required environment variables, and rollback steps as the week-4 deliverable.
