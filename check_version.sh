#!/bin/bash

set -e

# To prevent releasing a binary whose version is not updated, we compare the output of the
# --version flag to the release tag, trimming the leading characters ('v' from the tag,
# and 'steward ' from the cargo output)

BINARY_VERSION_STRING = $(cargo run --bin steward -- --version)
BINARY_VERSION = ${BINARY_VERSION_STRING: 8}
RELEASE_VERSION = ${GITHUB_REF_NAME: 1}

echo "Release version: $RELEASE_VERSION"
echo "Binary version: $BINARY_VERSION"

if [ "$RELEASE_VERSION" != "$BINARY_VERSION" ]; then
    >&2 echo "Version mismatch. Has the binary version been bumped?"
    exit 1
fi
