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

## Disclamer

At refactor stage of this repository, I have heavily used LLMs to do most
of the job for me since this was not the intention of the project.

I am doing an small refactor mainly to be able to add more stuff later on
and try out new things.

## Usage

Clone the repository and start the stack:

```sh
docker compose up --build
```

There is also a Makefile that can be of good use.
Peace!

