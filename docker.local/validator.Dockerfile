FROM blobber_base as validator_build

LABEL zchain="validator"


ENV SRC_DIR=/blobber
ENV GO111MODULE=on
#ENV GOPROXY=https://goproxy.cn,direct 


# Download the dependencies:
# Will be cached if we don't change mod/sum files
COPY .  $SRC_DIR
# COPY ./gosdk  /gosdk

RUN cd $SRC_DIR/ && go mod download 

WORKDIR $SRC_DIR/code/go/0chain.net/validator

RUN CGO_ENABLED=1 go build -v -tags "bn256 development" -ldflags "-X github.com/0chain/blobber/code/go/0chain.net/core/build.BuildTag=$GIT_COMMIT"

# Copy the build artifact into a minimal runtime image:
FROM golang:1.17.1-alpine3.14
RUN apk add gmp gmp-dev openssl-dev
COPY --from=validator_build  /usr/local/lib/libmcl*.so \
                        /usr/local/lib/libbls*.so \
                        /usr/local/lib/
ENV APP_DIR=/blobber
WORKDIR $APP_DIR
COPY --from=validator_build $APP_DIR/code/go/0chain.net/validator/validator $APP_DIR/bin/validator