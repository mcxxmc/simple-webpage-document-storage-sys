import React from "react";
import DirComponent from "./DirComponent"
import FileComponent from "./FileComponent";

class Hierarchy extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            version: 0
        }
        this.organizedData = []
    }

    componentDidMount() {
        this.recursivelyAppend(this.props.dirs, this.props.root, 0)
        this.setState({version: this.state.version + 1})
    }

    render() {
        return (
            <div>
                {this.organizedData}
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
            if (dir["dir"] === true) {
                let children = dir["children"];
                this.organizedData.push(<DirComponent id={dir["id"]} name={dir["name"]} level={dir["level"]}/>)
                if (children.length > 0) {  // append its children
                    for (let i = 0; i < children.length; i ++) {
                        this.recursivelyAppend(dirs, children[i], dep + 1);
                    }
                }
            } else {
                this.organizedData.push(<FileComponent id={dir["id"]} name={dir["name"]} level={dir["level"]}/>)
            }
        } catch(error) {
            console.error(error);
        }
    }
}

export default Hierarchy