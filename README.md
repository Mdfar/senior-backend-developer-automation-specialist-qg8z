Unió Digital Backend & Automation Hub

Comprehensive backend services and n8n workflows for MSP operations.

Architecture

API Layer: Go microservices for vendor communication.

Workflow Layer: n8n for business logic and ERP integration.

AI Layer: Python/RAG service for receipt and expense classification.

Database: PostgreSQL with pgvector (via Supabase).

Deployment

Set up your .env with Xero, HaloPSA, and OpenAI keys.

Run docker-compose up -d.

Import the n8n/workflow-erp-sync.json into your n8n instance.