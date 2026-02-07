# gopher-scope

A simple Go application instrumented with Prometheus and visualized in
Grafana.

Everything is Dockerized for quick setup. This repository is intended as
a trial-and-learn project and likely has no practical production use.

## Overview

The goal of this project is to explore:

- Basic Go application metrics
- Prometheus instrumentation
- Grafana dashboards
- Local observability tooling using Docker

## Stack

- Go
- Prometheus
- Grafana
- Docker / Docker Compose

## Usage

Clone the repository and start the stack:

```sh
docker compose up --build
```

There is also a Makefile that can be of good use.
Peace!
