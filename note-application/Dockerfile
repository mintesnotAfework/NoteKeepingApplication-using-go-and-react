FROM node:22.7-alpine

RUN mkdir -p /app/react-frontend
WORKDIR /app/react-frontend
COPY package-lock.json  .
COPY package.json . 
COPY . .

RUN npm install


EXPOSE 3000

CMD [ "npm", "run", "dev" ]