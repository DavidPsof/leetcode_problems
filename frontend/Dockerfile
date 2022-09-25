FROM node:18-alpine as base

WORKDIR /app/frontend/app

COPY ./app/package*.json ./

RUN npm install

ADD ./app .

CMD ["npm", "start"]