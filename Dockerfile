FROM alpine

# Install the required packages
RUN apk add --update git go musl-dev
# Install the required dependencies
RUN go get github.com/go-telegram-bot-api/telegram-bot-api
RUN go get github.com/tidwall/gjson
# Setup the proper workdir
WORKDIR /root/OxfordDict-Bot
# Copy indivisual files at the end to leverage caching
COPY ./LICENSE ./
COPY ./OxfordDict.go ./
COPY ./utils.go ./
RUN go build

#Executable command needs to be static
CMD ["/root/OxfordDict-Bot/OxfordDict-Bot"]
