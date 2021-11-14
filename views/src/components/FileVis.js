import React from "react";
import "./css/file-vis.css"
import {checkValidInput} from "../constants/validation";

const maxStackCapacity = 20;  // the max length of this.history

class FileVis extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            name: this.props.filename,
            id: this.props.id,
            index: this.props.index,
            content: this.props.content,
            rename: false,
            newName: this.props.filename,
        }
        this.handleChangeCtt = this.handleChangeCtt.bind(this)
        this.handleChangeRename = this.handleChangeRename.bind(this)
        // basically, undo and redo are 2 stacks (but when they are full they might behaviour like queues)
        this.undo = []  // every time the content is updated, an old version is added to undo
        this.redo = []  // every time undo pops something, the replaced version of content is added to redo
    }

    // does not need componentDidUpdate() here because this component will be destroyed when quitting

    // used to update this.undo & this.redo
    currentHistory() {
        if (this.state.content && this.state.content.length > 0) {
            return this.state.content.slice(0)  // make a copy
        } else {
            return ""
        }
    }

    /**
     * handles the change for modifying context
     * @param event
     */
    handleChangeCtt(event) {
        const copy = this.currentHistory()
        if (this.undo.length === maxStackCapacity) {
            this.undo = this.undo.slice(1)  // abandon the oldest version; this behaviour is like a queue
        }
        this.undo.push(copy)
        this.setState({content: event.target.value})
    }

    /**
     * handles the change for renaming
     * @param event
     */
    handleChangeRename(event) {
        this.setState({newName: event.target.value})
    }

    render() {
        let rename;
        if (this.state.rename) {
            rename = (
                <div>
                    <textarea value={this.state.newName} onChange={this.handleChangeRename}
                              className={"file-vis-textarea-rename"}/>
                    <button onClick={() => this.rename()}
                            className={"basic-btn button-confirm"}>Confirm</button>
                    <button onClick={() => this.setState({rename: false})}
                            className={"basic-btn button-cancel"}>Cancel</button>
                </div>
            )
        }

        return (
            <div>
                <button onClick={() => this.quit()}
                        className={"basic-btn button-cancel"}>Quit</button>
                <h1 className={"file-vis-h1"}>{this.state.name}</h1>
                <button onClick={() => this.setState({rename: true})}
                        className={"basic-btn button-rename"}>Rename</button>
                <button onClick={() => this.delete()}
                        className={"basic-btn button-delete"}>Delete</button>
                {rename}
                <br/>
                <textarea value={this.state.content} onChange={this.handleChangeCtt} className={"file-vis-textarea"}/>
                <button onClick={() => this.undo_f()} disabled={this.undo.length <= 0}
                    className={"small-basic-btn"}>Undo</button>
                <button onClick={() => this.redo_f()} disabled={this.redo.length <= 0}
                    className={"small-basic-btn"}>Redo</button>
                <button onClick={() => this.commit()}
                        className={"basic-btn button-confirm"}>Save</button>
            </div>
        )
    }

    commit() {
        this.props.callbackModifyFile(this.state.content)
    }

    rename() {
        let name = this.state.newName.trim();
        if (checkValidInput(name)) {
            this.props.callbackRename({"index": this.state.index, "objId": this.state.id, "isDir": false,
                "newName": name})
            this.setState({name: this.state.newName, rename: false})
        }
    }

    delete() {
        this.props.callbackDelete({"name": this.state.name, "objId": this.state.id, "isDir": false, "index": this.state.index})
    }

    quit() {
        this.props.callbackStopDisplayingFile({"quit": true})
    }

    undo_f() {
        if (this.undo.length === 0) {
            return
        }
        const pop = this.undo.pop()
        const copy = this.currentHistory()
        if (this.redo.length === maxStackCapacity) {
            this.redo = this.redo.slice(1)
        }
        this.redo.push(copy)
        this.setState({content: pop})
    }

    redo_f() {
        if (this.redo.length === 0) {
            return
        }
        const pop = this.redo.pop()
        const copy = this.currentHistory()
        if (this.undo.length === maxStackCapacity) {
            this.undo = this.undo.slice(1)
        }
        this.undo.push(copy)
        this.setState({content: pop})
    }
}

export default FileVis