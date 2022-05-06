# Emotebox
Emotebox is a web application for emote management in Discord.

## Features
- Easily search through emotes with standard filters
- Upload custom emotes for others to use


## Tech
- [React.js](https://reactjs.org/) - A JavaScript library for building user interfaces
- [Golang](https://go.dev/) - Compiled programming language designed by Google
- [Gin](https://gin-gonic.com/) - A full-featured web framework for go
- [Vite](https://vitejs.dev/) - The Faster solution to Webpack
- [PostgresSQL](https://www.postgresql.org/) - Reliable relation database

## Installation
Installing the project is as easy as cloning the repository using the git command
`git clone https://github.com/yeno-team/emotebox.git`

## Configuration
This project uses env files to configure the application.
> Env files can be made for different environments by naming it `.{environmentName}.env`
> Default environment is development, and the file name would be `.development.env`

| Name    | Description        |
|---------|--------------------|
| APP_ENV | Target environment (default: *development*) |
| DISCORD_TOKEN |   A Discord Application's bot token   |


## Contributing
If you want to work on an issue/feature, make sure to open an issue before committing too far otherwise small bug fixes are always welcome.

## Development

The project uses docker compose to easily setup a dev environment, just run the command for a live reload dev environment.
```
docker compose up -d
```

The server is exposed at `localhost:8080` while the client is at `localhost:4000`.
Happy developing!

## License
MIT
