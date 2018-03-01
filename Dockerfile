FROM scratch
ADD main /
ADD config.json /
CMD ["/main"]