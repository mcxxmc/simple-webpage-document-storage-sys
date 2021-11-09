import React from "react";
import "./css/file-vis.css"

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
    }

    // check if there is a change in the state (due to changes from the parent)
    componentDidUpdate(prevProps, prevState, snapshot) {
        if (prevState.name !== this.props.filename) {
            this.setState({
                name: this.props.filename,
                id: this.props.id,
                index: this.props.index,
                content: this.props.content,
                rename: false,
                newName: this.props.filename})
        }
    }

    /**
     * handles the change for modifying context
     * @param event
     */
    handleChangeCtt(event) {
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
                    <button onClick={() => this.rename()}>Confirm</button>
                    <button onClick={() => this.setState({rename: false})}>Cancel</button>
                </div>
            )
        }

        return (
            <div>
                <button onClick={() => this.quit()}>Quit</button>
                <h1 className={"file-vis-h1"}>{this.state.name}</h1>
                <button onClick={() => this.setState({rename: true})}>Rename</button>
                <button onClick={() => this.delete()}>Delete</button>
                {rename}
                <br/>
                <textarea value={this.state.content} onChange={this.handleChangeCtt} className={"file-vis-textarea"}/>
                <button onClick={() => this.commit()}>Commit</button>
            </div>
        )
    }

    commit() {
        this.props.callbackModifyFile(this.state.content)
    }

    rename() {
        this.props.callbackRename({"index": this.state.index, "objId": this.state.id, "isDir": false,
            "newName": this.state.newName.trim()})
        this.setState({rename: false})
    }

    delete() {
        this.props.callbackDelete({"name": this.state.name, "objId": this.state.id, "isDir": false, "index": this.state.index})
    }

    quit() {
        this.props.callbackStopDisplayingFile({"quit": true})
    }
}

export default FileVis