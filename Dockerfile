FROM gliderlabs/alpine:3.1
COPY . /src/github.com/progrium/toolbox
RUN apk-install -t build-deps go git mercurial \
	&& cd /src/github.com/progrium/toolbox \
	&& export GOPATH=/ \
	&& ./build.sh \
	&& rm -rf /src \
	&& apk del --purge build-deps
