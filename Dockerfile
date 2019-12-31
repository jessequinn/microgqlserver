FROM alpine
ADD microgqlserver-api /microgqlserver-api
ENTRYPOINT [ "/microgqlserver-api" ]
