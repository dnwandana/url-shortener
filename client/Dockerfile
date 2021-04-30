FROM node:14.15.0-alpine
ARG API_ENDPOINT
ENV NEXT_PUBLIC_API_ENDPOINT ${API_ENDPOINT}
ENV NEXT_TELEMETRY_DISABLED 1
WORKDIR /client
COPY package.json yarn.lock ./
RUN yarn install
COPY . .
RUN yarn build
CMD [ "yarn", "start" ]