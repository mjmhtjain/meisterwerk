# Current System Overview

Our platform currently consists of the following core services:

## 1. Order Scheduling Service
- Main domain service handling order management and scheduling
- Core business logic for order processing

## 2. Support Services
- **Authentication Service**: Handles user authentication and authorization
- **Document Service**: Manages document upload and download
- **Integration Service**: Handles communication with external services post-order completion

## 3. Technical Context
- Services are deployed on AWS
- Frontend communicates with backend via HTTP endpoints
- Primary languages: Golang and TypeScript/Node.js

## New Requirements
We need to design and implement a new domain for quote management with the following capabilities:
1. Product/Service Catalog
   - Manage master data for products and services
   - Handle pricing information
   - Configure and manage tax rates
2. Quote Management
   - Create and manage quotes for customers
   - Include products/services with respective prices and tax rates
   - Convert accepted quotes into orders in the existing order scheduling service
3. Integration Requirements
   - Support both single and bulk operations for all objects
   - Enable efficient UI interactions including sorting and filtering
4. Quote Management Domain
   - Design and implement a new domain for quote management with the following capabilities:
     1. Product/Service Catalog
        - Manage master data for products and services
        - Handle pricing information
        - Configure and manage tax rates
     2. Quote Management
        - Create and manage quotes for customers
        - Include products/services with respective prices and tax rates
        - Convert accepted quotes into orders in the existing order scheduling service
     3. Integration Requirements
        - Support both single and bulk operations for all objects
        - Enable efficient UI interactions including sorting and filtering
        - Facilitate external system integrations

## Interview Preparation Tasks
1. System Design
   Please prepare a system design diagram showing:
   - New services required
   - Integration with existing services
   - Data flow between services
   - AWS services to be utilized
   - APIs and communication patterns
2. Implementation Task
   Implement one of the following services in Golang:
   - Product Catalog Service with CRUD and bulk operations
   - Quote Management Service with basic quote creation and management
   Requirements for implementation:
   - OpenAPI 3.1 specification for the APIs
   - Clean architecture principles
   - Docker configuration

## Discussion Topics
During the interview, be prepared to discuss the following aspects of your design:
1. Architecture and System Design
   - Service boundaries and responsibilities
   - Data model and storage decisions
   - Caching strategy
   - Scalability considerations
   - Event-driven vs REST communication patterns
   - Error handling and resilience
   - Performance optimization strategies
2. API Design
   - RESTful API best practices
   - Bulk operation handling
   - Sorting and filtering
   - API versioning strategy
   - Documentation approach
   - Rate limiting and pagination
   - Error response standardization
3. Data Management
   - Data consistency patterns
   - Transaction handling
   - Data versioning approach
   - Master data management
   - Caching strategies
   - Data migration approach
4. Security
   - Authentication and authorization
   - API security best practices
   - Data encryption
   - Audit logging
   - Rate limiting
   - Security compliance considerations
5. Infrastructure and DevOps
   - AWS services selection and justification
   - Infrastructure as Code approach
   - CI/CD pipeline design
   - Monitoring and alerting strategy
   - Deployment strategies
   - Cost optimization
6. Testing Strategy
   - Unit testing approach
   - Integration testing
   - Performance testing
   - API contract testing
   - Test automation
   - Testing in CI/CD pipeline
7. Operational Excellence
   - Logging strategy
   - Monitoring and metrics
   - Alerting criteria
   - Incident response
   - Documentation practices
   - On-call considerations

## Evaluation Criteria
Your solution will be evaluated based on:
1. Architecture design clarity and justification
2. Code quality and organization
3. Testing approach and coverage
4. Security considerations
5. Operational readiness
6. Communication and explanation of technical decisions

## Submission Instructions
Please provide:
1. System design diagram (any standard format: draw.io preferred)
2. GitHub repository with your implementation
3. OpenAPI specification for your APIs
4. README or google doc/slides with:
   - Setup instructions
   - Design decisions and justifications
   - Testing approach
   - Deployment considerations
   Ideally send these materials 24 hours before the interview to allow for review.

## Interview Format
- 30 minutes: Your presentation of the system design
- 30 minutes: Deep dive into design decisions and implementation
- 30 minutes: Questions and discussion

## Additional Notes
- Feel free to make reasonable assumptions where requirements might be unclear
- Be prepared to explain and justify your technical choices
- Consider both immediate needs and future scalability
- Focus on clean, maintainable, and testable code
- Consider operational aspects such as monitoring and troubleshooting

Note: for any questions prior to the interview please feel free to reach out to lysanne@meisterwerk.com, mariano@meisterwerk.com or andi@meisterwerk.app