#!/bin/bash
# Principal-Grade Security CI/CD Gatekeeper
# This script enforces security policies before any code enters the main branch.

set -e

echo "--- Initializing Security Enforcement Gatekeeper ---"

# 1. SCA: Software Composition Analysis (Detecting vulnerable dependencies)
# Block if dependency has high/critical CVEs
echo "Running SCA check..."
snyk test --severity-threshold=high --file=go.mod

# 2. SAST: Static Application Security Testing (Code scanning for OWASP classes)
# Using Semgrep to enforce rules against SQLi, XSS, and Hardcoded Secrets
echo "Running SAST check (OWASP Baseline)..."
semgrep --config "p/owasp-top-ten" --error --timeout 300 ./src

# 3. Secret Scanning: Finding credentials in git history
# Fail if secrets are detected
echo "Running Secret Scanning..."
gitleaks detect --source . --verbose

# 4. Dependency Locking: Ensure no rogue packages
echo "Verifying Dependency Integrity..."
go mod verify

echo "--- Security Gatekeeper Passed: Production-Ready ---"
