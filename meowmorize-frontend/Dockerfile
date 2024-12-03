# meowmorize-frontend/Dockerfile

# Stage 1: Build the React application
FROM node:18-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy package.json and package-lock.json
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the source code
COPY . .

# Build the React app
RUN npm run build

# Stage 2: Serve the React app using nginx
FROM nginx:stable-alpine

# Copy the build output to nginx's html directory
COPY --from=builder /app/build /usr/share/nginx/html
 
EXPOSE 8999

# Start nginx
CMD ["nginx", "-g", "daemon off;"]
