#!/bin/bash

# .env ファイルから GOOGLE_PROJECT_ID, REGION を読み込む
# ない場合はエラーを出力して終了
if [ ! -f .env ]; then
  echo ".env file not found"
  exit 1
fi

if [ -z "$(cat .env | grep 'GOOGLE_PROJECT_ID=')" ]; then
  echo "GOOGLE_PROJECT_ID not found in .env file"
  exit 1
fi

if [ -z "$(cat .env | grep 'REGION=')" ]; then
  echo "REGION not found in .env file"
  exit 1
fi

export $(cat .env | grep 'GOOGLE_PROJECT_ID=*')
export $(cat .env | grep 'REGION=*')

gcloud run deploy counter-api --region=$REGION --source=. --update-env-vars HOST=0.0.0.0,GOOGLE_PROJECT_ID=$GOOGLE_PROJECT_ID --allow-unauthenticated --platform=managed

exit 0
