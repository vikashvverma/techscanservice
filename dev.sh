#!/bin/bash

export $(cat ./config/dev.env | xargs)

svc="techscanservice"

if [[ $(uname -s) == "Darwin" ]]; then
  ext=""
elif [[ $(uname -s) == *"MINGW"* ]]; then
  ext=".exe"
else
  ext=""
fi

#cp ./out/build/"${svc}${ext}" "./${svc}${ext}"
exec ./run_service.sh -port=${port}