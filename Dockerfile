
FROM golang as buildando

WORKDIR /app
EXPOSE 3000
# export SECRET_KEY_MP="APP_USR-5603718645176488-021311-c9e03f6eab82326f933417d74ab081d0-811772071"

CMD [ "tail", "-f", "/dev/null" ]

