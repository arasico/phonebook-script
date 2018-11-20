FROM golang

COPY ./ /go/src/phonebook
CMD dep ensure
CMD go install phonebook && phonebook