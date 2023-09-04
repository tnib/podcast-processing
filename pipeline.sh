#!/bin/bash
rm -rf ./shownotes.txt
rm -rf ./title.txt
rm -rf ./twitter.txt
rm -rf ./transcript.txt
python3 ./transcribe.py
. ./generate_shownotes.sh
. ./generate_title.sh
. ./generate_twitter_post.sh
. ./generate_facebook_post.sh
. ./generate_linkedin_post.sh
