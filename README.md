# MeowMorize

MeowMorize is a comprehensive flashcard application designed to help users effectively study and memorize information. Built with a robust Go backend and a dynamic React frontend, MeowMorize offers seamless interactions, efficient data management, and easy deployment using Docker.

Ok...that sounds good and created fron Chat.  It was.  This is really an experiment on using Chat to build an app from the ground up.   Its still a pretty decent flashcard app too.

## Architecture Overview

MeowMorize is composed of two main components:

- **Backend**: Implemented in Go using the Echo framework, the backend listens on port `8789` and manages the application's core functionalities, including user authentication, deck and card management, and session handling. Swagger documentation is available at `/swagger/index.html` for detailed API insights.

- **Frontend**: Developed with React, the frontend resides in the `/meowmorize-frontend` directory and listens on port `8999` as defined in the `.env` file. It provides an intuitive user interface for interacting with flashcards, importing data, and tracking study sessions.

The application leverages a SQLite database (`meowmorize.db`) located at the root directory for data persistence. To reset the application data, simply delete this database file.


## Project Status
This is still a work in progress and has a few things I would like to do, but its workable





## Getting Started

### Docker Deployment

MeowMorize is containerized using Docker, allowing for easy deployment and scalability. Below are the steps to deploy using Docker Compose.

#### Prerequisites

