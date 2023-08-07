# Ghstorage

SIMPELBI Module

## Release

```sh
go get -u all
go mod tidy
git tag                                 #check current version
git tag v0.0.1                          #set tag version
git push origin --tags                  #push tag version to repo
go list -m github.com/aiteung/ghstorage@v0.0.1   #publish to pkg dev, replace ORG/URL with your repo URL
```
