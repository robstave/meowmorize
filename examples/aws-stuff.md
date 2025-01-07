#flashcards #aws

status:: 
count:: 11
[[Flashcards]]

<!-- Card Start -->
### Front 
What is AWS CloudFormation and its core purpose?

### Back

AWS CloudFormation is an Infrastructure as Code (IaC) service that:
- **Automates** infrastructure deployment
- Creates resources in a **consistent, repeatable** manner
- Manages AWS resources through **templates**
- Handles **dependencies** automatically
- Treats infrastructure as **version-controlled code**
- Provides **rollback** capabilities
- Enables **stack** based resource management

<!-- Card End -->

<!-- Card Start -->
### Front

What is AWS CloudFormation and what are its main functions?
### Back

AWS CloudFormation is a service that:

- Allows users to model and provision their entire cloud infrastructure using configuration files
- Functions as an Infrastructure as Code (IaC) tool
- Enables creation, updating, and deletion of AWS resources in an automated way
- Uses JSON or YAML templates to define required AWS resources
- Manages resources as a single unit called a "stack"
- Supports provisioning across multiple AWS regions and accounts

<!-- Card End -->

<!-- Card Start -->
### Front

What are the main sections of a CloudFormation template?
### Back
A CloudFormation template consists of:

1. **Format Version**: Template version identifier
2. **Description**: Template documentation
3. **Parameters**: Input values
4. **Mappings**: Key-value lookups
5. **Conditions**: Conditional resource creation
6. **Resources** (Required): AWS resources to create
7. **Outputs**: Return values
8. **Metadata**: Template additional information
9. **Transform**: Include serverless transforms or macros

<!-- Card End -->
<!-- Card Start -->
### Front

How does CloudFormation handle stack updates?
### Back
CloudFormation stack updates involve:

- **Change Sets**: Preview of changes before execution
- **Update Behaviors**:
  - No interruption
  - Some interruption
  - Replacement
- **Rolling Updates**: Staged resource updates
- **Stack Policies**: Protect resources from updates
- **Rollback Triggers**: Automatic failure detection
- **Drift Detection**: Identify manual changes

<!-- Card End -->
<!-- Card Start -->
### Front 
Compare CloudFormation with Terraform

### Back
**CloudFormation**:
- Native AWS integration
- AWS-specific features
- JSON/YAML templates
- Stack-based management
- Free service (pay for resources)

**Terraform**:
- Multi-cloud support
- HCL syntax
- State file management
- Provider-based architecture
- Larger community modules
- More flexible state management

<!-- Card End -->
<!-- Card Start -->
### Front 

What are Nested Stacks and their benefits?
### Back
**Nested Stacks**:
- Reuse common template patterns
- Break down complex templates
- Maximum resources limit workaround
- Hierarchical stack management
- Modular infrastructure design

**Implementation**:
- Uses AWS::CloudFormation::Stack resource
- Supports parameters passing
- Enables template reuse
- Manages dependencies between stacks

<!-- Card End -->
<!-- Card Start -->
### Front 
What are key CloudFormation best practices?
### Back
1. **Template Structure**:
   - Use YAML for readability
   - Implement proper version control
   - Include comprehensive descriptions

2. **Security**:
   - Use stack policies
   - Implement least privilege
   - Encrypt sensitive parameters

3. **Operations**:
   - Test templates in dev environment
   - Use change sets
   - Implement proper tagging
   - Document dependencies

4. **Cost**:
   - Estimate costs before deployment
   - Use parameters for flexibility
   - Clean up unused resources

<!-- Card End -->
<!-- Card Start -->
### Front 
How to implement deletion protection in CloudFormation?

### Back
Multiple protection layers:

1. **Stack-level**:
   - EnableTerminationProtection flag
   - Stack policies

2. **Resource-level**:
   - DeletionPolicy attribute
   - RetentionPolicy settings
   - Termination protection

3. **Additional Measures**:
   - IAM policies
   - Resource lock settings
   - Service-specific protection

<!-- Card End -->
<!-- Card Start -->
### Front
What are CloudFormation Custom Resources?
### Back
Custom Resources allow:

- **Integration** with external resources
- **Custom provisioning** logic
- **Third-party service** management
- **Gap filling** for unsupported resources

**Implementation**:
- Lambda function backing
- SNS topic integration
- Response handling
- Lifecycle management
- Custom logic execution

<!-- Card End -->
<!-- Card Start -->
### Front 
How does CloudFormation Drift Detection work?

### Back
Drift Detection:

- **Purpose**: Identify unmanaged changes
- **Process**:
  1. Scan resource configuration
  2. Compare with template
  3. Report differences
  
- **Supported Operations**:
  - Stack level detection
  - Resource level detection
  - Drift status reporting
  
- **Limitations**:
  - Not all resources supported
  - Point-in-time detection
  - Manual initiation required

<!-- Card End -->
<!-- Card Start -->
### Front 

What are CloudFormation StackSets and their use cases?
### Back
**StackSets enable**:
- Deploy stacks across multiple accounts
- Regional deployment management
- Centralized operations
- Organizational unit deployment

**Features**:
- Concurrent deployment
- Automatic deployment ordering
- Failed operation handling
- Permission management
- Drift detection support
- Automatic retries

<!-- Card End -->
<!-- Card Start -->
### Front 
How to validate CloudFormation templates?
### Back
Validation methods:

1. **AWS Console**:
   - Template designer
   - Validate template option
   
2. **AWS CLI**:
   - validate-template command
   - describe-stack-events

3. **Best Practices**:
   - Linting tools
   - Static code analysis
   - Test deployments
   - Change set review
   - Parameter validation
   - Resource configuration checks