- **Docker**: [Install Docker](https://docs.docker.com/get-docker/)
- **Docker Compose**: [Install Docker Compose](https://docs.docker.com/compose/install/)

#### Steps to Deploy

1. **Build and Run Containers**

   From the project's root directory, execute:

   ```bash
   docker-compose up --build
   ```

   This command will build Docker images for both the backend and frontend services and start the containers.

2. **Accessing the Services**

   - **Backend**: Accessible at `http://localhost:8789`
   - **Frontend**: Accessible at `http://localhost:8999`

3. **Docker Compose Services**

   - **backend**
     - **Build Context**: Current directory (`.`)
     - **Dockerfile**: `Dockerfile.backend`
     - **Ports**: Maps host port `8789` to container port `8789`
     - **Environment Variables**:
       - `DB_PATH`: `/app/data/db.sqlite3`
     - **Volumes**:
       - `db-data`: Persists database data at `/app/data` inside the container
     - **Restart Policy**: `unless-stopped`

   - **frontend**
     - **Build Context**: `./meowmorize-frontend`
     - **Dockerfile**: `Dockerfile`
     - **Ports**: Maps host port `8999` to container port `80`
     - **Environment Variables**:
       - `REACT_APP_BACKEND_URL`: `http://backend:8789`
     - **Depends On**: `backend`
     - **Restart Policy**: `unless-stopped`

4. **Managing the Containers**

   - **Stop Containers**
     ```bash
     docker compose down
     ```

   - **Rebuild Containers**
     ```bash
     docker compose up --build
     ```

5. **Pushing Docker Images**

   Utilize the helper script to build and push Docker images to Docker Hub:

   ```bash
   ./helper push-docker
   ```

   **Note**: Ensure you are logged in to Docker Hub before executing this command.


### Building locally

#### Prerequisites

- **Go**: Ensure Go is installed on your machine. [Download Go](https://golang.org/dl/)
- **Node.js & npm**: Required for the React frontend. [Download Node.js](https://nodejs.org/)
- **Docker & Docker Compose**: For containerized deployment. [Install Docker](https://docs.docker.com/get-docker/)
- **Swag**: For generating Swagger documentation.
  ```bash
  go install github.com/swaggo/swag/cmd/swag@latest
  ```
- **Mockery**: For generating mock files.
  ```bash
  go install github.com/vektra/mockery/v2@latest
  ```

#### Installation

1. **Clone the Repository**
   ```bash
   git clone https://github.com/yourusername/meowmorize.git
   cd meowmorize
   ```

( Skip to docker compose if your not interested in actually putting this together)

2. **Set Up the Backend**
   
   The backend is managed through the helper script located at the root of the project.

   - **Make the Helper Script Executable**
     ```bash
     chmod +x helper
     ```

   - **Run the Backend Application**
     ```bash
     ./helper run
     ```
     This command executes `go run cmd/main/main.go`, starting the backend server on port `8789`.

3. **Set Up the Frontend**

   Navigate to the frontend directory and install dependencies:

   ```bash
   cd meowmorize-frontend
   npm install
   ```

   **Run the Frontend Application**
   ```bash
   npm start
   ```
   The frontend will be accessible at `http://localhost:8999`.

  **Run the Frontend Application from helper**
   ```bash
   ./helper npm-start
   ```
   The frontend will be accessible at `http://localhost:8999`.

### Building the Application

MeowMorize utilizes a helper script to streamline various build and deployment tasks. Below are the available commands and their descriptions:

```bash
./helper {run|concat|swagger|redoswagger|mocks|npm-start|npm-build|npm-test|push-docker|test|help}
```

##### Commands:

- **run**: Start the main backend application.
  ```bash
  ./helper run
  ```

- **concat**: Concatenate files for Chat pasting purposes.
  ```bash
  ./helper concat
  ```

- **swagger**: Initialize Swagger documentation.
  ```bash
  ./helper swagger
  ```

- **redoswagger**: Regenerate Swagger documentation from the `cmd/main` directory.
  ```bash
  ./helper redoswagger
  ```

- **mocks**: Generate mock files for testing.
  ```bash
  ./helper mocks
  ```

- **npm-build**: Build the React frontend application.
  ```bash
  ./helper npm-build
  ```

- **npm-start**: Start the React frontend application.
  ```bash
  ./helper npm-start
  ```

- **npm-test**: Run JavaScript tests in the frontend.
  ```bash
  ./helper npm-test
  ```

- **push-docker**: Build and push Docker images for both backend and frontend.
  ```bash
  ./helper push-docker
  ```

- **test**: Execute backend tests.
  ```bash
  ./helper test
  ```

- **help**: Display usage information.
  ```bash
  ./helper help
  ```

### Importing Flashcards

MeowMorize supports importing flashcards using a special Markdown format tailored for chat-friendly interactions. Each flashcard must adhere to the following structure:

```markdown
<!-- Card Start -->

### Front

[Question Here]

### Back

[Answer Here]

<!-- Card End -->
```

**Optional**: Include a link related to the flashcard by adding the `<!--- Card Link --->` comment before the end comment:

```markdown
<!-- Card Start -->

### Front

[Question Here]

### Back

[Answer Here]

<!--- Card Link ---> https://example.com/resource

<!-- Card End -->
```

#### Steps to Import:

1. **Prepare Your Markdown File**
   
   Ensure your flashcards are formatted correctly. Below is an example of 10 flashcards focusing on AWS Cloud certification topics related to AWS Lambda and AWS Step Functions:

   ```markdown
   <!-- Card Start -->
   ### Front
   What is AWS Lambda?
   
   ### Back
   AWS Lambda is a serverless compute service that lets you run code without provisioning or managing servers.
   
   <!-- Card End -->

   <!-- Card Start -->
   ### Front
   How do AWS Step Functions integrate with AWS Lambda?
   
   ### Back
   AWS Step Functions can coordinate multiple AWS Lambda functions into a workflow, managing state and transitions.
   
   <!-- Card End -->

   <!-- Add additional cards similarly -->
   ```

2. **Import via Frontend**
   
   - Navigate to the Import section in the frontend application.
   - Upload your Markdown file containing the flashcards.
   - The application will parse the file and add the flashcards to the selected deck.



## API Documentation

Comprehensive API documentation is available via Swagger at `http://localhost:8789/swagger/index.html`. It provides detailed information on all available endpoints, request parameters, and response structures.

## Running Tests

Ensure all dependencies are installed and run the following commands using the helper script:

- **Backend Tests**
  ```bash
  ./helper test
  ```

- **Frontend Tests**
  ```bash
  ./helper npm-test
  ```

## Notes

- **Database Reset**: To start fresh, simply delete the `meowmorize.db` file located at the root directory and restart the backend application.
- **Ports Configuration**: The frontend listens on port `8999` as specified in the `.env` file. Ensure this port is available or adjust the configuration as needed.
- **User Authentication**: The application initializes with a default user. Ensure to secure and manage user credentials appropriately.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
