#!/usr/bin/env bashset -e
# echo "" > coverage.txt

go test -v -coverprofile=coverage.txt -coverpkg=github.com/patil-ashutosh/RegexUtility/regex github.com/patil-ashutosh/RegexUtility/testing

# for d in $(go list ./... | grep -v vendor); do
#     go test -v -coverprofile=profile.out -coverpkg=-coverpkg=github.com/patil-ashutosh/RegexUtility/regex $d
#     if [ -f profile.out ]; then
#         cat profile.out >> coverage.txt
#         rm profile.out
#     fi
# done
