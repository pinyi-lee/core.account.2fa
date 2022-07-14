#!/bin/bash
export GO_HTTP_PORT='9999'
export LOG_LEVEL='INFO'
export DEPLOY_ENVIRONMENT='DEVELOP'
export MONGO_URI='mongodb://localhost:27017/core-account-2fa-service?rm.failover=1000ms:5x1&rm.monitorRefreshMS=100&rm.nbChannelsPerNode=1'