FROM golang

COPY ./ /go/src/phonebook

CMD go install phonebook && phonebook