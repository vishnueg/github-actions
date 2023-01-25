#! /bin/bash

docker build -t initsetupscript:$1 . && docker tag initsetupscript:$1 sarveshdockerrepo/initsetupscript:$1 && docker push sarveshdockerrepo/initsetupscript:$1