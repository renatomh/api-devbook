<h1 align="center"><img alt="API DevBook" title="API DevBook" src="https://go.dev/images/go-logo-blue.svg" width="250" /></h1>

# API DevBook

## 💡 Project's Idea

This project was developed while studying the Go programming language. It aims to create an API which provides routes to be used together with a web app or mobile application.

## 🔍 Features

* Creating new posts;
* Liking posts;
* Following other users;

## 🛠 Technologies

During the development of this project, the following techologies were used:

- [Go](https://go.dev/)
- [gorilla/mux](https://github.com/gorilla/mux)

## 💻 Project Configuration

### First, you must [install Go](https://go.dev/dl/) on your computer.

### Create the database tables according to the provided SQL scripts (located on the *sql* folder).

* The project was developed using MySQL;

## 🌐 Setting up config files

Create an *.env* file on the root directory, with all needed variables, credentials and API keys, according to the sample provided (*example.env*).

## ⏯️ Running

To run the application locally, execute the following command on the root directory (it'll show the application's available commands and options).

```bash
$ go run main.go 
```

## 🔨 Project's *Build*

In order to build the application and get the executable file, use the following command:

```bash
$ go build
```

Then, you can use this command to run the app (it'll show the application's available commands and options):

```bash
$ ./cli-app
```

### Documentation:
* [A Tour of Go](https://go.dev/tour/welcome/1)

## 📄 License

This project is under the **MIT** license. For more information, access [LICENSE](./LICENSE).
