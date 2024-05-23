# Dockerfile for React app
FROM node:14 as build
WORKDIR /app
COPY package.json ./
COPY package-lock.json ./
RUN npm install
COPY . ./
RUN npm run build

FROM node:14
WORKDIR /app
COPY --from=build /app/build .
RUN npm install -g serve
EXPOSE 3000
CMD ["serve", "-s", ".", "-l", "3000"]
