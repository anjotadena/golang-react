FROM node:14.1.0-alpine3.10 as dev
WORKDIR /app
ENV PATH /app/node_modules/.bin:$PATH

COPY package.json ./
COPY package-lock.json ./
RUN npm install
RUN npm install react-scripts -g

# add app
COPY . .
RUN ls -lah

# start app
CMD ["npm", "start"]
# FROM node:14.1.0-alpine3.10 as dev
# WORKDIR /app/
# CMD echo "getting latest npm modules" \
#     && npm install \
#     && echo "installation complete, starting" \
#     && npm start

# FROM node:14.1.0-alpine3.10 as build
# WORKDIR /app
# COPY ./package.json ./package.json
# RUN npm install
# COPY . . 
# # Run app
# RUN npm run build -o app

# # multistage build...
# # this runs a lil nginx instance to serve the app
# FROM nginx:alpine as prod
# EXPOSE 3000
# COPY --from=build /app/build /usr/share/nginx/html
# COPY nginx.conf /etc/nginx/nginx.conf
