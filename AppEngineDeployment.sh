#!/bin/bash

#You must be logged in to Google Cloud on the machine you run this script from. If you are not logged in, you will be prompted to login.

echo Please follow all on-screen prompts for the deployment process...

echo Installing Dependencies For the Project...
go get -u -v -f all
echo Initializing Google Cloud Deployment with gcloud SDK...
gcloud init --console-only
echo Now deploying to Google AppEngine...
gcloud app deploy
echo Configuring Firewall Rules for AppEngine...
gcloud app firewall-rules create 1 --action allow --source-range 0.0.0.0/0
gcloud app firewall-rules test-ip 72.69.124.135
echo Deployed

