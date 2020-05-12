#!/usr/bin/env bash

echo "Current version tag $(git describe)"

read -p 'New release semantic version tag? (v1.x.x) ' newtag
read -p 'New release comment? ' newcmmt
echo -e "new commit version: \"$newtag\" comment: \"$newcmmt\"\nchangelog:\n$(cat changelog.md)\n"
read -p 'confirm? [y/N] ' confirm

case $confirm in
y | yes | ok) ;;
*)
    exit 0
    ;;
esac

git status
git add . &&
    git commit -m "$newcmmt" &&
    git tag -a $newtag -m "$newcmmt" &&
    git push origin $newtag &&
    goreleaser --release-notes changelog.md --rm-dist &&
    go get github.com/bengarrett/cfw

# note to delete a remote tag
# git tag --delete tagname
