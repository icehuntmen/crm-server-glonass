#!/bin/bash

# Read current version from VERSION file
CURRENT_VERSION=$(cat VERSION)

# Split version into major, minor, patch
IFS='.' read -r -a VERSION_PARTS <<< "$CURRENT_VERSION"
MAJOR=${VERSION_PARTS[0]}
MINOR=${VERSION_PARTS[1]}
PATCH=${VERSION_PARTS[2]}

# Increment patch version
PATCH=$((PATCH + 1))

# Create new version string
NEW_VERSION="$MAJOR.$MINOR.$PATCH"

# Update VERSION file without leading/trailing spaces
echo "${NEW_VERSION}" | sed -e 's/^[[:space:]]*//' -e 's/[[:space:]]*$//' > VERSION

# Add updated VERSION file to git index
git add VERSION

# Print updated version for confirmation
echo "Version updated to $NEW_VERSION"