## Uptime Checker

### Description
Provide a brief overview of the project and its key features.

### Features
- Register and login (JWT), verify using OTP (email)
- CRUD for websites
- Check websites in 30 seconds, 1 minute, 5 minutes, 30 minutes intervals
- Save website data to InfluxDB
- Send email alert for website status code 500
- Queue and job system using Redis
- Retry failed jobs
- Docker-compose for containerization

### Dependencies
- [jwt-go](https://github.com/golang-jwt/jwt): Package for generating and verifying JWT tokens
- [gin-gonic/gin](https://github.com/gin-gonic/gin): HTTP web framework

### Usage
1. Clone the repository
   ```bash
   git clone https://github.com/iya30n/uptime-checker.git
   ```
2. Set up the environment variables
3. Run the application using Docker-compose
   ```bash
   docker-compose up
   ```

### Contribution
If you spot any bugs, mistakes, or want to contribute to the project, feel free to create an issue or pull request.

This README provides a high-level overview of the project, its features, setup instructions, and references for implementing key components such as JWT authentication and web server. You can further expand it with detailed installation, configuration, and usage instructions based on your project's specific implementation.

Please note that the provided information is based on the available search results. If you have any specific requirements or need further assistance, feel free to ask!