<!-- Card End -->

<!-- Card Start -->

### Front

What are the key components of an AWS CloudFormation template?

### Back

An AWS CloudFormation template typically includes:

1. Parameters: For user input and customization
2. Mappings: Key-value pairs for conditional resource definition
3. Conditions: For creating or modifying resources based on specific criteria
4. Resources: The only mandatory section, declaring all AWS resources
5. Outputs: For exporting information about created resources
6. Metadata: To organize parameters logically

<!-- Card End --> 
<!-- Card Start -->
### Front

What are some real-world use cases for AWS CloudFormation?

### Back

Real-world use cases for AWS CloudFormation include:

1. Infrastructure management with DevOps: Automating testing and deployment
2. Production stack scaling: From single EC2 instances to complex multi-region applications
3. Defining subnets and services: Easy provisioning of VPCs, ECS, or OpsWorks
4. Replicating environments: Creating identical dev, test, and production setups
5. Version control for infrastructure: Tracking changes and facilitating rollbacks
6. Multi-region deployments: Consistent resource provisioning across different regions
7. Access control: Managing permissions and security configurations
8. Application deployment: Automating the setup of entire application stacks

<!-- Card End -->



#flashcards #aws 
status:: 
count:: 27
[[Flashcards]]

<!-- Card Start -->

### Front  

Which of the following are benefits of migrating to the AWS Cloud? (Choose two.)  

- A. Operational resilience
- B. Discounts for products on Amazon.com
- C. Business agility
- D. Business excellence
- E. Increased staff retention

### Back

**A** and **C**

<!-- Card End --> 

<!-- Card Start -->

### Front  

What is AWS CloudFront?

### Back

**AWS CloudFront** is a **Content Delivery Network (CDN)** service that securely delivers data, videos, applications, and APIs to customers globally with low latency and high transfer speeds. It integrates with other AWS services to provide a seamless content distribution experience.

<!-- Card End -->
<!-- Card Start -->

### Front

What are the two types of CloudFront distributions?

### Back

1. **Web Distributions**: Used for serving static and dynamic content over HTTP and HTTPS.
2. **RTMP Distributions**: Used for streaming media files using the RTMP protocol. _(Note: RTMP distributions are being deprecated in favor of newer streaming technologies.)_

<!-- Card End -->
<!-- Card Start -->

### Front 

Define an Origin in AWS CloudFront.
### Back

An **Origin** is the source of the content that CloudFront will distribute. It can be an **AWS S3 bucket**, an **Elastic Load Balancer**, an **EC2 instance**, or any **custom HTTP server**.

<!-- Card End -->
<!-- Card Start -->
### Front 

What are Edge Locations in CloudFront?

### Back

**Edge Locations** are data centers located globally where CloudFront caches copies of your content. They are used to deliver content to end-users with low latency by serving requests from the nearest Edge Location.

<!-- Card End -->
<!-- Card Start -->
### Front 

Explain Regional Edge Caches in CloudFront.
### Back

**Regional Edge Caches** are larger caching layers located between the Edge Locations and the origin. They help reduce the load on the origin by serving content to multiple Edge Locations from a single Regional Edge Cache, improving cache hit ratios and reducing latency.

<!-- Card End -->
<!-- Card Start -->
### Front

What are Cache Behaviors in CloudFront?
### Back

**Cache Behaviors** define how CloudFront handles requests for specific paths within your distribution. They allow you to configure settings such as caching policies, origin selection, viewer protocols, and more for different parts of your website or application.

<!-- Card End -->
<!-- Card Start -->
### Front 

What is Lambda@Edge in CloudFront?
### Back

**Lambda@Edge** allows you to run **AWS Lambda functions** at CloudFront Edge Locations. This enables you to execute custom code to modify requests and responses, personalize content, perform authentication, and more, closer to the end-user for reduced latency.

<!-- Card End -->
<!-- Card Start -->
### Front  

How is AWS CloudFront priced?
### Back

CloudFront pricing is based on:

- **Data Transfer Out** to the internet.
- **Requests** made to your distribution.
- **Invalidation Requests** beyond the free tier.
- **Additional features** like Lambda@Edge and Field-Level Encryption.

Pricing varies by **region** and usage levels. There are no upfront fees, and you pay only for what you use.
<!-- Card End -->
<!-- Card Start -->
### Front 

List security features of AWS CloudFront.
### Back

- **HTTPS Support**: Secure content delivery.
- **AWS Shield Integration**: Protection against DDoS attacks.
- **AWS WAF Integration**: Web application firewall for filtering malicious traffic.
- **Origin Access Identity (OAI)**: Restricts access to S3 origins.
- **Signed URLs and Cookies**: Control access to content.

<!-- Card End --> <!-- Card Start -->

### Front 

What is an Origin Access Identity (OAI) in CloudFront?

### Back

An **Origin Access Identity (OAI)** is a special CloudFront user that you associate with your distribution to securely serve private content from an **Amazon S3 bucket**. It ensures that only CloudFront can access the S3 content, preventing direct access from the internet.

<!-- Card End -->
<!-- Card Start -->

### Front 

How does CloudFront handle caching?
### Back

CloudFront caches content at Edge Locations based on **cache behaviors** and **TTL (Time to Live)** settings. When a user requests content, CloudFront checks the cache:

- **Cache Hit**: Delivers content from the Edge Location.
- **Cache Miss**: Fetches content from the origin, caches it, and then delivers it to the user.

You can configure caching policies to control what is cached and for how long.

<!-- Card End --> 

<!-- Card Start -->

