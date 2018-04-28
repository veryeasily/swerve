FROM golang:1.10-stretch
ENV SWERVEROOT=/go/src/github.com/elju/swerve
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
VOLUME $SWERVEROOT
EXPOSE 8080
WORKDIR $SWERVEROOT
CMD dep ensure && go build && ./swerve
