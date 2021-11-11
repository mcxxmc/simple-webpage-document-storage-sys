import React from "react";
import Dir from "./Dir"
import File from "./File";
import FileVis from "./FileVis";
import Login from "./Login"
import {user2url} from "../constants/constants";
import "./css/hierachy.css"

class Hierarchy extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            organized: [],
            root: "",
            fileOnDisplay: -1,  // the index of the file that is now displayed. -1 means none.
            fileOnDisplayId: "",  // the id of the file that is now displayed
            filename: "",  // the filename of the file that is now displayed
            content: "",  // the content of the file that is now displayed
            markedIndex: -1,  // the index of the marked dir
            markedId: "",  // the id of the marked dir; used to check if the markedIndex is valid
            login: false,  // whether the user has logged in
        }
        this.cache = []
    }

    /**
     * sort the data from the promise and update the state
     * @param {Promise} promise
     */
    sort(promise) {
        if (!promise["ok"]) {
            console.error("unsuccessful transaction")
            return
        }
        this.recursivelyAppend(promise["dirs"], promise["top"], 0);
        this.setState({
            organized: this.cache,
            root: promise["top"],
            fileOnDisplay: -1,
            fileOnDisplayId: "",
            filename: "",
            content: "",
            markedIndex: -1,
            markedId: "",
        });
    }

    /**
     * fetch data from the backend, sort them and reset the state.
     * this function is expensive so try to avoid it when possible unless laziness got you
     */
    fetchAndSort() {
        /* fetch data from the backend */
        fetch(user2url["get"]["view"], {
            headers: {
                "Authorization": window.localStorage.getItem("token")
            }
        })
            .then(response => response.json())
            .catch(error => console.error("Error: ", error))
            .then(response => this.sort(response))
        this.cache = []  // empty the cache
    }

    /**
     * fetch the file content from the backend
     * @param {int} i the index of the dir in this.state.organized
     */
    fetchFile(i) {
        const dir = this.state.organized[i];
        fetch(user2url["post"]["readFile"], {
            body: JSON.stringify({"fid": dir["id"]}),
            headers: {
                "Authorization": window.localStorage.getItem("token")
            },
            method: 'POST'
        })
            .then(response => response.json())
            .catch(error => console.error("Error: ", error))
            .then(response => this.setState({fileOnDisplay: i, fileOnDisplayId: dir["id"],
                filename: response["file_name"], content: response["content"]}))
    }

    /**
     * stopping to display the file and return to the tree view.
     */
    callbackStopDisplayingFile = () => {
        this.setState({fileOnDisplay: -1});
    }

    /**
     * modifying file
     * The file modified will be the current file on display
     * @param {string} childData
     */
    callbackModifyFile = (childData) => {  //todo: better structured childData
        if (this.state.fileOnDisplay === -1) {
            console.error("modifyFile error: no file selected")
            return
        }
        fetch(user2url["post"]["modifyFile"], {
            body: JSON.stringify({
                "fid": this.state.organized[this.state.fileOnDisplay]["id"],
                "new_c": childData}),
            headers: {
                "Authorization": window.localStorage.getItem("token")
            },
            method: 'POST'
        }).catch(error => console.error("Error when modifying file: ", error))
            .then(response => response.json())
            .then(response => {
                alert(response["msg"]);
                if (response["ok"]) {
                    this.setState({content: childData})
                }
            })
    }

    /**
     * renaming a file or a directory.
     * The file renamed will be the current file on display.
     * @param childData
     */
    callbackRename = (childData) => {
        let index = childData["index"]
        let objId = childData["objId"]
        let isDir = childData["isDir"]
        let newName = childData["newName"]

        let tmp = this.state.organized[index];
        let type = "";

        if (tmp["dir"] !== isDir) {
            console.error("type does not match")
            return
        }
        if (tmp["id"] !== objId) {
            console.error("object id does not match")
            return
        }
        if (tmp["name"] === newName) {
            console.log("renaming has no effect")
            return
        }

        // for a directory
        if (tmp["dir"]) {
            type = " directory "
        }
        // for a file
        else {
            if (this.state.fileOnDisplay !== index) {
                console.error("wrong index for file")
                return
            }
            type = " file "
        }

        let msg = "Are you sure to rename" + type + tmp["name"] + " to " + newName + "?";

        if (window.confirm(msg)) {
            fetch(user2url["post"]["rename"], {
                body: JSON.stringify({
                    "obj_id": objId,
                    "dir": isDir,
                    "new_name": newName
                }),
                headers: {
                    "Authorization": window.localStorage.getItem("token")
                },
                method: 'POST'
            }).catch(error => console.error("Error when renaming", error))
                .then(() => alert("success"))
        }
        tmp = this.state.organized;
        tmp[index]["name"] = newName;
        this.setState({
            organized: tmp
        })
        if (!isDir) {
            this.setState({filename: newName})
        }
    }

    /**
     * creating a new file or a directory.
     * @param childData
     */
    callbackCreate = (childData) => {
        let parentIndex = childData["parentIndex"];
        let parentId = childData["parentId"];
        let isDir = childData["isDir"];
        let name = childData["name"];
        let content = "";  // only for new file; always initialized as ""
        if (this.state.organized[parentIndex]["id"] !== parentId) {
            console.error("Error when creating: index and id does not match")
            return
        }
        fetch(user2url["post"]["create"], {
            body: JSON.stringify({
                "dir": isDir,
                "name": name,
                "new_content": content,
                "parent_id": parentId
            }),
            headers: {
                "Authorization": window.localStorage.getItem("token")
            },
            method: 'POST'
        }).catch(error => console.error("Error when creating", error))
            .then(() => alert("success"))
        this.fetchAndSort()  //todo: change to a cheaper way
    }

    /**
     * deleting a dir or a file.
     * The dir to be deleted must have been empty.
     * @param childData
     */
    callbackDelete = (childData) => {
        let name = childData["name"];
        let objId = childData["objId"];
        let isDir = childData["isDir"];
        let index = childData["index"];
        if (this.state.organized[index]["id"] !== objId || this.state.organized[index]["dir"] !== isDir) {
            console.error("Error when deleting: id or type does not match")
            return
        }
        let type;
        if (isDir) {
            type = " directory "
        } else {
            type = " file "
        }
        const msg = "Are you sure to delete the" + type + name + "? Note that a directory cannot be deleted if it is not empty."
        if (!window.confirm(msg)) {
            return
        }
        fetch(user2url["post"]["delete"], {
            body: JSON.stringify({
                "obj_id": objId,
                "dir": isDir
            }),
            headers: {
                "Authorization": window.localStorage.getItem("token")
            },
            method: 'POST'
        }).catch(error => alert("Error when deleting. Probably the dir is not empty."))
            .then(() => this.fetchAndSort())  //todo: change to a cheaper way
    }

    /**
     * marking a dir as the potential new parent node.
     * @param childData
     */
    callbackMark = (childData) => {
        let index = childData["index"];
        let id = childData["id"];
        if (this.state.organized[index]["id"] !== id) {
            console.error("Error when marking: index and id does not match")
            alert("error")
            return
        }
        if (this.state.markedIndex === index && this.state.markedId === id) {
            console.log("the directory is already marked. It is now unmarked.")
            this.setState({markedIndex: -1, markedId: ""})
            return
        }
        this.setState({markedIndex: index, markedId: id})
    }

    /**
     * moving a dir or a file.
     * @param childData
     */
    callbackMove = (childData) => {
        if (this.state.markedIndex === -1 || this.state.markedId === "") {
            alert("No parent node selected yet.")
            return
        }
        let objId = childData["objId"];
        let isDir = childData["isDir"];
        let index = childData["index"];
        if (this.state.organized[this.state.markedIndex]["id"] !== this.state.markedId) {
            console.error("Error when marking: index and id do not match")
            alert("error")
            return
        }
        if (this.state.organized[index]["id"] !== objId || this.state.organized[index]["dir"] !== isDir) {
            console.error("Error when Moving: child index and id do not match, or the types do not match")
            alert("error")
            return
        }
        let type;
        if (isDir) {
            type = " directory "
        } else {
            type = " file "
        }
        let msg = "are you sure you want to move" + type + this.state.organized[index]["name"] + "?";
        if (window.confirm(msg)) {
            fetch(user2url["post"]["move"], {
                body: JSON.stringify({
                    "obj_id": objId,
                    "dir": isDir,
                    "new_parent_id": this.state.markedId
                }),
                headers: {
                    "Authorization": window.localStorage.getItem("token")
                },
                method: 'POST'
            }).catch(error => alert("Error moving"))
                .then(() => this.fetchAndSort())  //todo: change to a cheaper way
        }
    }

    /**
     * logging in (callback part)
     * @param {JSON} childData
     */
    callbackLogin = (childData) => {
        let name = childData["name"];
        let pwd = childData["pwd"];
        fetch(user2url["post"]["login"], {
            body: JSON.stringify({
                "username": name,
                "password": pwd
            }),
            method: 'POST'
        }).catch(error => console.error(error))
            .then(response => response.json())
            .then(response => this.login(response))
    }

    /**
     * logging in (continuing the callback part)
     * @param {JSON} response
     */
    login(response) {
        if (response["ok"]) {
            window.localStorage.setItem("token", response["token"])
            this.fetchAndSort()
            this.setState({login: true})
        } else {
            alert("Fail to login. Please check your username and password.")
        }
    }

    /**
     * logging out
     */
    logout() {
        let tk =  window.localStorage.getItem("token");
        if (tk === null) {
            alert("you have not logged in yet")
            return
        }

        fetch(user2url["get"]["logout"], {
            headers: {
                "Authorization": tk
            },
            method: 'GET'
        }).catch(error => console.error(error))
            .then(response => response.json())
            .then(response => alert(response["msg"]))

        window.localStorage.removeItem("token");
        this.setState({login: false})
    }

    /**
     * used for the render() method
     * @param {JSON} x
     * @param {int} i
     */
    display(x, i) {
        if (x["dir"] === true) {
            return <Dir id={x["id"]} name={x["name"]} level={x["level"]} index={i}
                        key={i}
                        style={{backgroundColor:
                                this.state.markedIndex === i && this.state.markedId === x["id"]? "red": "transparent"}}
                        callbackRename={this.callbackRename}
                        callbackCreate={this.callbackCreate}
                        callbackDelete={this.callbackDelete}
                        callbackMark={this.callbackMark}
                        callbackMove={this.callbackMove}/>
        } else {
            return <File id={x["id"]} name={x["name"]} level={x["level"]} index={i}
                         key={i} onClick={() => this.fetchFile(i)}
                         callbackMove={this.callbackMove}/>
        }
    }



    render() {
        let logout;
        let login;
        let visFile;
        let tree;

        logout = <button className={"basic-btn"} onClick={() => this.logout()}>Logout</button>

        if (!this.state.login) {
            login = <Login callbackLogin={this.callbackLogin}/>
        } else {
            //this.fetchAndSort()
            if (this.state.fileOnDisplay !== -1) {
                visFile = <FileVis filename={this.state.filename} content={this.state.content}
                                   index={this.state.fileOnDisplay} id={this.state.fileOnDisplayId}
                                   callbackStopDisplayingFile={this.callbackStopDisplayingFile}
                                   callbackModifyFile={this.callbackModifyFile}
                                   callbackRename={this.callbackRename}
                                   callbackDelete={this.callbackDelete}/>
            }

            if (this.state.fileOnDisplay === -1) {
                tree = this.state.organized.map((x, i) => this.display(x, i));
            }
        }

        return (
            <div className={"div-hierarchy"}>
                {logout}
                {login}
                {tree}
                {visFile}
            </div>
        )
    }

    /**
     * append the whole structure
     * @param {JSON} dirs
     * @param {string} cur the current directory id to append
     * @param {number} dep the current depth of recursion
     * @param {number} maxDep the max depth of recursion
     */
    recursivelyAppend(dirs, cur, dep, maxDep=24) {
        if (dep >= maxDep) {
            return
        }
        try {
            let dir = dirs[cur];
            this.cache.push(dir)
            if (dir["dir"] === true) {  // if it is a directory (not a file)
                let children = dir["children"];
                if (children.length > 0) {  // append its children
                    for (let i = 0; i < children.length; i ++) {
                        this.recursivelyAppend(dirs, children[i], dep + 1);
                    }
                }
            }
        } catch(error) {
            console.error(error);
        }
    }
}

export default Hierarchy