### Front 

What is an invalidation in CloudFront?

### Back

An **invalidation** is a request to remove one or more objects from CloudFront's cache before they expire. This ensures that subsequent requests for the invalidated objects are fetched from the origin, allowing you to update or delete content as needed.

<!-- Card End --> 
<!-- Card Start -->

### Front

How does CloudFront handle SSL/TLS?

### Back

CloudFront supports **HTTPS** to encrypt data between viewers and Edge Locations. You can use **default CloudFront certificates** or **custom SSL certificates** via **AWS Certificate Manager (ACM)** for your own domain names. It also supports **TLS 1.2** for enhanced security.

<!-- Card End --> 
<!-- Card Start -->

### Front 

What is Geo-Restriction in CloudFront?

### Back

**Geo-Restriction** allows you to **allow** or **block** content delivery to users based on their geographic locations (countries). This is useful for complying with licensing agreements, regional regulations, or content distribution policies.

<!-- Card End --> <!-- Card Start -->

### Front 

What are CloudFront Access Logs?

### Back

**CloudFront Access Logs** provide detailed records of every request made to your distribution. Logs include information such as the **viewer’s IP address**, **request time**, **HTTP status code**, **bytes served**, and more. They are useful for monitoring, analyzing traffic, and troubleshooting.

<!-- Card End --> 
<!-- Card Start -->

### Front 

Describe CloudFront Functions.

### Back

**CloudFront Functions** are lightweight JavaScript functions that run at Edge Locations. They allow you to **modify HTTP requests and responses** quickly and efficiently, enabling tasks like URL rewrites, header manipulations, and simple authentication without the overhead of Lambda@Edge.

<!-- Card End -->

<!-- Card Start -->

### Front 

How do Signed URLs and Cookies work in CloudFront?

### Back

**Signed URLs** and **Signed Cookies** are used to restrict access to your content:

- **Signed URLs**: Provide time-limited access to specific files.
- **Signed Cookies**: Allow access to multiple restricted files within a set time period.

They ensure that only authorized users can access your private content.

<!-- Card End --> <!-- Card Start -->

### Front 

How does CloudFront integrate with Amazon S3?

### Back

CloudFront can use an **Amazon S3 bucket** as an origin. When configured with an **Origin Access Identity (OAI)**, CloudFront securely serves private content from the S3 bucket, ensuring that content is accessible only through CloudFront and not directly from S3.

<!-- Card End --> <!-- Card Start -->

### Front


What are CNAMEs in CloudFront?

### Back

**CNAMEs (Canonical Names)** allow you to use your own domain names with CloudFront distributions. By adding **alternate domain names (CNAMEs)** to your distribution and configuring DNS records, you can serve content using custom URLs like `cdn.yourdomain.com`.

<!-- Card End --> <!-- Card Start -->

### Front

Explain Field-Level Encryption in CloudFront.

### Back

**Field-Level Encryption** allows you to encrypt specific data fields within HTTPS requests, protecting sensitive information like credit card numbers or personal data. CloudFront decrypts these fields before forwarding the request to the origin, ensuring data security end-to-end.

<!-- Card End -->

<!-- Card Start -->

### Front 

How does CloudFront handle compression?

### Back

CloudFront can automatically **compress** certain types of files (e.g., HTML, CSS, JavaScript) before delivering them to viewers. This reduces the amount of data transmitted, improving load times and reducing bandwidth usage. Compression settings can be enabled in cache behaviors.

<!-- Card End -->

<!-- Card Start -->

### Front 

What monitoring options does CloudFront provide?

### Back

CloudFront offers **real-time metrics and logs** through **Amazon CloudWatch**. You can monitor key performance indicators like **cache hit ratios**, **latency**, **request counts**, and **error rates**. Additionally, **access logs** provide detailed request information for in-depth analysis.

<!-- Card End --> 
<!-- Card Start -->

### Front 

What is Origin Shield in CloudFront?

### Back

**Origin Shield** is an additional caching layer that acts as a central point for all requests to your origin. It reduces the number of direct requests to the origin by consolidating requests from multiple Edge Locations, improving cache efficiency and reducing origin load.

<!-- Card End --> 
<!-- Card Start -->

### Front 

How can you customize error responses in CloudFront?

### Back

CloudFront allows you to define **custom error responses** for specific HTTP status codes (e.g., 404, 500). You can specify custom error pages, response codes, and caching durations, enhancing the user experience by providing branded or more informative error messages.

<!-- Card End --> 
<!-- Card Start -->

### Front  

Does CloudFront support HTTP/2?

### Back

Yes, **AWS CloudFront** fully supports **HTTP/2**, allowing for improved performance through features like multiplexing, header compression, and server push, enhancing the efficiency of content delivery to compatible clients.

<!-- Card End -->
<!-- Card Start -->

### Front

How does CloudFront manage access control?

### Back

CloudFront manages access control through:

- **Origin Access Identity (OAI)** for S3 buckets.
- **Signed URLs and Cookies** for restricting content access.
- **AWS IAM Policies** to control who can create and manage distributions.
- **AWS WAF** for filtering and blocking malicious requests.

<!-- Card End --> <!-- Card Start -->

### Front

How does CloudFront integrate with AWS Shield?

### Back

CloudFront integrates with **AWS Shield**, providing built-in protection against DDoS attacks. **AWS Shield Standard** is automatically enabled with CloudFront distributions, offering protection without additional costs, while **AWS Shield Advanced** provides enhanced protections and support for larger, more sophisticated attacks.

<!-- Card End -->


#flashcards #aws

status:: 
count:: 25
[[Flashcards]]

