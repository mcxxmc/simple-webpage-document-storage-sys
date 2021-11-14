import React from "react";
import {tab} from "../constants/constants";
import {checkValidInput} from "../constants/validation";
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
            show: this.props.show,
            collapse: false,
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
                show: this.props.show,
                collapse: false
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
        if (event.target.value === "true" || event.target.value === "false") {
            this.setState({createDir: JSON.parse(event.target.value)})
        } else {
            this.setState({createDir: "undefined"})
        }
    }

    rename() {
        if (checkValidInput(this.state.newName)) {
            this.props.callbackRename({"index": this.state.index, "objId": this.state.id, "isDir": true,
                "newName": this.state.newName.trim()})
            this.setState({rename: false})
        }
    }

    create() {
        let name = this.state.createName.trim();
        if (checkValidInput(name)) {
            this.props.callbackCreate({"parentIndex": this.state.index, "parentId": this.state.id,
                "isDir": this.state.createDir, "name": name})
            this.setState({create: false, createName: "new", createDir: true})
        }
    }

    delete() {
        this.props.callbackDelete({"name": this.state.name, "objId": this.state.id, "isDir": true, "index": this.state.index})
    }

    mark() {
        this.props.callbackMark({"index": this.state.index, "id": this.state.id})
    }

    move() {
        this.props.callbackMove({"objId": this.state.id, "isDir": true, "index": this.state.index})
    }

    collapse() {

    }

    render() {
        if (!this.state.show) {
            return
        }
        let collapse = (
            <button className={"btn-collapse"} onClick={() => this.collapse()}>
                {this.state.collapse? "^": ">"}</button>
        );
        let rename;
        if (this.state.rename) {
            rename = (
                <div className={"div-dir-1 div-flex"}>
                    <textarea value={this.state.newName} onChange={this.handleChangeRename}
                              className={"dir-textarea-1"} />
                    <button onClick={() => this.rename()}
                            className={"basic-btn button-confirm"}>Confirm</button>
                    <button onClick={() => this.setState({rename: false})}
                            className={"basic-btn button-cancel"}>Cancel</button>
                </div>
            )
        }
        let create;
        if (this.state.create) {
            create = (
                <div className={"div-dir-1 div-flex"}>
                    <select onChange={this.handleChangeCreateType}>
                        <option value={true}>Directory</option>
                        <option value={false}>txt file</option>
                    </select>
                    <textarea value={this.state.createName} onChange={this.handleChangeCreateName}
                              className={"dir-textarea-1"}/>
                    <button onClick={() => this.create()}
                            className={"basic-btn button-confirm"}>Confirm</button>
                    <button onClick={() => this.setState({create: false, createName: "new", createDir: true})}
                            className={"basic-btn button-cancel"}>Cancel</button>
                </div>
            )
        }
        /* for input, using props instead of state to allow it to update dynamically */
        return (
            <div className={"div-dir-2"}>
                <div className={"div-flex"}>
                    {collapse}
                    <input readOnly={true} className={"input-readOnly"} style={this.props.style}
                           value={tab.repeat(this.props.level) + this.props.name}/>
                    <button className={"small-basic-btn button-mark"}
                            onClick={() => this.mark()}>Mark</button>
                    <button className={"small-basic-btn button-move"}
                            onClick={() => this.move()}>Move</button>
                    <button onClick={() => this.setState({rename: true, create: false})}
                            className={"basic-btn button-rename"}>Rename</button>
                    <button onClick={() => this.setState({rename: false, create: true})}
                            className={"basic-btn button-create"}>New</button>
                    <button onClick={() => this.delete()}
                            className={"basic-btn button-delete"}>delete</button>
                </div>
                {rename}
                {create}
            </div>
        )
    }

}

export default Dir