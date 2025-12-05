# CI/CD Pipeline Documentation

## Overview

This project implements a comprehensive CI/CD pipeline for a microservices architecture with 5 services:
- **Go Services**: hotel-service, location-service, restaurant-service, vehicle-service
- **Python Service**: user-service (FastAPI)

## Pipeline Architecture

### Workflows

1. **Main CI/CD Pipeline** (`.github/workflows/ci-cd.yml`)
   - Triggers on push to `main`/`develop` branches and pull requests
   - Runs tests, security scans, builds Docker images, and deploys

2. **Pull Request Checks** (`.github/workflows/pr-checks.yml`)
   - Code quality checks (linting, formatting)
   - Security vulnerability scanning
   - Test coverage reporting

3. **Release Management** (`.github/workflows/release.yml`)
   - Creates GitHub releases when tags are pushed
   - Generates changelog automatically
   - Builds and pushes versioned Docker images

4. **Manual Deployment** (`.github/workflows/manual-deploy.yml`)
   - Emergency deployments with manual triggers
   - Environment and service selection

## Pipeline Stages

### 1. Testing Stage
- **Go Services**: Unit tests, integration tests, coverage reporting
- **Python Service**: pytest with coverage
- **Parallel Execution**: All services tested simultaneously using matrix strategy

### 2. Security Stage
- **Trivy Vulnerability Scanner**: Scans for known vulnerabilities
- **Dependency Scanning**: Checks for outdated or vulnerable dependencies
- **Results Upload**: Uploads findings to GitHub Security tab

### 3. Build Stage
- **Docker Image Building**: Multi-stage builds for each service
- **Image Tagging**: Automatic tagging based on branch, commit, and version
- **Registry Push**: Pushes to GitHub Container Registry (ghcr.io)

### 4. Deploy Stage
- **Development**: Auto-deploy on `develop` branch pushes
- **Production**: Auto-deploy on `main` branch pushes
- **Health Checks**: Post-deployment verification

## Best Practices Implemented

### 1. Branch Strategy
```
main (production)
├── develop (staging)
├── feature/xyz
└── hotfix/xyz
```

### 2. Environment Management
- **Development**: Auto-deploy from `develop` branch
- **Staging**: Manual deployment option
- **Production**: Auto-deploy from `main` branch with approval

### 3. Security
- Vulnerability scanning with Trivy
- Dependency scanning
- Code quality checks
- Secure secrets management

### 4. Performance
- Parallel job execution
- Docker layer caching
- Go module caching
- Python pip caching

### 5. Monitoring
- Test coverage reporting
- Security scan results
- Deployment notifications
- Health check monitoring

## Configuration Files

### Code Quality
- `.golangci.yml`: Go linting configuration
- `pyproject.toml`: Python formatting and linting
- `.github/workflows/`: GitHub Actions workflows

### Docker
- Each service has its own `Dockerfile`
- Multi-stage builds for optimization
- Health checks included

## Usage Guide

### For Developers

1. **Creating a Feature**
   ```bash
   git checkout -b feature/new-feature
   # Make changes
   git commit -m "feat: add new feature"
   git push origin feature/new-feature
   # Create PR to develop
   ```

2. **Testing Locally**
   ```bash
   # Go services
   cd hotel-service
   go test ./...
   
   # Python service
   cd user-service
   pip install -r requirements.txt
   pytest tests/
   ```

3. **Code Quality**
   ```bash
   # Go
   golangci-lint run
   
   # Python
   black src/
   isort src/
   flake8 src/
   ```

### For DevOps/Operations

1. **Manual Deployment**
   - Go to GitHub Actions
   - Select "Manual Deployment" workflow
   - Choose environment, service, and version
   - Click "Run workflow"

2. **Creating a Release**
   ```bash
   git tag v1.0.0
   git push origin v1.0.0
   ```

3. **Monitoring Deployments**
   - Check GitHub Actions tab for pipeline status
   - Review security scan results in Security tab
   - Monitor health checks post-deployment

## Environment Variables

### Required Secrets
- `GITHUB_TOKEN`: Automatically provided
- `DOCKER_REGISTRY_TOKEN`: For external registries
- `KUBECONFIG`: For Kubernetes deployments
- `SLACK_WEBHOOK`: For notifications (optional)

### Service-Specific Variables
Each service can have its own environment variables defined in:
- `.env` files (for local development)
- GitHub repository secrets (for production)
- Kubernetes ConfigMaps/Secrets (for deployment)

## Troubleshooting

### Common Issues

1. **Build Failures**
   - Check Go version compatibility
   - Verify Python dependencies
   - Review Dockerfile syntax

2. **Test Failures**
   - Run tests locally first
   - Check for flaky tests
   - Review test coverage

3. **Deployment Issues**
   - Verify environment secrets
   - Check Kubernetes cluster access
   - Review health check endpoints

4. **Security Scan Failures**
   - Update vulnerable dependencies
   - Review false positives
   - Address critical vulnerabilities first

### Debugging

1. **Local Testing**
   ```bash
   # Test Docker builds locally
   docker build -t service-name ./service-name
   
   # Run with docker-compose
   docker-compose up --build
   ```

2. **Pipeline Debugging**
   - Check GitHub Actions logs
   - Review step-by-step execution
   - Verify environment variables

## Future Enhancements

1. **Advanced Monitoring**
   - Prometheus metrics integration
   - Grafana dashboards
   - APM (Application Performance Monitoring)

2. **Advanced Security**
   - SAST (Static Application Security Testing)
   - DAST (Dynamic Application Security Testing)
   - Container image signing

3. **Performance Optimization**
   - Multi-architecture builds
   - Image optimization
   - CDN integration

4. **Advanced Deployment**
   - Blue-green deployments
   - Canary deployments
   - Rollback automation

## Support

For issues with the CI/CD pipeline:
1. Check the GitHub Actions logs
2. Review this documentation
3. Create an issue in the repository
4. Contact the DevOps team 