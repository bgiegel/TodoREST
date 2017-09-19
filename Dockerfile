FROM iron/go:dev

WORKDIR /app

ADD . /app

# Set an env var that matches your github repo name, replace treeder/dockergo here with your repo name
ENV SRC_DIR=/go/src/github.com/bgiegel/TodoREST/

RUN go get github.com/gorilla/mux
RUN go get github.com/lib/pq

# Add the source code:
ADD . $SRC_DIR

# Build it:
RUN cd $SRC_DIR; go build -o todorest; cp todorest /app/

ENTRYPOINT ["./todorest"]