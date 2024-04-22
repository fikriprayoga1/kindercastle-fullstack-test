# Documentation
This program run on Infrastructur as Service(IaaS) and using docker

## Server
### Spesification
- Languange : Go
- Framework : Echo
- Database : Firebase Firestore

### Quick Start
#### Step 1
Ensure your PC has installed Docker & Git

#### Step 2
Open your terminal for Linux Based

#### Step 3
type `git clone https://github.com/fikriprayoga1/kindercastle-fullstack-test.git` and than enter

#### Step 4
type `cd kindercastle-fullstack-test` and than enter

#### Step 5
type `docker build -t kindercastle_image:1.0.0 .` and than enter

#### Step 6
type `docker compose up -d`

## Warning
#### Server not running
Ensure server running. You can type `docker logs kindercastle_container` and than enter to check your server

## Server Tester
You can test the server using postman collection. Import postman collection v2.1 in this folder

## Integration with Flutter
This server integrate with flutter app. You can use [this project link](https://app.flutterflow.io/project/kindercastle-6ha2bo) to try flutter app consume this server API

### Spesification
- This flutter application only consume 1 API because this flutter application build from flutter flow who restricted API only
- This flutter application using Firebase Authentication for login
