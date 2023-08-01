# Project Golang: QuizIslamicAI

QuizIslamicAI is a Golang-based project that aims to revolutionize the learning experience by combining personalization and automated content creation using the latest AI technology.

## Running the Golang Application Locally

To run the Golang application locally, follow these steps:

1. Make sure you have Go installed and have correctly configured the Go environment on your computer.

2. Open a terminal and navigate to your Golang project directory.

3. Run the following command to execute the application:
   ```
   go run main.go
   ```
   Alternatively, if you are using [Air](https://github.com/cosmtrek/air) as a runner for hot reload:
   ```
   air -c .air.toml
   ```

4. Your Golang application will be running and accessible at `http://localhost:8080/`.

## Postman Documentation

The Postman API documentation can be accessed at the following link: [Postman Documentation](https://documenter.getpostman.com/view/6565461/2s946icBSw)

## Deploying to the Public Server

This Golang application has been deployed to a public server with the IP address: `http://44.206.239.176:8080/`. You can access the application using this link.

## CI/CD with GitHub Actions

In this repository, we utilize GitHub Actions for Continuous Integration/Continuous Deployment (CI/CD) of the Golang application.

### CI/CD Process

Every time there is a push to the `main` branch, GitHub Actions will perform the following steps:

1. Checkout the source code from the repository.

2. Set up the SSH private key and allow access to the AWS server using webfactory/ssh-agent.

3. Pull from the repository on the AWS server to get the latest changes.

4. SSH to the AWS server, build a Docker image from the Golang code in the `/var/www/be-alta-quizislamic-ai/` directory, and run the Docker container using the newly created Docker image.

### GitHub Actions Configuration

The GitHub Actions configuration can be found in the `.github/workflows/go-push-docker.yml` file. You can customize the necessary steps for CI/CD according to your project's requirements.

Make sure you have set up the SSH_PRIVATE_KEY as a secret in your repository for successful SSH access to the AWS server.

Note: Ensure you have properly configured and set up the AWS server before performing CI/CD with GitHub Actions.

---

## Table of Contents
- [Project Golang: QuizIslamicAI](#project-golang-quizislamicai)
  - [Running the Golang Application Locally](#running-the-golang-application-locally)
  - [Postman Documentation](#postman-documentation)
  - [Deploying to the Public Server](#deploying-to-the-public-server)
  - [CI/CD with GitHub Actions](#cicd-with-github-actions)
    - [CI/CD Process](#cicd-process)
    - [GitHub Actions Configuration](#github-actions-configuration)
  - [Table of Contents](#table-of-contents)
  - [Pertemuan 1](#pertemuan-1)
  - [Pertemuan 2](#pertemuan-2)
  - [Pertemuan 3](#pertemuan-3)
  - [Nice To Have (Optional)](#nice-to-have-optional)

## Pertemuan 1
- [x] Ide: Understand the initial idea behind QuizIslamicAI and what problem it aims to solve in the education sector.
- [x] Repository: Set up the GitHub repository for the QuizIslamicAI project, including defining the project structure and initializing version control.

## Pertemuan 2
- [x] User Login/Register: Implement the user login and registration functionality for CourseGenAI. Users should be able to create accounts and log in to access personalized learning content.
- [x] CRUD (table apapun): Create basic CRUD operations for one of the tables in the database to efficiently manage data within the application.
- [x] Postman: Set up Postman to test the APIs and ensure they are working correctly for CourseGenAI.
- [x] Setup Docker: Dockerize the CourseGenAI application to make it easier to deploy and manage across different environments.

## Pertemuan 3
- [x] Live di AWS: Deploy CourseGenAI live on AWS (Amazon Web Services) to make the application accessible to users over the internet.
- [x] Continue Rest API: Continue working on the Rest API to enhance its functionality and make it more robust.

## Nice To Have (Optional)
- [x] Pagination: Implement pagination in CourseGenAI to efficiently handle large amounts of data and improve user experience.
- [x] Middleware (JWT, Static): Set up middleware for CourseGenAI, such as JWT (JSON Web Tokens) for authentication and Static Middleware for serving static assets.
- [x] Docs: Create comprehensive documentation for CourseGenAI, including API documentation and guidelines for developers who want to contribute to the project.
