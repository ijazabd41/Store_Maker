#!/bin/bash

# Get current commit message
COMMIT_MSG=$(git log -1 --pretty=%s)

case "$COMMIT_MSG" in
    "hellllooooo")
        export GIT_AUTHOR_DATE="2024-10-01 10:00:00"
        export GIT_COMMITTER_DATE="2024-10-01 10:00:00"
        ;;
    "<:>")
        export GIT_AUTHOR_DATE="2024-11-15 10:00:00"
        export GIT_COMMITTER_DATE="2024-11-15 10:00:00"
        ;;
    "blahhhh")
        export GIT_AUTHOR_DATE="2024-11-20 14:00:00"
        export GIT_COMMITTER_DATE="2024-11-20 14:00:00"
        ;;
    "Components")
        export GIT_AUTHOR_DATE="2024-12-01 09:00:00"
        export GIT_COMMITTER_DATE="2024-12-01 09:00:00"
        ;;
    "frontend done")
        export GIT_AUTHOR_DATE="2024-12-15 16:00:00"
        export GIT_COMMITTER_DATE="2024-12-15 16:00:00"
        ;;
    *)
        echo "No date change needed for: $COMMIT_MSG"
        ;;
esac

# Amend the commit with the new date
git commit --amend --no-edit --date="$GIT_AUTHOR_DATE"