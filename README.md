# DB Composer

DB Composer is a simple tool to quickly spin up local database instances using Docker Compose. It supports multiple database engines, including PostgreSQL, MySQL, MariaDB, SQL Server, Oracle, and MongoDB.

## 📹 Demo

[![DB Composer Execution](https://img.youtube.com/vi/YOUR_VIDEO_ID/0.jpg)](https://github.com/iamgabs/database-composer/blob/main/assets/DB-Composer.mp4)

> Click the image to watch the demo.

## 🚀 Getting Started

### 1. Download the Binary

Go to the [Releases](https://github.com/YOUR-REPO/releases) page and download the `app` binary for your system.

### 2. Give Execution Permission

```sh
chmod +x app
```

### 3. Run DB Composer

```sh
./app
```

This will launch a terminal user interface (TUI) where you can select a database, its version, and start a local instance using `docker compose up`.

## 🛠 Supported Databases

- **PostgreSQL** (Versions: 16, 15, 14, 13, 12, 11)
- **MySQL** (Versions: 8.0, 5.7, 5.6, 5.5)
- **MariaDB** (Versions: 10.5, 10.4, 10.3, 10.2)
- **SQL Server** (Versions: 2019, 2017, 2016, 2014)
- **Oracle** (Versions: 19c, 18c, 12c, 11g)
- **MongoDB** (Versions: 6.0, 5.0, 4.4, 4.2)

## 📦 Docker Requirements

Ensure you have Docker and Docker Compose installed before running DB Composer.

- [Docker Installation Guide](https://docs.docker.com/get-docker/)
- [Docker Compose Installation Guide](https://docs.docker.com/compose/install/)

## 🤝 Contributing

Feel free to submit issues or contribute to the project by opening a pull request.

## 📝 License

This project is licensed under the MIT License.

