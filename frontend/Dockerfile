FROM node:21-alpine AS builder

WORKDIR /build

# Install npm packages
COPY package.json package-lock.json ./
RUN npm ci

COPY . .

ARG BACKEND_CONNECTOR
ARG BACKEND_URL

# Bundle app
RUN VITE_BACKEND_CONNECTOR=$BACKEND_CONNECTOR \
    VITE_BACKEND_URL=$BACKEND_URL \
    npm run build

FROM nginx:1.25-alpine

# Copy app bundle
COPY --from=builder /build/dist /usr/share/nginx/html

# Add nginx config
COPY nginx.conf /etc/nginx/conf.d/default.conf
