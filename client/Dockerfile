FROM node:14.17.5-alpine
ARG API_ENDPOINT
ARG DOMAIN
ENV NEXT_PUBLIC_API_ENDPOINT ${API_ENDPOINT}
ENV NEXT_PUBLIC_DOMAIN ${DOMAIN}
ENV NEXT_TELEMETRY_DISABLED 1
WORKDIR /client
COPY package.json package-lock.json ./
RUN npm install
COPY . .
RUN npm run build
CMD [ "npm", "run", "start" ]
