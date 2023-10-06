#!/bin/bash
echo "starting transcribe step"
python3 ./transcribe.py || exit 1
echo "starting generation of show notes and titles"
. ./generate_shownotes_and_titles.sh || exit 1
echo "starting generation of the twitter"
. ./generate_twitter_post.sh || exit 1
echo "starting generation of the facebook"
. ./generate_facebook_post.sh || exit 1
echo "starting generation of the linkedin"
. ./generate_linkedin_post.sh || exit 1
