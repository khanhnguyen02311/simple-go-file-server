#!/bin/sh
set -e

./bin/server \
    --port=${PORT:-1323} \
    --type=${TYPE:-local} \
    --upload-auth=${UPLOAD_AUTH:-false} \
    --download-auth=${DOWNLOAD_AUTH:-false} \
    --allowed-list=${ALLOWED_LIST:-image/png,image/jpeg,image/jpg,image/gif,image/webp} \
    --max-file-size=${MAX_FILE_SIZE:-10}