<!-- Card Start -->
### Front 
What is AWS Lambda and what is its primary purpose?

### Back
AWS Lambda is a serverless compute service that runs code in response to events without requiring server management. Its primary purposes are:

- **Event-driven execution**: Runs code in response to triggers from AWS services or HTTP endpoints
- **Automatic scaling**: Scales automatically from a few requests per day to thousands per second
- **Pay-per-use**: Charges only for compute time consumed
- **Supports multiple languages**: Including Node.js, Python, Java, Go, Ruby, and .NET Core

<!-- Card End -->

<!-- Card Start -->
### Front
What are the maximum execution duration limits for AWS Lambda functions?

### Back
- **Maximum execution duration**: 15 minutes (900 seconds)
- **Default timeout**: 3 seconds
- **Minimum timeout**: 1 second
- **Configuration**: Timeout can be set at deployment or update time
- **Best practice**: Set the timeout value based on expected function execution time plus a buffer

<!-- Card End -->

<!-- Card Start -->
### Front
Compare AWS Lambda and AWS Step Functions

### Back
**AWS Lambda**:
- Single-purpose functions
- Short-lived executions
- Event-driven processing
- Standalone compute service

**AWS Step Functions**:
- Orchestrates multiple AWS services
- Long-running workflows
- Visual workflow designer
- State machine-based coordination
- Maintains workflow state

<!--- Card Link --->
<!-- Card End -->

<!-- Card Start -->
### Front 
What are the main state types in AWS Step Functions?

### Back
1. **Task States**: Execute work (Lambda, AWS services)
2. **Choice States**: Add branching logic
3. **Parallel States**: Execute branches simultaneously
4. **Map States**: Process items in arrays
5. **Wait States**: Pause execution
6. **Pass States**: Pass data without work
7. **Succeed States**: End execution successfully
8. **Fail States**: End execution with failure

<!-- Card End -->

<!-- Card Start -->
### Front  
How does memory allocation work in AWS Lambda?

### Back
- **Range**: 128 MB to 10,240 MB (10 GB)
- **Increments**: 1 MB steps
- **CPU Scaling**: CPU power scales linearly with memory
- **Pricing**: Charged based on memory-time (GB-seconds)
- **Best practice**: Test different memory configurations to optimize cost/performance

<!-- Card End -->

<!-- Card Start -->
### Front 
What are the two types of workflows in Step Functions?

### Back
**Standard Workflows**:
- Long-running executions (up to 1 year)
- At-least-once execution model
- Full execution history
- Higher price point

**Express Workflows**:
- Short-lived executions (up to 5 minutes)
- At-most-once or at-least-once execution
- Limited execution history
- Lower price point, higher throughput

<!-- Card End -->

<!-- Card Start -->
### Front
What is a cold start in AWS Lambda and how to minimize it?

### Back
A cold start occurs when a new Lambda container is initialized, causing additional latency:

**Minimization strategies**:
- Use Provisioned Concurrency
- Implement function warming
- Optimize code package size
- Choose runtimes with faster startup (Node.js, Python)
- Keep dependencies minimal

<!-- Card End -->

<!-- Card Start -->
### Front
How are environment variables handled in AWS Lambda?

### Back
- **Storage**: Key-value pairs stored in function configuration
- **Size limit**: 4 KB total for all environment variables
- **Encryption**: Can be encrypted using KMS
- **Runtime access**: Available through standard runtime interfaces
- **Best practice**: Use for configuration that changes between environments

<!-- Card End -->

<!-- Card Start -->
### Front 
How does Step Functions handle data input and output?

### Back
- Uses **JsonPath** syntax for data manipulation
- **InputPath**: Selects portion of input JSON
- **OutputPath**: Filters output JSON
- **ResultPath**: Specifies where to place result
- **Parameters**: Manipulates input data
- Maximum event size: 256 KB

<!-- Card End -->

<!-- Card Start -->
### Front 
What is the default concurrent execution limit for Lambda?

### Back
- **Default limit**: 1,000 concurrent executions per region
- **Reserved concurrency**: Can be set per function
- **Unreserved concurrency**: Shared pool for all functions
- **Burst concurrency**: Initial burst of 500-3000 based on region
- Can request limit increase via AWS Support

<!-- Card End -->

<!-- Card Start -->
### Front
What error handling features are available in Step Functions?

### Back
- **Retry**: Configure retry attempts with backoff
- **Catch**: Handle specific error types
- **TimeoutSeconds**: Set state execution timeout
- **HeartbeatSeconds**: Monitor long-running tasks
- **FallbackStates**: Define alternate execution paths
- Error types: States.ALL, States.Timeout, States.TaskFailed

<!-- Card End -->

<!-- Card Start -->
### Front 
Compare Lambda@Edge with regular Lambda functions

### Back
**Lambda@Edge**:
- Runs at CloudFront edge locations
- Lower memory limit (128MB)
- Shorter timeout (5-30 seconds)
- Limited runtime support
- No VPC access

**Regular Lambda**:
- Runs in single AWS region
- Up to 10GB memory
- Up to 15 minutes timeout
- Full runtime support
- VPC access available

<!-- Card End -->

<!-- Card Start -->
### Front 
What are key security best practices for Lambda?

### Back
1. **IAM Roles**: Use least privilege permissions
2. **VPC**: Run in VPC when accessing private resources
3. **Secrets**: Use Secrets Manager or Parameter Store
4. **Dependencies**: Regularly update runtime and dependencies
5. **Logging**: Enable CloudWatch logging
6. **API Security**: Use API Gateway with authorization
7. **Code Signing**: Enable code signing for functions

