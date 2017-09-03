#!/bin/bash

exec ./out/build/techscanservice \
  -dbPort=${dbPort} \
  -dbName=${dbName} \
  -dbServer=${dbServer} \
  -dbUserName=${dbUserName} \
  -dbPassword=${dbPassword} \
  -seedDataPath=${seedDataPath} \
  -originAllowed=${originAllowed} \
  $1
