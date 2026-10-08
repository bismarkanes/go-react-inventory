# Use the official Node image
FROM node:22-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy dependency files first to utilize Docker's layer caching
COPY package*.json ./

# Install all dependencies (including devDependencies needed for npm run dev)
RUN npm install

# Copy the rest of your application code
COPY . .

# Expose the port your dev server runs on (e.g., 5173 for Vite, 3000 for Next.js)
EXPOSE 5173

# Start the application in development mode
CMD ["npm", "run", "dev"]
