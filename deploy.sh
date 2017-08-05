#!/bin/bash

go build -o build/scrumbot
scp build/scrumbot vds:~/scrumbot
