
---

# Project Report: Automated CI/CD Pipeline for Go-Based Web Application with Kubernetes and Argo CD on GKE

**Author:** Prajwal  
**Date:** 27-08-2024 

## Project Architecture
![Project Diagram drawio (1)](https://github.com/user-attachments/assets/4c1784f9-f1d5-4cfb-8b7e-46effdb44d1d)

## 1. Introduction

This project report outlines the creation and implementation of a comprehensive CI/CD pipeline for a Go-based web application, deployed on a Google Kubernetes Engine (GKE) cluster. The project integrates various DevOps tools and practices, including Docker, Kubernetes, Helm, and Argo CD, to achieve an automated and streamlined deployment process.

The project was inspired by a YouTube tutorial by Abhishek Veeramalla, which provided valuable guidance on setting up the core components of the CI/CD pipeline.

## 2. Project Objectives

- Develop a Go-based web application and test it locally.
- Containerize the application using Docker and push the image to Docker Hub.
- Deploy the application on a GKE cluster using Kubernetes.
- Set up Nginx Ingress to manage external access to the application.
- Implement Helm to manage Kubernetes deployments and configurations.
- Establish a CI pipeline using GitHub Actions to automate testing, Docker image creation, and Helm chart updates.
- Deploy Argo CD for Continuous Deployment, enabling automated updates to the Kubernetes cluster based on changes in the Git repository.

## 3. Project Setup and Implementation

### 3.1 Go Web Application Development
- Developed a simple web server using Go and tested it locally to ensure correct operation.
- The application served a static website built with HTML.

### 3.2 Containerization with Docker
- Created a Dockerfile to containerize the Go-based web application.
- Tested the Docker container locally, built, and tagged the Docker image.
- Pushed the image to a Docker Hub repository for later use in the deployment process.

### 3.3 Google Kubernetes Engine (GKE) Cluster Setup
- Set up a GKE cluster consisting of 3 nodes on Google Cloud Platform (GCP) to host the application.
- Created Kubernetes manifest files:
  - `deployment.yaml`: Defined the deployment of the Dockerized web application.
  - `service.yaml`: Set up a service to expose the application within the cluster.
  - `ingress.yaml`: Configured an Ingress resource to manage external access to the application.

### 3.4 Nginx Ingress Controller
- Installed an Nginx Ingress Controller on the GKE cluster to manage incoming HTTP/HTTPS traffic.
- Accessed the application via the load balancer IP address assigned to the Nginx Ingress.

### 3.5 Helm Integration
- Installed Helm on the GKE cluster to manage the deployment process more efficiently.
- Created a Helm chart named `go-site.chart`, with the Kubernetes manifest files added to the templates directory.
- Updated the `values.yaml` file to include the Docker image details, allowing for easy updates during deployment.

### 3.6 Source Control Management with Git
- Stored all project files, including the Go application source code, Dockerfile, Kubernetes manifests, and Helm chart, in a Git repository.
- This repository served as the Source Control Management (SCM) system for the project.

### 3.7 Continuous Integration with GitHub Actions
- Established a CI pipeline using GitHub Actions to automate the following tasks:
  - **Build and Test:** Compiled the Go application and ran unit tests to ensure code quality.
  - **Static Code Analysis:** Checked the code for potential issues using static analysis tools.
  - **Docker Image Creation and Push:** Built a new Docker image upon every commit and pushed it to Docker Hub.
  - **Helm Chart Update:** Updated the Helm chart values with the new Docker image tag, corresponding to the commit ID.
- Configured secrets required for the pipeline, such as Docker Hub credentials, securely in GitHub Secrets.

### 3.8 Continuous Deployment with Argo CD
- Deployed Argo CD on the GKE cluster to manage Continuous Deployment.
- Configured the project in Argo CD to monitor the Git repository for changes.
- Upon detecting a new commit, Argo CD automatically pulled the latest changes and deployed them to the GKE cluster.
- Ensured that the latest version of the application was always deployed and accessible via the Nginx Ingress load balancer.

## 4. Results and Outcomes
- **Automated CI/CD Pipeline:** Successfully implemented a fully automated CI/CD pipeline that streamlined the development, testing, and deployment of the Go-based web application.
- **Continuous Deployment with Argo CD:** Enabled seamless updates to the Kubernetes cluster, ensuring that the latest application version was always live without manual intervention.
- **Efficient Cluster Management:** Leveraged Helm and Argo CD for efficient management of Kubernetes resources, allowing for easy rollbacks and upgrades.

## 5. Challenges and Solutions
- **Kubernetes Configuration:** Ensuring proper configuration of Kubernetes resources, especially the Ingress Controller, required careful attention to detail. This was resolved through iterative testing and validation.
- **Pipeline Automation:** Integrating the CI pipeline with GitHub Actions and ensuring secure handling of secrets were critical challenges. These were addressed by following best practices for GitHub Actions workflows and secret management.

## 6. Conclusion
This project demonstrates the power and flexibility of modern DevOps tools and practices in automating the development and deployment processes for web applications. By integrating Docker, Kubernetes, Helm, GitHub Actions, and Argo CD, the project achieved a fully automated CI/CD pipeline that can serve as a template for future cloud-native application deployments.

The skills and knowledge gained from this project are highly applicable to real-world scenarios, where automation and scalability are key to managing complex applications in cloud environments.

## References
- **Video Tutorial:** [YouTube Video by Abhishek Veeramalla](https://www.youtube.com/watch?v=HGu9sgoHaJ0)
- **Documentation and Tools:** Docker, Kubernetes, Helm, GitHub Actions, Argo CD, Google Cloud Platform
- **Project Repository:** [devopsified-web-app](https://github.com/Praj0496/devopsified-web-app.git)
- **Docker image Repository:** [praj0404/go-webapp](https://hub.docker.com/r/praj0404/go-webapp)

**Contact Information:** [Portfolio site](https://www.prajwal.site/)

--- 
