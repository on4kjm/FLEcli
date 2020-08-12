#!/bin/bash
set -eu -o pipefail

CURRENT_DIR="$(cd "$(dirname "$0")" && pwd -P)"
TMP_DIR=${TMP_DIR:-/tmp}

## Cleanup
rm -rf "${CURRENT_DIR}/bats" "${TMP_DIR}/bats-core"

## Prepare
mkdir -p "${CURRENT_DIR}/helpers"

cat >"${CURRENT_DIR}/helpers/load.bash" <<EOF
#!/bin/bats
load "\${BATS_HELPERS_DIR}/bats-support/load.bash"
load "\${BATS_HELPERS_DIR}/bats-assert/load.bash"
load "\${BATS_HELPERS_DIR}/bats-file/load.bash"
load "\${BATS_HELPERS_DIR}/helpers.bash"
EOF

## Install bats from https://github.com/bats-core/bats-core
git clone https://github.com/bats-core/bats-core.git "${TMP_DIR}/bats-core"
bash "${TMP_DIR}/bats-core/install.sh" "${CURRENT_DIR}/bats"

## Install bats helper libraries from https://github.com/ztombol/bats-docs
for BATS_HELPER in bats-support bats-assert bats-file
do
    git clone "https://github.com/ztombol/${BATS_HELPER}" "${CURRENT_DIR}/helpers/${BATS_HELPER}"
    rm -rf "${CURRENT_DIR}/helpers/${BATS_HELPER}/.git" # No git tracking
done