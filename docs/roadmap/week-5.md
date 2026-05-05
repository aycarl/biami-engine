# Week 5: Cloud Orchestration (ECS Fargate)

## Tasks

- Deploy containers to Amazon ECS (Fargate) for serverless scaling.
- Configure Application Load Balancers (ALB) and Auto-scaling.
- Deliverable: A highly available, auto-scaling cluster.

## Implementation

1. Build and tag the application container image, then push it to Amazon ECR so ECS can pull a versioned artifact.
2. Create an ECS task definition for Fargate that specifies CPU, memory, container port mappings, environment variables, and logging configuration.
3. Deploy an ECS service across multiple subnets and availability zones to ensure high availability.
4. Attach the service to an Application Load Balancer, configure a target group, and add health checks for the application endpoint.
5. Configure security groups and networking so the ALB accepts public traffic while ECS tasks only accept traffic from the ALB.
6. Enable ECS service auto-scaling based on CPU or memory utilization and verify that new tasks register cleanly behind the load balancer during scale-out events.
