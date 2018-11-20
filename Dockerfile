FROM golang

COPY ./ /go/src/phonebook-script
CMD dep ensure
CMD go install phonebook-script && phonebook-script