# Go Template

## About

This is my personal template when creating a new Golang Project. It is highly opinionated, so please bear with me at how things are structured - as it may be (or is) highly unconventional.

### How to run/build the code

Simply run `go run server.go` as it is the entrypoint of the code itself. When you're ready to build the code, simply run `go build` and a binary executable will be created for you.

### Any dependencies needed to be installed?

You do not have to install anything but Golang itself - all dependencies are handled by the go modules (see `go.mod` and `go.sum`).

**Important:** When you're tweaking your own version of this template, ensure that you **do not** edit `go.sum` directly - it should be updated by go modules through the `go.mod`!

### How to test if the template is working?

Here are the steps you might want to do to ensure the template is working as intended:
1. Clone this repository (or download it)
2. Copy `.env.example` to `.env`
3. Run `go server.go`
4. Using your browser, access http://localhost:4000/api (or whatever you define the running port in `APP_PORT`).
5. If you receive the following json:
```json
{
  "data": {
    "app_description": "Public API Endpoint for Clevecord's Experimental Transcoding Service",
    "app_version": "0.0.1",
    "app_environment": "development"
  }
}
```
Congratulations! The server is running properly, and you're free to code!

6. (Optional) Test out the `/env` route: simply access http://localhost:4000/api/env/APP_DESCRIPTION for example. It should output something along the lines of:
```json
{
  "data": {
    "key": "APP_DESCRIPTION",
    "value": "Public API Endpoint for Clevecord's Experimental Transcoding Service"
  }
}
```
7. Once you've ensured that the code is running properly, you *should* remove all environment routes alongside its related files (as they are sample codes that should *never* reach production). Here are the list of directories you will need to remove (using the root folder as the base relative URL):
```
server/apimodels/env
server/controllers/env
server/repository/env
server/services/env
```
Don't forget to remove the following lines from `server/routes/root.go` as well:
```go
// Should be line 7
envController "template.clevecord.me/server/controllers/env"

// Should be lines 19 - 21
// Env Routes
e.GET(APIPrefix+EnvPrefix+"/", envController.GetEnv)
e.GET(APIPrefix+EnvPrefix+"/:env", envController.GetEnv)
```

Happy coding!

### About Environment Variables

Environment variables are read through the `godotenv` package. It reads a `.env` file that *must* be located at the same place you execute the command to run or build. For a binary file, you just need to have the executable to be at the same place as the executable.

This repository comes with a environment variable endpoint as the sample code (don't forget to remove it, see the previous point on how to); feel free to use it to test the functionality of the template.

## How to use the template/what to do?

Feel free to check the wiki to understand how to do a few things such as:
1. Adding routes
2. Returning values using the ResponseWrapper
