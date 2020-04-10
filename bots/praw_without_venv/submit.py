# $pip freeze > requirements.txt to extract the list of them.
# $pip install -r requirements.txt to install them later.

import os
import praw  # https://www.reddit.com/prefs/apps/
from termcolor import colored  # https://pypi.org/project/termcolor/

from settings import login

reddit = login()

targets = ["bottesting", "golang", "rust"]

for target in targets:
    subreddit = reddit.subreddit(target)
    subreddit.submit('praw', selftext='praw')