<!-- Card End -->

<!-- Card Start -->
### Front 
What is a Task Token in Step Functions and when to use it?

### Back
- **Purpose**: Enable callback pattern for long-running tasks
- **Usage**: External systems can signal task completion
- **Implementation**: Using .waitForTaskToken
- **Timeout**: Can be configured up to 1 year
- **Best for**: Human approval workflows, external processing

<!-- Card End -->

<!-- Card Start -->
### Front 
What are Lambda Layers and their benefits?

### Back
**Lambda Layers**:
- Shared code and dependencies across functions
- Up to 5 layers per function
- Maximum size: 250 MB unzipped
- Can include custom runtimes
- Versioning support
- Reduces deployment package size
- Promotes code reuse

<!-- Card End -->

<!-- Card Start -->
### Front 
Compare Step Functions and SQS for workflow management

### Back
**Step Functions**:
- Visual workflow management
- Complex orchestration
- Built-in error handling
- State tracking
- Higher cost

**SQS**:
- Simple message queue
- Decoupled components
- Message retention
- Lower cost
- No visual workflow

<!-- Card End -->

<!-- Card Start -->
### Front 
What are Lambda Function URLs and their features?

### Back
- **Dedicated HTTPS endpoint** for Lambda function
- **Built-in CORS** support
- **Authentication**: IAM or NONE
- **Dual stack**: Supports IPv4 and IPv6
- **Request context**: Includes HTTP context
- **No additional charge**: Pay only for Lambda invocations
- Alternative to API Gateway for simple HTTP APIs

<!-- Card End -->

<!-- Card Start -->
### Front 
What are Step Functions Activities and their use cases?

### Back
- **Poll-based tasks** executed by external workers
- Used for **legacy system integration**
- Workers can run **anywhere** (on-premises, cloud)
- Supports **long-polling** (up to 60 seconds)
- Requires **manual acknowledgment**
- Good for **hybrid cloud** scenarios
- Alternative to Lambda for specialized workloads

<!-- Card End -->

<!-- Card Start -->
### Front 
How does Lambda integrate with VPC resources?

### Back
- Requires **ENI** (Elastic Network Interface)
- Needs **subnet IDs** and **security groups**
- Access to private resources like **RDS**, **ElastiCache**
- **Cold start** impact when using VPC
- Requires **IAM permissions** for ENI management
- No direct internet access without **NAT Gateway**
- Best practice: Only enable if required

<!-- Card End -->

<!-- Card Start -->
### Front 
What are the integration patterns for Step Functions?

### Back
**Integration Patterns**:
1. **Request Response**: Synchronous tasks
2. **Run a Job**: Asynchronous operations
3. **Wait for Callback**: External completion signal
4. **Optimized Integrations**: Direct service integration

**Supported Services**:
- Lambda
- AWS Batch
- DynamoDB
- ECS/Fargate
- SNS/SQS
- and many more

<!-- Card End -->

<!-- Card Start -->
### Front 
How do Dead Letter Queues work with Lambda?

### Back
- **Purpose**: Capture failed event processing
- **Supported destinations**: SQS or SNS
- **Configuration**: Set at function level
- **Event retention**: Based on destination service
- **Async invocation only**: Not for synchronous calls
- **Alternative**: Use destination configuration
- **Best practice**: Monitor DLQ for troubleshooting

<!-- Card End -->

<!-- Card Start -->
### Front 
What are key Lambda performance optimization techniques?

### Back
1. **Code Optimization**:
   - Minimize initialization code
   - Reuse connections and clients
   - Implement caching

2. **Configuration**:
   - Right-size memory allocation
   - Use Provisioned Concurrency
   - Optimize deployment package

3. **Architecture**:
   - Use async patterns when possible
   - Implement fan-out patterns
   - Consider Step Functions for orchestration

<!-- Card End -->

<!-- Card Start -->
### Front 
Compare AWS Step Functions and EventBridge Rules

### Back
**Step Functions**:
- Complex workflow orchestration
- Visual workflow designer
- State management
- Error handling and retry logic
- Long-running processes

**EventBridge Rules**:
- Event routing and filtering
- Simple event-driven patterns
- Schedule-based execution
- Fan-out pattern support
- No state management

<!-- Card End -->

<!-- Card Start -->
### Front
What are the characteristics of Lambda Container Images?

### Back
- **Size limit**: Up to 10 GB
- **Base images**: AWS-provided or custom
- **Registry**: Store in Amazon ECR
- **Runtime API**: Must implement Lambda Runtime API
- **Benefits**:
  - Consistent build environment
  - Larger deployment packages
  - Container tooling compatibility
- **Limitations**:
  - Longer cold starts
  - More complex deployment

<!-- Card End -->

#flashcards #aws

status:: 
count:: 29
[[Flashcards]]



<!-- Card Start -->

### Front

** Amazon QuickSight** 
### Back

Amazon QuickSight is a cloud-native, fully managed business intelligence (BI) service that enables organizations to create interactive dashboards, visualize data, and gain insights through machine learning capabilities.


<!-- Card End -->

<!-- Card Start -->

### Front

** S3 Express One Zone** Use cases
### Back

- Machine learning and artificial intelligence training
- Interactive data analytics 
- High performance computing (HPC)
- Data streaming
- Real-time advertising
- Media content workloads


<!-- Card End -->

## AWS EC2 Instance Connect

<!-- Card Start -->


### Front 

What is AWS EC2 Instance Connect?

### Back

AWS EC2 Instance Connect is a simple and secure way to connect to your Amazon EC2 instances using Secure Shell (SSH). It provides:

