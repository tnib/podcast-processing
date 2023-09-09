#!/bin/bash
rm -rf ./shownotes.txt
rm -rf ./title.txt
rm -rf ./twitter.txt
rm -rf ./transcript.txt
echo "starting transcribe step"
python3 ./transcribe.py || exit 1
echo "starting generation of show notes"
. ./generate_shownotes.sh || exit 1
echo "starting generation of the title"
. ./generate_title.sh || exit 1
echo "starting generation of the twitter"
. ./generate_twitter_post.sh || exit 1
echo "starting generation of the facebook"
. ./generate_facebook_post.sh || exit 1
echo "starting generation of the linkedin"
. ./generate_linkedin_post.sh || exit 1
