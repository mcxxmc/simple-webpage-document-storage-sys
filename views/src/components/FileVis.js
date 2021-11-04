import React from "react";

class FileVis extends React.Component {

    render() {
        return (
            <div>
                <h1>{this.props.fileName}</h1>
                <p>{this.props.content}</p>
            </div>
        )
    }
}

export default FileVis