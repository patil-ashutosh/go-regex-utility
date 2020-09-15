GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
SRCPKG = regex
SRCPKGPATH  = github.com/patil-ashutosh/RegexUtility/regex
TESTPKGPATH = github.com/patil-ashutosh/RegexUtility/testing


build:
		$(GOBUILD) ./$(SRCPKG)/...

test:
		$(GOTEST) -v -coverprofile=coverage.txt -coverpkg=$(SRCPKGPATH) $(TESTPKGPATH)
 
