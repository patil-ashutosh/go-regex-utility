GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOLANGCI = golangci-lint
SRCPKG = regex
SRCPKGPATH1  = github.com/patil-ashutosh/RegexUtility/regex/rvalidation
SRCPKGPATH2	 =  github.com/patil-ashutosh/RegexUtility/regex/rstring
TESTPKGPATH = github.com/patil-ashutosh/RegexUtility/testing


build:
		$(GOBUILD) ./$(SRCPKG)/...

test:
		$(GOTEST) -v -coverprofile=coverage.txt -coverpkg=$(SRCPKGPATH1),$(SRCPKGPATH2) $(TESTPKGPATH)

golinter:
		$(GOLANGCI) run --enable-all --exclude-use-default=false
