# Go Fiber Project

## What is Fiber?
Fiber is an Express-inspired web framework built on top of Fasthttp, known for its high performance and minimal configuration.

## Why Fiber?
Fiber is chosen for this project due to its:
- Simplified routing and middleware setup.
- High performance and low latency.
- Robust error handling and logging capabilities.

## Project Description
This project is a REST API for managing leads, utilizing Fiber for routing and Gorm for database interactions.

## Example Usage
The application sets up routes for managing leads:
- `GET /api/v2/leads` - Retrieve all leads
- `GET /api/v2/lead/:id` - Retrieve a lead by ID
- `POST /api/v2/lead` - Create a new lead
- `DELETE /api/v2/lead/:id` - Delete a lead by ID

Middleware is used for error handling and logging, ensuring a robust application.
