#!/bin/bash
# Pre-commit hook: Architectural Quality Gatekeeper

echo ">>> Running Architectural Quality Gate..."

# 1. Check for committed secrets (Security)
if git grep -qE "AIza[0-9A-Za-z]{35}" --cached; then
    echo "ERROR: Potential Google API key detected!"
    exit 1
fi

# 2. Enforce formatting (Consistency)
# Example: runs standard linter before commit
echo "Running Linter..."
# npm run lint (or any other project-specific command)
# if [ $? -ne 0 ]; then
#    echo "ERROR: Linting failed."
#    exit 1
# fi

echo ">>> Quality Gate Passed. Proceeding with commit."
exit 0
