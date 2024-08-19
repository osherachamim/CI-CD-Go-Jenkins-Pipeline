# CI-CD-Go-Jenkins-Pipeline
GoCI/CDFlow is a streamlined CI/CD pipeline project that automates the building and deployment of Go applications using Jenkins and Docker. It clones a Go repository, builds a Docker image, and runs the application in a container, showcasing an efficient continuous integration and delivery process.

# GoCI/CDFlow

GoCI/CDFlow is a CI/CD pipeline project designed to automate the building and deployment of a simple Go application using Jenkins and Docker. This project showcases the integration of Jenkins, Docker, and Go in a seamless CI/CD process.

## Tools Used

- **Jenkins**: An open-source automation server that facilitates continuous integration and continuous delivery (CI/CD).
- **Go**: A statically typed, compiled programming language designed for simplicity and efficiency.
- **Docker**: A platform for developing, shipping, and running applications in containers.

## Setting Up Jenkins Server + Docker Engine

To set up Jenkins and Docker on your server, run the following bash script:

```bash
#!/bin/bash

# Step 1: Update and upgrade the package list
echo "Updating package list and upgrading installed packages..."
apt update && sudo apt upgrade -y

# Step 2: Install Java and Fontconfig
echo "Installing OpenJDK 17 and fontconfig..."
apt-get install -y fontconfig openjdk-17-jre

# Verify Java installation
echo "Verifying Java installation..."
java -version

# Step 3: Add the Jenkins Repository
echo "Adding the Jenkins GPG key and repository..."
wget -O /usr/share/keyrings/jenkins-keyring.asc https://pkg.jenkins.io/debian-stable/jenkins.io-2023.key

echo "Adding Jenkins repository entry..."
echo "deb [signed-by=/usr/share/keyrings/jenkins-keyring.asc] https://pkg.jenkins.io/debian-stable binary/" | sudo tee /etc/apt/sources.list.d/jenkins.list > /dev/null

# Step 4: Install Jenkins
echo "Updating package list and installing Jenkins..."
apt update && sudo apt install -y jenkins

# Step 5: Start and enable Jenkins service
echo "Starting Jenkins service..."
systemctl start jenkins

echo "Enabling Jenkins to start at boot..."
systemctl enable jenkins

# Step 6: Adjust the Firewall
echo "Allowing Jenkins through the firewall on port 8080..."
ufw allow 8080
ufw allow 8081

echo "Checking firewall status..."
ufw status

echo "Jenkins installation and setup complete!"

# Install required packages for Docker
echo "Installing required packages for Docker..."
apt install apt-transport-https ca-certificates curl software-properties-common -y

# Add Docker’s official GPG key
echo "Adding Docker’s official GPG key..."
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg

# Add Docker repository
echo "Adding Docker repository..."
echo "deb [arch=amd64 signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

# Update package database with Docker packages
echo "Updating package database..."
apt update -y

# Install Docker
echo "Installing Docker..."
apt install docker-ce -y

# Start and enable Docker
echo "Starting and enabling Docker service..."
systemctl start docker
systemctl enable docker
```

# Jenkins Setup for Docker and Go Pipeline

## Initial Jenkins Setup

Once Jenkins is installed, visit `http://your-server-ip:8080` for the initial setup.

### Required Plugins

On Jenkins, install the following plugins:

- Docker API
- Docker Commons
- Docker Pipeline
- Docker
- Go

**Note:** After installing these plugins, a restart is required.

### Docker Permissions for Jenkins

To allow Jenkins to run Docker commands, add the Jenkins user to the Docker group with the following command:

```bash
sudo usermod -aG docker jenkins
```

## Restart Jenins Server 

```bash
sudo systemctl restart jenkins
```
# Jenkins CI/CD Pipeline for Go Application

Before you begin, make sure Jenkins is up and running, and you have Docker installed on the server where Jenkins is running.

## Setting Up the Pipeline

Once Jenkins is up and running, follow these steps:

1. **Configure Go in Jenkins:**
   - Go to **"Manage Jenkins"** > **"Global Tool Configuration"**.
   - Add a new Go installation and name it `go-1.23.0` (or the version you are using).

2. **Create a New Pipeline Job:**
   - Go to **"New Item"**, enter a name for your job, and choose **"Pipeline"**.
   - Scroll down to the **"Pipeline"** section and paste the following script:

   ```
   pipeline {
       agent any
       tools {
           go 'go-1.23.0'
       }

       stages {
           stage('Clone Repo') {
               steps {
                   // Get the code from a GitHub repository
                   git branch: 'main', url: 'https://github.com/osherachamim/CI-CD-Go-Jenkins-Pipeline.git'
               }
           }

           stage('Building Application Image') {
               steps {
                   script {
                       // Build the Docker image
                       app = docker.build("osherachamim/go-webapp-sample")
                   }
               }
           }

           stage('Running Application') {
               steps {
                   script {
                       // Run the Docker container
                       app.run("-p 8081:8081")
                   }
               }
           }
       }
   }


## Save and Build

Save the pipeline and click **"Build Now"**.

## What Does the Pipeline Do?

The pipeline performs the following tasks:

- **Clone the Repository:** Clones the Go application from the GitHub repository.
- **Build Docker Image:** Builds a Docker image using the `Dockerfile` and `hello-world.go` provided in the repository.
- **Run the Application:** Runs the application in a Docker container, exposing it on port `8081`.

## Accessing the Application

After the build completes, you can access the application by navigating to `http://your-server-ip:8081` in your web browser.

