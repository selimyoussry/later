FROM alpine

RUN mkdir /app
ADD ./mybin /app/mybin

WORKDIR /app
RUN chmod +x ./mybin

EXPOSE 9080 8080
CMD ./mybin
