import React from "react";
import Dir from "./Dir"
import File from "./File";
import ReactDOM from "react-dom";
import FileVis from "./FileVis";
import {user2url} from "../constants/constants";

class Hierarchy extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            organized: [],
            user: this.props.userId
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
        fetch(user2url["post"]["readFile"][this.state.user], {  //TODO user
            body: JSON.stringify({"user": this.state.user, "fid": dir["id"]}),
            method: 'POST'
        })
            .then(response => response.json())
            .catch(error => console.error("Error: ", error))
            .then(response => visFile(response))
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
        return (
            <div>
                {this.state.organized.map((x, i) => this.display(x, i))}
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

/**
 * renders the file visualization
 * @param {JSON} response
 */
function visFile(response) {
    ReactDOM.render(
        <FileVis fileName={response["file_name"]} content={response["content"]}/>,
        document.getElementById("div_txt")
    )
}

export default Hierarchy