FROM alpine

RUN wget -q -O /etc/apk/keys/sgerrand.rsa.pub https://alpine-pkgs.sgerrand.com/sgerrand.rsa.pub

RUN wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.33-r0/glibc-2.33-r0.apk

RUN apk add glibc-2.33-r0.apk

# https://github.com/danielmiessler/SecLists/blob/master/Discovery/Web-Content/common.txt
COPY data/common.txt /usr/share/common.txt

# Copy go-loader
COPY go-loader /dist/go-loader

# Command to run when starting the container
ENTRYPOINT ["/dist/go-loader"]