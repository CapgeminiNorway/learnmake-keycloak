#FROM node:10-alpine
FROM bitnami/node:10
LABEL maintainer="Yilmaz Guleryuz 397125+zeusbaba@users.noreply.github.com"

ENV APP_HOME=/usr/src/app

# optionally if you want to run npm global bin without specifying path
ENV NPM_CONFIG_PREFIX=$APP_HOME/.npm-global
ENV PATH=$PATH:$APP_HOME/.npm-global/bin

RUN mkdir -p $APP_HOME
WORKDIR $APP_HOME

# Install app dependencies
# A wildcard is used to ensure both package.json AND package-lock.json are copied
# where available (npm@5+)
COPY package*.json ./

RUN npm config set strict-ssl false
RUN npm install
# If you are building your code for production
# RUN npm install --only=production

## used to check if DB is up
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.4.0/wait ./wait
RUN chmod +x ./wait

# Bundle app source
#COPY . .
COPY src ./src
COPY config ./config
COPY public ./public
COPY .* ./

EXPOSE 2080

CMD /wait && npm start