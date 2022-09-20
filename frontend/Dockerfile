FROM node:latest

RUN npm install -g http-server

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .
RUN npm run build

EXPOSE 8000

CMD ["http-server", "dist", "-p 8000"]