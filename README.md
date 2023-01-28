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

### Then, install the dependencies for the project

```bash
$ go mod download
```

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

Then, you can use this command to run the app:

```bash
$ ./api
```

### ⌨ Linux
In Linux systems, you can use the *api-devbook.service* file to run the application as a system service. The file must be updated and placed in the '/etc/systemd/system/' directory. After that, you should execute the following commands to enable and start the service:

```bash
$ sudo systemctl start api-devbook
$ sudo systemctl enable api-devbook
$ sudo systemctl status api-devbook
```

In order to serve the application with Nginx, it can be configured like so (adjusting the paths, server name, etc.):

```
# API DevBook
server {
    listen 80;
    server_name api-devbook.mhsw.com.br;

    location / {
        include proxy_params;
        proxy_pass http://localhost:5055;
        client_max_body_size 16M;
    }
}
```

#### 📜 SSL/TLS

You can also add security with SSL/TLS layer used for HTTPS protocol. One option is to use the free *Let's Encrypt* certificates.

For this, you must [install the *Certbot*'s package](https://certbot.eff.org/instructions) and use its *plugin*, with the following commands (also, adjusting the srver name):

```bash
$ sudo apt install snapd # Installs snapd
$ sudo snap install core; sudo snap refresh core # Ensures snapd version is up to date
$ sudo snap install --classic certbot # Installs Certbot
$ sudo ln -s /snap/bin/certbot /usr/bin/certbot # Prepares the Certbot command
$ sudo certbot --nginx -d api-devbook.mhsw.com.br
```

### Documentation:
* [A Tour of Go](https://go.dev/tour/welcome/1)
* [How To Install Go on Ubuntu 20.04](https://www.digitalocean.com/community/tutorials/how-to-install-go-on-ubuntu-20-04)
* [How To Deploy a Go Web Application Using Nginx on Ubuntu 18.04](https://www.digitalocean.com/community/tutorials/how-to-deploy-a-go-web-application-using-nginx-on-ubuntu-18-04)

## 📄 License

This project is under the **MIT** license. For more information, access [LICENSE](./LICENSE).
