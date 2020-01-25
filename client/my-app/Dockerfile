
FROM node:9.5 as builder

RUN mkdir /usr/src/app
WORKDIR /usr/src/app

#COPY packages.json packages.json

ENV PATH /usr/src/app/node_modules/.bin:$PATH

ADD package.json /usr/src/app/package.json
RUN npm install
RUN npm install -g react-scripts@1.1.1

COPY . .
RUN npm run build
RUN npm install -g serve

RUN ["pwd"]
RUN ["ls"]

CMD [ "serve", "-s", "build"]

#Run Stage Start
# FROM nginx:1.13.3-alpine

# RUN rm -rf /usr/share/nginx/html/*

# COPY nginx/nginx.conf /etc/nginx/conf.default
# COPY --from=builder /usr/src/app/build /usr/share/nginx/html

# CMD ["nginx", "-g", "daemon off;"]