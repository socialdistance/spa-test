Example project
---

#### Build API server:
```bash
# Clone the repo
git clone https://github.com/socialdistance/spa-test
make buildx

# Up database
cd deployments
docker-compose up

# Do migrations with test data
make migration

#Run application
make run

# or manually

# Create container
docker-compose build
# Start container
docker-compose up

#Build from Makefile with test data
make build-test

#Start from Makefile with test data
make run-test

#Build from Makefile without data
make build-prod

#Start from Makefile without data
make run-prod

#Test
make test
```

