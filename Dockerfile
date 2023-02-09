
FROM golang as buildando



WORKDIR /app

CMD [ "tail", "-f", "/dev/null" ]



# FROM alpine

# WORKDIR /app
# COPY --from=buildando /app/gopix /app/



# ENTRYPOINT ./gopix