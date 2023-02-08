# Checks Documentation

## Cognito

### AWS_COG_001: Cognito allows unauthenticated users

Amazon Cognito has authenticated and unauthenticated mode to generate AWS temporary credentials for users. If unauthenticated access is enable, anyone can obtain AWS credentials by using a specific API call. If the AWS Credentials obtained have liberal AWS permissions, it might be possible for an unauthenticated user to access sensitive AWS services.

References:
- https://blog.appsecco.com/exploiting-weak-configurations-in-amazon-cognito-in-aws-471ce761963#03ec
- https://docs.aws.amazon.com/cognito/latest/developerguide/getting-credentials.html

## Lambda

### AWS_LMD_001: Lambdas are private

By default, Lambda runs functions in a secure VPC with access to AWS services and the internet. Lambda owns this VPC, which isn't connected to the account's default VPC. Internet access from a private subnet requires Network Address Translation (NAT).
To give your function access to the internet, route outbound traffic to a NAT gateway in a public subnet.

## IAM

### AWS_IAM_003: IAM User can't elevate rights

In AWS IAM, some permissions specific combinations could lead to compromise. These permissions could allowing an non-admin user to elevate its privileges to become admin on the AWS account.

References:
- https://github.com/RhinoSecurityLabs/AWS-IAM-Privilege-Escalation
- https://hackingthe.cloud/aws/exploitation/iam_privilege_escalation/

### AWS_IAM_004: IAM Role can't elevate rights

In AWS IAM, some permissions specific combinations could lead to compromise. These permissions could allowing an non-admin role to elevate its privileges to become admin on the AWS account.

References:
- https://github.com/RhinoSecurityLabs/AWS-IAM-Privilege-Escalation
- https://hackingthe.cloud/aws/exploitation/iam_privilege_escalation/

## S3

### AWS_S3_002: S3 buckets are not global but in one zone

All S3 buckets should to be in only one region.

### AWS_S3_005: S3 bucket have public access block enabled

When you create an S3 bucket, it is good practice to enable public access block to ensure the bucket is never accidentally public.
