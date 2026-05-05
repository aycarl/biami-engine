# Week 3: Networking with Terraform

## Tasks

- Script a Multi-AZ VPC with public/private subnets and NAT Gateways.
- Define Security Groups for strict traffic flow.
- Deliverable: A version-controlled network environment.

## Implementation

1. Create a Terraform configuration for a VPC spanning at least two Availability Zones.
2. Define public and private subnets in each AZ, and associate public subnets with a route table that uses an Internet Gateway.
3. Provision a NAT Gateway in each public subnet and route private subnet outbound traffic through the NAT Gateway in the same AZ.
4. Add Security Groups that allow only required traffic between tiers, such as SSH from approved admin sources and application traffic from trusted subnets.
5. Validate the plan with `terraform fmt`, `terraform validate`, and `terraform plan`, then commit the version-controlled network configuration as the weekly deliverable.
