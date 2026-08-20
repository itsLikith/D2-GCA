<div align="center">
  <h1>Deployment Guide</h1>
  <h3>PRISM — Positioning & Resilience through Intelligent Signal Mapping</h3>

  <p>
    <img src="https://img.shields.io/badge/Deploy-Docker%20Compose-2496ED.svg?style=for-the-badge&logo=docker&logoColor=white" alt="Docker">
    <img src="https://img.shields.io/badge/Status-Production%20Ready-success.svg?style=for-the-badge" alt="Status">
  </p>
</div>

<br>

## Overview

PRISM is designed for simple, reliable deployment using Docker Compose.  
A single command is all that is required to bring the complete system online.

---

## Prerequisites

Ensure the following are installed on your system:

| Requirement | Minimum Version | Notes |
| :--- | :--- | :--- |
| **Docker** | 20.10+ | Engine must be running |
| **Docker Compose** | v2.0+ | Included with modern Docker Desktop |

> Verify installation with:
>
> ```bash
> docker --version
> docker compose version
> ```

---

## Quick Start

From the project root directory, run:

```bash
docker compose up -d