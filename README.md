# GoLang based Microservices with Istio

## Architecture

Very simpled, without load balancer, istio and etc

![](https://i.imgur.com/GGcMXMb.png)

## Deployment

The DEPLOYING.md (ToDo) outlines deploying the stack to Google Kubernetes Engine (GKE)
on the Google Cloud Platform (GCP), with Istio and all associated telemetry
components: Prometheus, Grafana, Zipkin, Jaeger, Service Graph, and Kiali.
This README outlines deploying the Microservices/PostgreSQL/Redis stack locally to Docker Swarm.

### Requirements

- Docker
- Helm
- gcloud CLI
- Istio 1.1.x
- Jinja2 (pip install) - optional

#### Build images (Optional)

All Docker images, references in the Docker Swarm and Kubernetes resource files,
for the microservices are available on Docker Hub. To build all images yourself,
modify and use these two scripts.

```bash
bash ./1_build_images.sh
bash ./2_push_images.sh
```

Also you can remove all images use these script.

```bash
bash ./7_push_images.sh
```
