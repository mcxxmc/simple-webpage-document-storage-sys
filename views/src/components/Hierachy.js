import React from "react";
import Dir from "./Dir"
import File from "./File";
import FileVis from "./FileVis";
import {defaultUserId, user2url} from "../constants/constants";

class Hierarchy extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            organized: [],
            root: "",
            user: this.props.userId,
            fileOnDisplay: -1,  // the index of the file that is now displayed. -1 means none.
            fileOnDisplayId: "",  // the id of the file that is now displayed
            indexesModified: false,  // if the indexes of the organized array is modified
            filename: "",  // the filename of the file that is now displayed
            content: "",  // the content of the file that is now displayed
        }
        this.cache = []
    }

    /**
     * sort the data from the promise and update the state (the user id won't change)
     * @param {Promise} promise
     */
    sort(promise) {
        this.recursivelyAppend(promise["dirs"], promise["top"], 0);
        this.setState({
            organized: this.cache,
            root: promise["top"],
            fileOnDisplay: -1,
            fileOnDisplayId: "",
            indexesModified: false,
            filename: "",
            content: ""});
    }

    /**
     * fetch data from the backend, sort them and reset the state.
     * this function is expensive so try to avoid it when possible unless laziness got you
     */
    fetchAndSort() {
        /* fetch data from the backend */
        const url = user2url["get"][defaultUserId]
        fetch(url)
            .then(response => response.json())
            .catch(error => console.error("Error: ", error))
            .then(response => this.sort(response))
        this.cache = []  // empty the cache
    }

    componentDidMount() {
        this.fetchAndSort()
    }

    /**
     * fetch the file content from the backend
     * @param {int} i the index of the dir in this.state.organized
     */
    fetchFile(i) {
        const dir = this.state.organized[i];
        fetch(user2url["post"]["readFile"], {
            body: JSON.stringify({"user": this.state.user, "fid": dir["id"]}),
            method: 'POST'
        })
            .then(response => response.json())
            .catch(error => console.error("Error: ", error))
            .then(response => this.setState({fileOnDisplay: i, fileOnDisplayId: dir["id"],
                filename: response["file_name"], content: response["content"]}))
    }

    /**
     * The callback function for modifying file
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
                "user": this.state.user,
                "fid": this.state.organized[this.state.fileOnDisplay]["id"],
                "new_c": childData}),
            method: 'POST'
        }).catch(error => console.error("Error when modifying file: ", error))
            .then(() => alert("success"))
        this.setState({content: childData})
    }

    /**
     * The callback function for renaming a file or a directory.
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
                    "user": this.state.user,
                    "obj_id": objId,
                    "dir": isDir,
                    "new_name": newName
                }),
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
     * The callback function for creating a new file or a directory.
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
                "user": this.state.user,
                "dir": isDir,
                "name": name,
                "new_content": content,
                "parent_id": parentId
            }),
            method: 'POST'
        }).catch(error => console.error("Error when creating", error))
            .then(() => alert("success"))
        this.fetchAndSort()  //todo: change to a cheaper way
    }

    /**
     * The callback function for deleting a dir or a file.
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
        let type = "";
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
                "user": this.state.user,
                "obj_id": objId,
                "dir": isDir
            }),
            method: 'POST'
        }).catch(error => alert("Error when deleting. Probably the dir is not empty."))
            .then(() => this.fetchAndSort())  //todo: change to a cheaper way
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
                        callbackRename={this.callbackRename}
                        callbackCreate={this.callbackCreate}
                        callbackDelete={this.callbackDelete}/>
        } else {
            return <File id={x["id"]} name={x["name"]} level={x["level"]} index={i}
                         key={i} onClick={() => this.fetchFile(i)}/>
        }
    }

    render() {
        let visFile;
        if (this.state.fileOnDisplay !== -1 && !this.state.indexesModified) {
            visFile = <FileVis filename={this.state.filename} content={this.state.content}
                               index={this.state.fileOnDisplay} id={this.state.fileOnDisplayId}
                               callbackModifyFile={this.callbackModifyFile}
                               callbackRename={this.callbackRename}
                               callbackDelete={this.callbackDelete}/>
        }
        return (
            <div>
                {this.state.organized.map((x, i) => this.display(x, i))}
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