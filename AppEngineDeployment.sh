#!/bin/bash


echo Please follow all on-screen prompts for the deployment process...

echo Installing Dependencies For the Project...
go get -u -v -f all
echo Initializing Google Cloud Deployment with gcloud SDK...
gcloud init --console-only
echo Now deploying to Google AppEngine...
gcloud app deploy
echo Deployed

