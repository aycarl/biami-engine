# Week 7: Zero-Trust Security & DevSecOps

## Tasks

- Migrate API keys to AWS Secrets Manager.
- Deploy an AWS WAF to protect against SQLi and DDoS on mobile webhooks.
- Deliverable: A hardened infrastructure with least-privilege IAM roles.

## Implementation

1. Inventory every API key currently stored in application config, CI variables, or deployment scripts, and classify each secret by service owner and runtime usage.
2. Create AWS Secrets Manager entries for each key, apply consistent naming conventions, and restrict access with least-privilege IAM policies for only the workloads that need them.
3. Update application and deployment configuration so services read secrets from AWS Secrets Manager at runtime instead of using hard-coded values or checked-in environment files.
4. Rotate the migrated API keys after cutover, verify that old credentials are revoked, and document the rollback procedure in case any downstream integration fails.
5. Deploy AWS WAF in front of the mobile webhook endpoints and enable managed rules that cover common SQL injection, bad input, and volumetric attack patterns.
6. Add rate-limiting and IP reputation rules for webhook traffic, then test expected requests and blocked attack simulations to confirm legitimate traffic still succeeds.
7. Review IAM roles attached to the application, CI/CD pipeline, and webhook infrastructure, remove unused permissions, and confirm access is limited to required AWS resources only.
8. Deliver the week by publishing the final architecture changes, IAM policy updates, WAF rule summary, and validation evidence showing the infrastructure is hardened and operating under least privilege.
