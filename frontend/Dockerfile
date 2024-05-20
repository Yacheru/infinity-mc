FROM node:alpine as frontend

WORKDIR /infinitymc

COPY package.json yarn.lock ./

RUN apk add yarn
RUN yarn install

COPY . /infinitymc

RUN yarn build


FROM nginx:alpine

WORKDIR /usr/share/nginx/html

RUN rm -rf ./*

COPY --from=frontend /infinitymc/dist /usr/share/nginx/html

COPY nginx.conf /etc/nginx/nginx.conf

ENTRYPOINT ["nginx", "-g", "daemon off;"]