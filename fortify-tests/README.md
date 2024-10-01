# Fortify Testing
This project is used to test fortify

## Prerequisites
- Python
- Poetry
- Docker
- Just

Run the following commands to install all of Playwrights dependencies
```bash
poetry run playwright install
poetry run playwright install-deps
```

## Test Types
- Integration(mark = `it`) - Integration tests assess if the interaction points between applications work as expected
- End To End(mark = `e2e`) - E2E tests run using Playwright and check the overall system and UI is working from an end user perspective
- Performance(mark = `perf`) - Performance tests run using Locust

## Getting Started
When testing files a inside the `env/` need to be passed in to define application service's addresses and other configuration options.
- `local-env` - can be used for testing against localhost

### Common Commands
- `just test-all <envFile>` - Run all integration and e2e tests. 
- `just test <envFile> <mark>` Run all tests that are marked with the selected mark. `pyproject.toml` has a list of all the marks that can be used
