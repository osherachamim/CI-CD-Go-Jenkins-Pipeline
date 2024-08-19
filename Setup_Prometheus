# CI-CD-Go-Jenkins-Pipeline

#GoCI/CDFlow is a streamlined CI/CD pipeline project that automates the testing, building, and deployment of Go applications using Jenkins and Docker. It clones a Go repository, builds a Docker image, and runs the application in a container, showcasing an efficient continuous integration and delivery process.

## Why Use Prometheus

Prometheus is a powerful monitoring and alerting tool designed to handle the high demands of modern infrastructure and applications. Here are some reasons why integrating Prometheus into your CI/CD pipeline is beneficial:

- **Comprehensive Metrics Collection**: Prometheus allows you to collect detailed metrics from your applications, infrastructure, and services. This helps in understanding the performance and behavior of your Go applications in real-time.

- **Real-Time Monitoring**: With Prometheus, you can monitor your application and infrastructure in real-time. This is crucial for detecting and resolving issues before they impact your end-users.

- **Flexible Querying**: Prometheus uses a powerful query language called PromQL, which allows for flexible and complex queries to analyze and visualize metrics, making it easier to troubleshoot issues.

- **Alerting Capabilities**: Prometheus can be configured to trigger alerts based on specific conditions, allowing you to proactively manage and resolve potential issues.

- **Scalability**: Prometheus is designed to scale, making it suitable for environments of all sizes, from a single application to large, distributed systems.

- **Integration with Grafana**: Prometheus can be easily integrated with Grafana for enhanced visualization of your metrics, providing a more intuitive and interactive way to explore data.

# Prometheus Installation Guide on Ubuntu

This guide will help you set up Prometheus on an Ubuntu system. Follow the steps below to install, configure, and start Prometheus.

## Step 1 - Update System Packages

You should first update your system's package list to ensure that you are using the most recent packages. To accomplish this, issue the following command:

```bash
sudo apt update
```

## Step 2 - Create a System User for Prometheus

Now create a group and a system user for Prometheus. To create a group and then add a user to the group, run the following commands:

```bash
sudo groupadd --system prometheus
sudo useradd -s /sbin/nologin --system -g prometheus prometheus
```

This will create a system user and group named "prometheus" for Prometheus with limited privileges, reducing the risk of unauthorized access.

## Step 3 - Create Directories for Prometheus

To store configuration files and libraries for Prometheus, you need to create a few directories. The directories will be located in the `/etc` and the `/var/lib` directory respectively. Use the commands below to create the directories:

```bash
sudo mkdir /etc/prometheus
sudo mkdir /var/lib/prometheus
```

## Step 4 - Download Prometheus and Extract Files

To download the latest update, go to the Prometheus official downloads site and copy the download link for the Linux Operating System. Download using `wget` and the link you copied like so:

```bash
wget https://github.com/prometheus/prometheus/releases/download/v2.54.0/prometheus-2.54.0.linux-amd64.tar.gz
```

After the download has been completed, run the following command to extract the contents of the downloaded file:

```bash
tar vxf prometheus*.tar.gz
```

## Step 5 - Navigate to the Prometheus Directory

After extracting the files, navigate to the newly extracted Prometheus directory using the following command:

```bash
cd prometheus*/
```

Changing to the Prometheus directory allows for easier management and configuration of the installation. Subsequent steps will be performed within the context of the Prometheus directory.

## Configuring Prometheus on Ubuntu 22.04

With Prometheus downloaded and extracted on Ubuntu 22.04, you can go on to configure it. Configuring Prometheus could involve specifying the metrics to be collected, defining targets to scrape metrics from, and configuring alerting rules and recording rules. This guide focuses on the default setup that enables you to access Prometheus through your web browser.

### Step 1 - Move the Binary Files & Set Owner

First, you need to move some binary files (`prometheus` and `promtool`) and change the ownership of the files to the "prometheus" user and group. You can do this with the following commands:

```bash
sudo mv prometheus /usr/local/bin
sudo mv promtool /usr/local/bin
sudo chown prometheus:prometheus /usr/local/bin/prometheus
sudo chown prometheus:prometheus /usr/local/bin/promtool
```

### Step 2 - Move the Configuration Files & Set Owner

Next, move the configuration files and set their ownership so that Prometheus can access them. To do this, run the following commands:

```bash
sudo mv consoles /etc/prometheus
sudo mv console_libraries /etc/prometheus
sudo mv prometheus.yml /etc/prometheus
sudo chown prometheus:prometheus /etc/prometheus
sudo chown -R prometheus:prometheus /etc/prometheus/consoles
sudo chown -R prometheus:prometheus /etc/prometheus/console_libraries
sudo chown -R prometheus:prometheus /var/lib/prometheus
```

The `prometheus.yml` file is the main Prometheus configuration file. It includes settings for targets to be monitored, data scraping frequency, data processing, and storage. You can set alerting rules and notification conditions in the file. You don't need to modify this file for this demonstration but feel free to open it in an editor to take a closer look at its contents.

```bash
sudo nano /etc/prometheus/prometheus.yml
```

### Step 3 - Create Prometheus Systemd Service

Now, you need to create a system service file for Prometheus. Create and open a `prometheus.service` file with the Nano text editor using:

```bash
sudo nano /etc/systemd/system/prometheus.service
```

Include these settings to the file, save, and exit:

```ini
[Unit]
Description=Prometheus
Wants=network-online.target
After=network-online.target

[Service]
User=prometheus
Group=prometheus
Type=simple
ExecStart=/usr/local/bin/prometheus \
    --config.file /etc/prometheus/prometheus.yml \
    --storage.tsdb.path /var/lib/prometheus/ \
    --web.console.templates=/etc/prometheus/consoles \
    --web.console.libraries=/etc/prometheus/console_libraries

[Install]
WantedBy=multi-user.target
```

### Step 4 - Reload Systemd

You need to reload the system configuration files after saving the `prometheus.service` file so that changes made are recognized by the system. Reload the system configuration files using the following:

```bash
sudo systemctl daemon-reload
```

### Step 5 - Start Prometheus Service

Next, you want to enable and start your Prometheus service. Do this using the following commands:

```bash
sudo systemctl enable prometheus
sudo systemctl start prometheus
```

### Step 6 - Check Prometheus Status

After starting the Prometheus service, you may confirm that it is running or if you have encountered errors using:

```bash
sudo systemctl status prometheus
```

## Access Prometheus Web Interface

Prometheus runs on port 9090 by default, so you need to allow port 9090 on your firewall. Do that using the command:

```bash
sudo ufw allow 9090/tcp
```

## Conclusion

You have successfully installed and configured Prometheus on your Ubuntu system. You can now access the Prometheus web interface by navigating to `http://your-server-ip:9090` in your web browser.
