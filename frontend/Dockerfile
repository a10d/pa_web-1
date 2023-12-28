FROM node:21-alpine AS builder

WORKDIR /build

# Install npm packages
COPY package.json package-lock.json ./
RUN npm ci

COPY . .

# Bundle app
RUN npm run build

FROM nginx:1.25-alpine

# Copy app bundle
COPY --from=builder /build/dist /usr/share/nginx/html

# Add nginx config
COPY nginx.conf /etc/nginx/conf.d/default.conf
