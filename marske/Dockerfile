# Use an official Ubuntu runtime as a parent image
FROM ubuntu:20.04

# Set the working directory to /app
WORKDIR /app

# Set noninteractive mode for apt-get
ENV DEBIAN_FRONTEND=noninteractive

# Update and install necessary packages
RUN apt-get update && \
    apt-get install -y \
    curl \
    build-essential \
    python3 \
    python3-pip  # Install Python and pip

# Alias python3 to python
RUN ln -s /usr/bin/python3 /usr/bin/python

# Install Node.js using NodeSource repository
RUN curl -fsSL https://deb.nodesource.com/setup_20.x | bash -
RUN apt-get install -yq nodejs

# Install Rust
RUN curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y

# Add Rust binaries to the PATH
ENV PATH="/root/.cargo/bin:${PATH}"

# Install TypeScript globally
RUN npm install -g typescript

# Copy the current directory contents into the container at /app
COPY . /app

# Run your command (replace with your actual command)
CMD ["./run_steps.sh"]
