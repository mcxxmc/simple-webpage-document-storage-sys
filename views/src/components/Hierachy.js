import React from "react";
import Dir from "./Dir"
import File from "./File";
import FileVis from "./FileVis";
import {user2url} from "../constants/constants";

class Hierarchy extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            organized: [],
            user: this.props.userId,
            fileOnDisplay: -1,  // the index of the file that is now displayed. -1 means none.
            filename: "",  // the filename of the file that is now displayed
            content: "",  // the content of the file that is now displayed
        }
        this.cache = []
    }

    componentDidMount() {
        this.recursivelyAppend(this.props.dirs, this.props.root, 0)
        this.setState({organized: this.cache})
        this.cache = []
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
            .then(response => this.setState({fileOnDisplay: i,
                filename: response["file_name"], content: response["content"]}))
    }

    /**
     * The callback function for modifying file
     * The file modified will be the current file on display
     * @param {string} childData
     */
    callbackModifyFile = (childData) => {
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
    }

    /**
     * The callback function for renaming a file.
     * The file renamed will be the current file on display
     * @param childData
     */
    callbackRenameFile = (childData) => {
        if (this.state.fileOnDisplay === -1) {
            console.error("renameFile error: no file selected")
            return
        }
        let tmp = this.state.organized[this.state.fileOnDisplay];
        if (tmp["dir"]) {
            console.error(("renameFile error: not a file"))
            return
        }
        let msg = "Are you sure to rename file " + tmp["name"] + " to " + childData + "?";
        if (window.confirm(msg)) {
            fetch(user2url["post"]["rename"], {
                body: JSON.stringify({
                    "user": this.state.user,
                    "obj_id": tmp["id"],
                    "dir": tmp["dir"],
                    "new_name": childData
                }),
                method: 'POST'
            }).catch(error => console.error("Error when renaming file", error))
                .then(() => alert("success"))
        }
        tmp = this.state.organized;
        tmp[this.state.fileOnDisplay]["name"] = childData;
        this.setState({
            organized: tmp, filename: childData
        })
    }

    /**
     * used for the render() method
     * @param {JSON} x
     * @param {int} i
     */
    display(x, i) {
        if (x["dir"] === true) {
            return <Dir id={x["id"]} name={x["name"]} level={x["level"]}
                        key={"d_" + i}/>
        } else {
            return <File id={x["id"]} name={x["name"]} level={x["level"]}
                         key={"f_" + i} onClick={() => this.fetchFile(i)}/>
        }
    }

    render() {
        let visFile;
        if (this.state.fileOnDisplay !== -1) {
            visFile = <FileVis filename={this.state.filename} content={this.state.content}
                               callbackModifyFile={this.callbackModifyFile}
                               callbackRenameFile={this.callbackRenameFile}/>
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