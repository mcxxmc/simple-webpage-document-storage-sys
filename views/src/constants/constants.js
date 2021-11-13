export const tab = "    "
export const bearer = "Bearer "

export const user2url = {
    "get": {
        "view": "http://localhost:8080/filesystem/view",
        "logout": "http://localhost:8080/filesystem/logout"
    },
    "post": {
        "readFile": "http://localhost:8080/filesystem/read",
        "modifyFile": "http://localhost:8080/filesystem/rewrite",
        "rename": "http://localhost:8080/filesystem/rename",
        "create": "http://localhost:8080/filesystem/create",
        "delete": "http://localhost:8080/filesystem/remove",
        "move": "http://localhost:8080/filesystem/move",
        "login": "http://localhost:8080/filesystem/login",
        "register": "http://localhost:8080/filesystem/register",
    }
}