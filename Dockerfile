FROM golang:alpine

# Set an env var that matches your github repo name, replace treeder/dockergo here with your repo name
ENV SRC_DIR=/go/src/github.com/bgiegel/TodoREST/

# Add the source code:
ADD . $SRC_DIR

# on travail dans le répertoire SRC
WORKDIR $SRC_DIR

# on récupère git..
RUN apk add --no-cache git

# ... puis govendor (outil de gestion des dépendances)
RUN go get github.com/kardianos/govendor

# On récupère toutes les dépendances du projet
RUN govendor sync

# On build
RUN cd $SRC_DIR; go build -o todorest

# et on lance
ENTRYPOINT ["./todorest"]