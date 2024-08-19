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
