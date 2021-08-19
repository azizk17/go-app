# build environment
FROM node:16-alpine3.11 as build
WORKDIR /web

COPY package*.json ./
RUN npm install

# COPY vite.config.ts  ./

COPY . .
RUN node node_modules/esbuild/install.js

EXPOSE 3000

# # production stage
# FROM nginx:stable-alpine as production-stage
# COPY --from=build-stage /app/dist /usr/share/nginx/html
# EXPOSE 80
# CMD ["nginx", "-g", "daemon off;"]