- Browser-based SSH access
- No need to manage SSH keys
- Integration with IAM for access control
- Support for Linux-based EC2 instances

<!-- Card End -->

<!-- Card Start -->

### Front

What are the main benefits of using EC2 Instance Connect?

### Back

The main benefits of EC2 Instance Connect include:

1. **Enhanced security**: No need to store or manage SSH keys
2. **Simplified access**: Connect directly from the AWS Management Console
3. **Fine-grained control**: Use IAM policies to manage access
4. **Audit trail**: All connection attempts are logged in AWS CloudTrail
5. **No additional cost**: Free feature included with EC2

<!-- Card End -->

<!-- Card Start -->

### Front

How does EC2 Instance Connect work?

### Back

EC2 Instance Connect works as follows:

1. User initiates a connection request
2. AWS generates a one-time SSH key pair
3. The public key is pushed to the instance metadata
4. The private key is used to establish the SSH connection
5. The key pair is automatically deleted after use

This process ensures secure, temporary access without long-term key management.

<!-- Card End -->

<!-- Card Start -->

### Front

What are the prerequisites for using EC2 Instance Connect?

### Back

Prerequisites for using EC2 Instance Connect:

1. EC2 instance running a supported Amazon Linux or Ubuntu AMI
2. Instance in a VPC with internet access (or VPC endpoints for EC2 and EC2 Instance Connect)
3. Security group allowing inbound traffic on port 22 (SSH)
4. IAM permissions to use EC2 Instance Connect
5. EC2 Instance Connect CLI (optional, for command-line access)

<!-- Card End -->

<!-- Card Start -->

### Front

What are the security best practices for EC2 Instance Connect?

### Back

Security best practices for EC2 Instance Connect:

1. Use IAM policies to restrict access to specific instances or users
2. Enable AWS CloudTrail to audit connection attempts
3. Use VPC endpoints to access EC2 Instance Connect without internet
4. Regularly review and update security group rules
5. Implement the principle of least privilege for IAM permissions
6. Use Multi-Factor Authentication (MFA) for IAM users
7. Monitor and analyze EC2 Instance Connect usage patterns

<!-- Card End -->

<!-- Card Start -->

### Front 

How does EC2 Instance Connect compare to Systems Manager Session Manager?

### Back

EC2 Instance Connect vs Systems Manager Session Manager:

| Feature | EC2 Instance Connect | Session Manager |
|---------|----------------------|-----------------|
| Protocol | SSH only | Multiple (SSH, RDP, CLI) |
| OS Support | Linux only | Linux, Windows, macOS |
| Bastion Host | Not required | Not required |
| Logging | CloudTrail | CloudTrail, S3, CloudWatch Logs |
| Port 22 | Required | Not required |
| IAM Integration | Yes | Yes |
| Additional Cost | No | No (but may incur data transfer costs) |

Choose based on specific requirements and use cases.

There many nuanced differences between these services but the basic idea is that EC2 Instance Connect allows for a convenient and secure native SSH connection using short-lived keys while Session Manager permits an SSH connection tunneled over a proxy connection.

<!-- Card End -->

<!-- Card Start -->

### Front

What are the limitations of EC2 Instance Connect?

### Back

Limitations of EC2 Instance Connect:

1. Only supports Linux-based instances (not Windows)
2. Requires internet access or VPC endpoints
3. Limited to SSH protocol
4. Not available for all instance types or regions
5. Requires specific AMIs or manual configuration
6. No support for custom SSH key management
7. May not be suitable for compliance requirements that mandate long-term key storage
8. Cannot be used with instances in EC2-Classic

Consider these limitations when deciding on remote access solutions.

<!-- Card End -->
## Questions

<!-- Card Start -->

### Front

What are some open-source alternatives to AWS Kinesis Data Streams?

### Back

Several open-source alternatives to AWS Kinesis Data Streams include:

1. **Apache Kafka**: A distributed streaming platform for high-throughput, fault-tolerant real-time data feeds.
2. **Apache Flink**: A stream processing framework for distributed, high-performing applications.
3. **Apache Storm**: A real-time computation system for processing large volumes of data.
4. **Spark Streaming**: Part of Apache Spark, it enables scalable, high-throughput stream processing.
5. **Apache NiFi**: A platform for automating data flow between software systems.

These tools offer similar capabilities for real-time data streaming and processing. 

<!-- Card End -->


<!-- Card Start -->

### Front

What are the main pricing tiers for AWS Support plans?

### Back

AWS offers four main support pricing tiers:

1. **Developer**: Minimum $29/month or 3% of monthly AWS charges
2. **Business**: Minimum $100/month or tiered percentage (10% for first $0-$10K, 7% for $10K-$80K, 5% for $80K-$250K, 3% over $250K)
3. **Enterprise On-Ramp**: Minimum $5,500/month or 10% of monthly AWS charges
4. **Enterprise**: Minimum $15,000/month or tiered percentage (10% for first $0-$150K, 7% for $150K-$500K, 5% for $500K-$1M, 3% over $1M)

Charges are based on the higher of the minimum or calculated percentage[1][7].

<!-- Card End -->

<!-- Card Start -->

### Front

How is the Developer Support plan fee calculated?

### Back

The Developer Support plan fee is calculated as follows:

- Minimum charge: $29/month
- If monthly AWS charges exceed $967 (3% of which is more than $29):
  - 3% of total monthly AWS charges

Example:
For $2,000 in monthly AWS charges:
$2,000 x 3% = $60

The higher amount ($60) would be charged as the support fee[5][7].

<!-- Card End -->

<!-- Card Start -->

## Enterprise Support Plan Benefits
### Front 

