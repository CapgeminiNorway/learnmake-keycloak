#FROM node:10-alpine
FROM bitnami/node:10
LABEL maintainer="Yilmaz Guleryuz 397125+zeusbaba@users.noreply.github.com"

ENV APP_HOME=/usr/src/app

# optionally if you want to run npm global bin without specifying path
ENV NPM_CONFIG_PREFIX=$APP_HOME/.npm-global
ENV PATH=$PATH:$APP_HOME/.npm-global/bin

RUN mkdir -p $APP_HOME
WORKDIR $APP_HOME

# Install APP dependencies
# A wildcard is used to ensure both package.json AND package-lock.json are copied
# where available (npm@5+)
COPY package*.json ./

RUN npm config set strict-ssl false
RUN npm install
# If you are building your code for production
# RUN npm install --only=production

# add @vue-cli for vue-cli-service
RUN npm install -g @vue/cli

# Bundle APP source
#COPY . .
COPY src ./src
#COPY config ./config
COPY public ./public
#COPY .* ./
COPY *.js ./

EXPOSE 8024

CMD npm start
