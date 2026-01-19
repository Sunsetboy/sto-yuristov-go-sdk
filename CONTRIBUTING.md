# Contributing to sto-yuristov-go-sdk

## Development Setup

1. Clone the repository:
   ```bash
   git clone git@github.com:your-org/sto-yuristov-go-sdk.git
   cd sto-yuristov-go-sdk
   ```

2. Verify the code compiles:
   ```bash
   go build ./...
   go vet ./...
   ```

## Making Changes

1. Create a feature branch from `master`:
   ```bash
   git checkout master
   git pull origin master
   git checkout -b feature/your-feature-name
   ```

2. Make your changes and ensure they compile:
   ```bash
   go build ./...
   go vet ./...
   ```

3. Commit your changes:
   ```bash
   git add .
   git commit -m "Add description of your changes"
   ```

4. Push and create a pull request:
   ```bash
   git push origin feature/your-feature-name
   ```

5. After review, merge to `master`.

## Versioning

This SDK follows [Semantic Versioning](https://semver.org/):

- **MAJOR** (v1.0.0 → v2.0.0): Breaking changes to the API
- **MINOR** (v1.0.0 → v1.1.0): New features, backwards compatible
- **PATCH** (v1.0.0 → v1.0.1): Bug fixes, backwards compatible

### What counts as a breaking change?

- Removing or renaming exported functions, types, or methods
- Changing function signatures (parameters or return types)
- Changing struct field names or types
- Changing behavior in a way that breaks existing code

### What is NOT a breaking change?

- Adding new exported functions, types, or methods
- Adding new optional fields to request structs (with `omitempty`)
- Adding new error types
- Bug fixes that don't change the API

## Creating a New Release

### 1. Update CHANGELOG (if present)

Document changes in CHANGELOG.md before releasing.

### 2. Ensure master is up to date

```bash
git checkout master
git pull origin master
```

### 3. Verify the code

```bash
go build ./...
go vet ./...
```

### 4. Create and push a git tag

For a minor release (e.g., v1.1.0):
```bash
git tag v1.1.0
git push origin v1.1.0
```

For a patch release (e.g., v1.0.1):
```bash
git tag v1.0.1
git push origin v1.0.1
```

For a major release (e.g., v2.0.0):
```bash
git tag v2.0.0
git push origin v2.0.0
```

### 5. Verify the release

The new version is immediately available via Go modules:
```bash
go list -m github.com/your-org/sto-yuristov-go-sdk@v1.1.0
```

## Updating Consumers

After releasing a new version, update the worker projects:

```bash
# In each worker project directory
go get github.com/your-org/sto-yuristov-go-sdk@v1.1.0
go mod tidy
```

Or to get the latest version:
```bash
go get github.com/your-org/sto-yuristov-go-sdk@latest
```

## Major Version Releases (v2+)

Go modules require a different import path for major versions v2 and above.

### 1. Update go.mod

```go
module github.com/your-org/sto-yuristov-go-sdk/v2
```

### 2. Update all internal imports

All imports within the SDK must include `/v2`:
```go
import "github.com/your-org/sto-yuristov-go-sdk/v2"
```

### 3. Tag the release

```bash
git tag v2.0.0
git push origin v2.0.0
```

### 4. Consumers update their imports

```go
import stoyuristov "github.com/your-org/sto-yuristov-go-sdk/v2"
```

## Quick Reference

| Action | Command |
|--------|---------|
| Build | `go build ./...` |
| Vet | `go vet ./...` |
| Tag release | `git tag v1.x.x && git push origin v1.x.x` |
| List versions | `git tag -l` |
| Delete tag (if mistake) | `git tag -d v1.x.x && git push origin :refs/tags/v1.x.x` |
