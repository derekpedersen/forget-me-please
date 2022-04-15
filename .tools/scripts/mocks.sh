#!/bin/bash

go version

# clean out existing mocks
rm -fr mock
mkdir mock

# example
# mockgen -source=repository/skateparks.go -destination=mock/mock_skateparks_repository.go -package=mock