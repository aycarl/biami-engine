# Week 8: Edge Computing & Global Delivery

## Tasks

- Host PWA assets on S3 and distribute via CloudFront.
- Implement Lambda@Edge for security header injection.
- Deliverable: High-performance static delivery for global mobile users.

## Implementation

1. Create an S3 bucket for the PWA build artifacts and configure it for static asset delivery.
2. Create a CloudFront distribution with the S3 bucket as the origin, enable compression, and set cache behaviors for versioned static assets.
3. Attach a Lambda@Edge function on the viewer response path to inject security headers such as `Content-Security-Policy`, `X-Content-Type-Options`, and `Strict-Transport-Security`.
4. Invalidate the CloudFront cache after each deployment so updated assets are served globally without stale content.
5. Verify the rollout by testing cache hits, response latency, and header presence from multiple geographic regions on mobile devices or throttled network profiles.
