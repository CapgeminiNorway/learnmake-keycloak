FROM alpine:3.8
LABEL maintainer="Yilmaz Guleryuz 397125+zeusbaba@users.noreply.github.com"

# copy executable prepared via go build
COPY demoserver ./
RUN chmod +x ./demoserver

EXPOSE 8026
CMD ./demoserver
