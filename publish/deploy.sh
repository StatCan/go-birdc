#!/bin/bash
# Deploy Script for Deploying the goproxy for the module

set -eo pipefail

[[ -z $VERCEL_TOKEN ]] && echo "VERCEL_TOKEN is not set" && exit 1
[[ -z $ENVIRONMENT ]] && echo "ENVIRONMENT is not set" && exit 1

[[ ! $(command -v vangen) ]] && echo "Please install vangen" && exit 1
[[ ! $(command -v vercel) ]] && echo "Please install vercel" && exit 1

# build the static site with go meta tags
vangen -config .vangen/vangen.json -out www/

# deploy the static site to vercel
vercel pull --yes --cwd ./www/ --environment=$ENVIRONMENT --token=$VERCEL_TOKEN
vercel build --cwd ./www/ --token=$VERCEL_TOKEN
[[ $ENVIRONMENT == "preview" ]] && vercel deploy --cwd ./www/ --prebuilt --token=$VERCEL_TOKEN || vercel deploy --cwd ./www/ --prebuilt --prod --token=$VERCEL_TOKEN

echo "Deployment to $ENVIRONMENT complete!"
