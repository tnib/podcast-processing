#!/bin/bash
rm -rf ./shownotes.txt
rm -rf ./title.txt
rm -rf ./twitter.txt
rm -rf ./transcript.txt
python3 ./transcribe.py
. ./generate_shownotes.sh
. ./generate_title.sh
. ./generate_twitter_link.sh