What are the key benefits of the Enterprise Support plan?

### Back

The Enterprise Support plan offers:

1. All features of Business Support
2. Dedicated Technical Account Manager (TAM)
3. Concierge Support Team
4. Support for use case reviews
5. Architectural guidance
6. Custom training options
7. Option to consolidate support across multiple AWS accounts within an organization
8. Simplified billing with a single invoice for the entire organization

Minimum fee: $15,000 per month (excludes additional AWS service usage charges)[2][7].

<!-- Card End -->


<!-- Card Start -->

## operational excellence
### Front 

What are the five design principles for operational excellence

### Back

 There are five design principles for operational excellence in the cloud: 
 
 - Perform operations as code 
 - Make frequent, small, reversible changes 
 - Refine operations procedures frequently 
 - Anticipate failure 
 - Learn from all operational failures

<!-- Card End -->

## AMI 

<!-- Card Start -->

### Front

What is an Amazon Machine Image (AMI)?
### Back

An Amazon Machine Image (AMI) is:

- A template for creating virtual servers (EC2 instances) in AWS
- Contains a pre-configured operating system and software stack
- Specific to region, OS, processor architecture, and virtualization type
- Used to launch multiple instances with the same configuration
- Can be created, acquired, or purchased from AWS Marketplace

AMIs include:
- A root volume template (OS, applications, etc.)
- Launch permissions
- Block device mapping for attached volumes

[1][5][7]

<!-- Card End -->

<!-- Card Start -->

### Front

What are the Types and Sources of AMIs
### Back

AMI types include:

1. Public AMIs: Provided by AWS, free to use
2. Paid AMIs: Available on AWS Marketplace, involve fees
3. Shared AMIs: Created and shared by other AWS users
4. Custom AMIs: Created by users for specific requirements

Sources of AMIs:
- AWS-provided
- AWS Marketplace
- Community AMIs
- Self-created custom AMIs

[2][8]

<!-- Card End -->

<!-- Card Start -->

### Front

Benefits of Using AMIs
### Back

Key benefits of using Amazon Machine Images:

1. Speed and efficiency in deployment
2. Consistency across multiple instances
3. Pre-configured software and permissions
4. Cost-effective scaling
5. Flexibility in OS and service options
6. Simplified software procurement (via AWS Marketplace)
7. Centralized policy enforcement
8. Version control for easy revision management
9. Integration with AWS services (e.g., Resource Access Manager, Organizations)
10. Reduced manual configuration and potential errors

[1][2]

<!-- Card End -->

<!-- Card Start -->

### Front

Creating a Custom AMI
### Back

Steps to create a custom AMI from an EC2 instance:

1. In AWS Management Console, go to EC2 Instances
2. Right-click the desired instance
3. Select "Create Image"
4. Provide a name and description for the AMI
5. Configure additional options if needed
6. Click "Create Image"
7. AWS will create a snapshot and register the AMI
8. The new AMI will appear in the AMIs list (may need to refresh)

Note: Creating an AMI may briefly stop the instance unless "No reboot" option is selected

[4][6]

<!-- Card End -->

## ETL

<!-- Card Start -->

### Front 
What is ETL?
 
