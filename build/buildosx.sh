#!/bin/bash

FILE="dist/ipfool-osx_darwin_${ARCH}/ipfool"
PACKAGE="dist/ipfool_${VERSION}_darwin_${ARCH}.zip"

echo "Signing ${FILE} with ${APPLE_DEVELOPER_ID}"
codesign --timestamp \
  --options=runtime \
  -s "${APPLE_DEVELOPER_ID}" \
  -v \
  "${FILE}"

# we need to create our own archive, goreleaser locks post-archival hooks
# away in their paid-tier
echo "Creating ${PACKAGE}"
zip "${PACKAGE}" "${FILE}"

# This submits a notarization request, the response may take hours so
# people on OSX might not actually be able to use the new version immediately
echo "Submitting ${PACKAGE} for notarization"
xcrun notarytool submit "${PACKAGE}" -v
  --apple-id "chadr@zang.com" \
  --team-id ${APPLE_DEVELOPER_ID} \
  --password ${APPLE_SIGNING_PASSWORD}

