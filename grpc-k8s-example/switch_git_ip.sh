#!/bin/bash

LOCAL_IP=`ipconfig getifaddr en0`
echo "LOCAL_IP:$LOCAL_IP"
grro="git remote remove origin"
`$grro`
grao="git remote add origin git@$LOCAL_IP:developer-learning/gcd.git"
`$grao`
git remote -v
