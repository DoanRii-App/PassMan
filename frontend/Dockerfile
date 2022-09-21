FROM node:latest

WORKDIR /app

RUN mkdir temp && mkdir dist

COPY package*.json ./temp
RUN cd ./temp && npm install
RUN npm install -g http-server

COPY . ./temp
RUN  cd ./temp && npm run build
RUN mv ./temp/dist/* ./dist
RUN rm -r ./temp

EXPOSE 8000

CMD ["http-server", "dist", "-p 8000"]