### Back
 Extract, transform, and load (ETL) is the process of combining data from multiple sources into a large, central repository called a data warehouse. ETL uses a set of business rules to clean and organize raw data and prepare it for [storage](https://aws.amazon.com/what-is/cloud-storage/), [data analytics](https://aws.amazon.com/what-is/data-analytics/), and [machine learning (ML)](https://aws.amazon.com/what-is/machine-learning/). You can address specific business intelligence needs through data analytics (such as predicting the outcome of business decisions, generating reports and dashboards, reducing operational inefficiency, and more).

<!-- Card End -->

 

<!-- Card Start -->

### Front 

 Difference between data lake and data Warehouse
### Back
 
**Data warehouses**

A [data warehouse](https://aws.amazon.com/what-is/data-warehouse/) is a central repository that can store multiple databases. Within each database, you can organize your data into tables and columns that describe the data types in the table. The data warehouse software works across multiple types of storage hardware—such as solid state drives (SSDs), hard drives, and other cloud storage—to optimize your data processing.

**Data lakes**

With a [data lake](https://aws.amazon.com/what-is/data-lake/), you can store your structured and unstructured data in one centralized repository and at any scale. You can store data as is without having to first structure it based on questions you might have in the future. Data lakes also allow you to run different types of analytics on your data, like SQL queries, big data analytics, full-text search, real-time analytics, and machine learning (ML) to guide better decisions.

<!-- Card End -->

## Glue

<!-- Card Start -->

### Front 

What does ETL stand for and what is its purpose?

### Back

ETL stands for Extract, Transform, Load. It is a process used to:

- Move data from various sources to a unified repository
- Prepare data for analysis and business intelligence
- Consolidate data from multiple databases into a single repository
- Ensure data is properly formatted and qualified for analysis

ETL enables businesses to create a single source of truth for enterprise data, ensuring consistency and up-to-date information[1].

<!-- Card End -->

<!-- Card Start -->

### Front ETL Process Steps

What are the three main steps in the ETL process?

### Back

The three main steps in the ETL process are:

1. **Extract**: Pull raw data from multiple sources (e.g., CRM, ERP, IoT sensors)
2. **Transform**: Update data to match organizational needs and storage requirements
   - Standardize data types
   - Cleanse and resolve inconsistencies
   - Apply rules and functions
3. **Load**: Deliver and secure data for sharing within the organization

This process ensures data is ready for analysis and business intelligence purposes[1].

<!-- Card End -->

<!-- Card Start -->

### Front AWS Glue Overview

What is AWS Glue and what are its main features?

### Back

AWS Glue is a serverless data integration service that offers:

- Data discovery and organization
- Data transformation and preparation
- Building and monitoring of data pipelines
- Automatic ETL code generation
- Integration with other AWS services
- Support for various data sources and formats

AWS Glue simplifies the process of extracting, transforming, and loading data for analytics and data warehousing[2][6].

<!-- Card End -->

<!-- Card Start -->

### Front AWS Glue Benefits

What are the key advantages of using AWS Glue?

### Back

Key benefits of AWS Glue include:

1. Serverless architecture (no infrastructure management)
2. Automatic scaling based on data volume and complexity
3. Pay-only-for-what-you-use pricing model
4. Integration with other AWS services
5. Automated ETL code generation
6. Support for multiple data sources and formats
7. Built-in job scheduling and monitoring
8. Reduced time for data analysis
9. Collaboration features for teams[4][6][8]

<!-- Card End -->

<!-- Card Start -->

### Front AWS Glue Data Catalog

What is the AWS Glue Data Catalog and its purpose?

### Back

The AWS Glue Data Catalog is:

- A centralized metadata repository
- Used to store and manage information about data sources and stores
- Helps in unifying and searching across multiple data stores
- Allows for automatic data discovery using crawlers
- Enables schema management and permission control
- Increases data visibility across the organization

It serves as a foundation for data governance and discovery in AWS Glue[2][8].

<!-- Card End -->

<!-- Card Start -->

### Front ETL vs ELT

How does ETL differ from ELT?

### Back

ETL (Extract, Transform, Load) and ELT (Extract, Load, Transform) differ in the order of operations:

- ETL: Data is transformed before loading into the destination
- ELT: Data is loaded into the destination before transformation

ELT is more common in cloud-based data warehouses due to their scalable processing capabilities. It allows for:

- Preloading of raw data
- More flexibility for data scientists
- Faster processing by leveraging modern data processing engines
- Reduced data movement[1]

<!-- Card End -->

<!-- Card Start -->

### Front AWS Glue ETL Jobs

How does AWS Glue handle ETL jobs?

### Back

AWS Glue handles ETL jobs by:

1. Automatically generating ETL code in Scala or Python
2. Providing both visual and code-based interfaces
3. Offering job scheduling based on events or time
4. Supporting on-demand job execution
5. Automating logging, monitoring, and alerting
6. Facilitating job restarts in case of failures
7. Enabling parallel processing for heavy workloads
8. Allowing customization of generated ETL scripts

This approach simplifies ETL processes and improves efficiency[4][6][8].

<!-- Card End -->

<!-- Card Start -->

### Front AWS Glue Limitations

What are some limitations of AWS Glue?

### Back

While AWS Glue offers many benefits, it has some limitations:

1. Limited control over the underlying infrastructure
2. Potential higher costs for large-scale, continuous processing
3. Learning curve for those new to AWS services
4. Less flexibility compared to custom-built ETL solutions
5. Dependency on AWS ecosystem
6. Potential performance issues with very complex transformations
7. Limited support for real-time streaming data processing

Consider these factors when deciding to use AWS Glue for your ETL needs[6][10].

<!-- Card End -->

<!-- Card Start -->

### Front ETL Tools

What are some popular ETL tools besides AWS Glue?

### Back

Popular ETL tools include:

1. Informatica PowerCenter
2. Talend
3. IBM DataStage
4. Microsoft SQL Server Integration Services (SSIS)
5. Oracle Data Integrator
6. Pentaho Data Integration
7. Apache NiFi
8. Stitch
9. Fivetran
10. Matillion

These tools offer various features for data extraction, transformation, and loading, catering to different organizational needs and technical requirements[1][3].

<!-- Card End -->

<!-- Card Start -->

### Front ETL Best Practices

What are some best practices for implementing ETL processes?

### Back

Best practices for ETL implementation include:

1. Clearly define data quality standards and rules
2. Implement data profiling and cleansing
3. Use incremental loading when possible
4. Optimize for performance (e.g., parallel processing)
5. Implement error handling and logging
6. Version control your ETL code
7. Schedule regular maintenance and updates
8. Monitor and alert on job failures
9. Document data lineage and transformations
10. Ensure compliance with data governance policies

Following these practices can improve the reliability and efficiency of your ETL processes[1][3][7].

<!-- Card End -->

## more

<!-- Card Start -->

### Front  
Which of the following are pillars of the AWS Well-Architected Framework? **(Choose two.**)

- A. High availability
- B. Performance efficiency
- C. Cost optimization
- D. Going global in minutes
- E. Continuous development
 
### Back

B. Performance efficiency 
C. Cost optimization

<!-- Card End -->

<!-- Card Start -->

### Front  
A company wants to migrate unstructured data to AWS. The data needs to be securely moved with inflight encryption and end-to-end data validation.  
  
Which AWS service will meet these requirements?

- A. AWS Application Migration Service
- B. Amazon Elastic File System (Amazon EFS)
- C. AWS DataSync
- D. AWS Migration Hub
 
### Back

C. AWS DataSync

AWS DataSync is a service designed for securely transferring large amounts of data between on-premises storage and Amazon S3, Amazon EFS, or Amazon FSx for Windows File Server. It supports in-flight encryption and end-to-end data validation during the transfer process. In the context of the question, AWS DataSync is a suitable choice for securely moving unstructured data to AWS while ensuring encryption and data integrity throughout the migration process. Therefore, option C, AWS DataSync, would meet the specified requirements.

<!-- Card End -->

