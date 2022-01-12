FROM cosmtrek/air

WORKDIR /usr/app

COPY . .

RUN go get github.com/lib/pq

RUN go mod download
