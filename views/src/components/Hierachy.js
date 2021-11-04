import React from "react";
import Dir from "./Dir"
import File from "./File";

class Hierarchy extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            organized: []
        }
        this.cache = []
    }

    componentDidMount() {
        this.recursivelyAppend(this.props.dirs, this.props.root, 0)
        this.setState({organized: this.cache})
        this.cache = []
    }

    render() {
        return (
            <div>
                {this.state.organized}
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
                this.cache.push(
                    <Dir id={dir["id"]} name={dir["name"]} level={dir["level"]} userId={this.props.userId}
                         key={"key" + dir["id"]}/>
                )
                if (children.length > 0) {  // append its children
                    for (let i = 0; i < children.length; i ++) {
                        this.recursivelyAppend(dirs, children[i], dep + 1);
                    }
                }
            } else {
                this.cache.push(
                    <File id={dir["id"]} name={dir["name"]} level={dir["level"]} userId={this.props.userId}
                          key={"key" + dir["id"]}/>
                )
            }
        } catch(error) {
            console.error(error);
        }
    }
}

export default Hierarchy