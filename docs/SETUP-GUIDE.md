# Quick Setup Guide for CI/CD Pipeline

## Prerequisites

1. **GitHub Repository**: Ensure your repository is on GitHub
2. **GitHub Actions**: Enabled by default on public repos, enable on private repos
3. **Docker**: For building container images
4. **Kubernetes Cluster** (optional): For deployment

## Initial Setup

### 1. Repository Configuration

```bash
# Clone your repository
git clone https://github.com/your-username/trip-to-bd.git
cd trip-to-bd

# Create the necessary directories
mkdir -p .github/workflows
mkdir -p docs
```

### 2. Environment Setup

#### For Go Services
```bash
# Ensure Go 1.23+ is installed
go version

# Install golangci-lint for code quality
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

#### For Python Service
```bash
# Ensure Python 3.11+ is installed
python --version

# Install development dependencies
cd user-service
pip install flake8 black isort pytest
```

### 3. GitHub Secrets Configuration

Go to your GitHub repository → Settings → Secrets and variables → Actions

Add the following secrets:
- `DOCKER_REGISTRY_TOKEN` (if using external registry)
- `KUBECONFIG` (for Kubernetes deployments)
- `SLACK_WEBHOOK` (for notifications, optional)

### 4. Branch Protection Rules

Set up branch protection for `main` and `develop`:

1. Go to Settings → Branches
2. Add rule for `main`:
   - Require pull request reviews
   - Require status checks to pass
   - Require branches to be up to date
3. Add rule for `develop`:
   - Require pull request reviews
   - Require status checks to pass

## First Deployment

### 1. Test Locally

```bash
# Test Go services
cd hotel-service
go test ./...

# Test Python service
cd ../user-service
pytest tests/

# Test Docker builds
docker build -t hotel-service ./hotel-service
```

### 2. Push to Develop Branch

```bash
git checkout -b develop
git add .
git commit -m "feat: add CI/CD pipeline"
git push origin develop
```

### 3. Monitor Pipeline

1. Go to GitHub → Actions tab
2. Monitor the pipeline execution
3. Check for any failures and fix them

### 4. Deploy to Production

```bash
# Merge develop to main
git checkout main
git merge develop
git push origin main
```

## Customization

### 1. Service-Specific Configuration

Each service can have its own configuration:

```bash
# For Go services
cd hotel-service
# Add service-specific tests
# Update Dockerfile if needed

# For Python service
cd user-service
# Add service-specific dependencies
# Update requirements.txt
```

### 2. Environment Variables

Create environment-specific files:

```bash
# Development
cp .env.example .env.dev

# Production
cp .env.example .env.prod
```

### 3. Deployment Targets

Update the deployment scripts in the workflows:

```yaml
# In .github/workflows/ci-cd.yml
- name: Deploy to development environment
  run: |
    # Add your deployment logic here
    # Example: kubectl apply -f k8s/dev/
```

## Monitoring and Alerts

### 1. GitHub Notifications

- Watch the repository for all activities
- Set up email notifications for workflow failures

### 2. External Monitoring (Optional)

- Set up Slack webhook for deployment notifications
- Configure monitoring tools (Prometheus, Grafana)
- Set up APM tools (New Relic, DataDog)

## Troubleshooting

### Common Issues

1. **Build Failures**
   ```bash
   # Check Go version
   go version
   
   # Check Python version
   python --version
   
   # Test locally first
   docker build -t test-service ./service-name
   ```

2. **Test Failures**
   ```bash
   # Run tests locally
   go test ./...
   pytest tests/
   ```

3. **Deployment Issues**
   ```bash
   # Check Kubernetes access
   kubectl get pods
   
   # Check service health
   curl http://service-url/health
   ```

### Getting Help

1. Check GitHub Actions logs
2. Review the full documentation in `docs/CI-CD-PIPELINE.md`
3. Create an issue in the repository
4. Contact the DevOps team

## Next Steps

1. **Set up monitoring**: Configure Prometheus and Grafana
2. **Add security scanning**: Integrate SAST/DAST tools
3. **Implement blue-green deployments**: For zero-downtime deployments
4. **Add performance testing**: Load testing in the pipeline
5. **Set up disaster recovery**: Backup and restore procedures 