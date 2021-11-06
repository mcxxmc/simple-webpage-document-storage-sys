export const tab = "  "

export const defaultUserId = "0"

export const user2url = {
    "get": {
        "0": "http://localhost:8080/default-view/view"
    },
    "post": {
        "readFile": "http://localhost:8080/filesystem/read",
        "modifyFile": "http://localhost:8080/filesystem/rewrite",
        "rename": "http://localhost:8080/filesystem/rename",
        "create": "http://localhost:8080/filesystem/create",
        "delete": "http://localhost:8080/filesystem/remove",
        "move": "http://localhost:8080/filesystem/move"
    }
}