import React from "react";
import {tab} from "../constants/constants";
import "./css/dir.css"

class Dir extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            name: this.props.name,
            level: this.props.level,
            index: this.props.index,
            id: this.props.id,
            rename: false,
            newName: this.props.name,
            create: false,  // if the user is creating a new dir or file
            createName: "new",  // the name of the new dir or file to be created
            createDir: true,  // if the newly-created object is a dir
            show: true,
        }
        this.handleChangeRename = this.handleChangeRename.bind(this);
        this.handleChangeCreateName = this.handleChangeCreateName.bind(this);
        this.handleChangeCreateType = this.handleChangeCreateType.bind(this);
    }

    componentDidUpdate(prevProps, prevState, snapshot) {
        if (prevState.name !== this.props.name) {
            this.setState({
                name: this.props.name,
                level: this.props.level,
                index: this.props.index,
                id: this.props.id,
                rename: false,
                newName: this.props.name,
                create: false,
                createName: "new",
                createDir: true,
                show: true,
            })
        }
    }

    handleChangeRename(event) {
        this.setState({newName: event.target.value})
    }

    handleChangeCreateName(event) {
        this.setState({createName: event.target.value})
    }

    handleChangeCreateType(event) {
        this.setState({createDir: event.target.value})
    }

    rename() {
        this.props.callbackRename({"index": this.state.index, "objId": this.state.id, "isDir": true,
            "newName": this.state.newName.trim()})
        this.setState({rename: false})
    }

    create() {
        let name = this.state.createName.trim();
        if (!this.state.createDir) {
            name = name + ".txt"
        }
        this.props.callbackCreate({"parentIndex": this.state.index, "parentId": this.state.id,
            "isDir": this.state.createDir, "name": name})
        this.setState({create: false, createName: "new", createDir: true})
    }

    delete() {
        this.props.callbackDelete({"name": this.state.name, "objId": this.state.id, "isDir": true, "index": this.state.index})
    }

    render() {
        if (!this.state.show) {
            return
        }
        let rename;
        if (this.state.rename) {
            rename = (
                <div>
                    <textarea value={this.state.newName} onChange={this.handleChangeRename}
                              className={"dir-textarea-rename"} />
                    <button onClick={() => this.rename()}>Confirm</button>
                    <button onClick={() => this.setState({rename: false})}>Cancel</button>
                </div>
            )
        }
        let create;
        if (this.state.create) {
            create = (
                <div>
                    <select onChange={this.handleChangeCreateType}>
                        <option value={true}>Directory</option>
                        <option value={false}>txt file</option>
                    </select>
                    <textarea value={this.state.createName} onChange={this.handleChangeCreateName}
                              className={"dir-textarea-create"}/>
                    <button onClick={() => this.create()}>Confirm</button>
                    <button onClick={() => this.setState({create: false, createName: "new", createDir: true})}>Cancel</button>
                </div>
            )
        }
        return (
            <div className={"div_dir"}>
                <pre>
                    <p>{tab.repeat(this.state.level) + this.state.name}</p>
                </pre>
                <button onClick={() => this.setState({rename: true, create: false})}
                        className={"dir-button1"}>Rename</button>
                <button onClick={() => this.setState({rename: false, create: true})}
                        className={"dir-button1"}>New</button>
                <button onClick={() => this.delete()}
                        className={"dir-button1"}>delete</button>
                {rename}
                {create}
            </div>
        )
    }

}

export default Dir