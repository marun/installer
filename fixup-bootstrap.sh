#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

MOCP_REPO_ROOT=/home/maru/src/mocp/repos
SSH_ARGS='-o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -i ~/.ssh/ostest_rsa'
SCP_CMD="scp ${SSH_ARGS} ${MOCP_REPO_ROOT}"
HOST='core@192.168.126.10'
SSH_CMD="ssh ${SSH_ARGS} ${HOST}"

${SCP_CMD}/machine-api-operator/install/0000_30_machine-api-operator_02_machine.crd.yaml "${HOST}:machine-crd.yaml"
${SCP_CMD}/machine-api-operator/install/0000_30_machine-api-operator_03_machineset.crd.yaml "${HOST}:machineset-crd.yaml"
${SCP_CMD}/machine-config-operator/install/0000_80_machine-config-operator_01_machineconfig.crd.yaml "${HOST}:machine-config-crd.yaml"
${SSH_CMD} 'sudo oc --kubeconfig=/etc/kubernetes/kubeconfig create -f /home/core/machine-crd.yaml' || true
${SSH_CMD} 'sudo oc --kubeconfig=/etc/kubernetes/kubeconfig create -f machine-config-crd.yaml' || true
${SSH_CMD} 'sudo oc --kubeconfig=/etc/kubernetes/kubeconfig create -f machineset-crd.yaml' || true
${SSH_CMD} 'sudo oc --kubeconfig=/etc/kubernetes/kubeconfig create ns openshift-config' || true
${SSH_CMD} 'sudo oc --kubeconfig=/etc/kubernetes/kubeconfig create ns openshift-config-managed' || true
