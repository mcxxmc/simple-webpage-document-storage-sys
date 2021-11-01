import React from "react";
import "./css/FD.css";

const tab = "  "

class FileComponent extends React.Component {

    render() {
        return (
            <div className={"div_file"}>
                <pre>
                    <p>{tab.repeat(this.props.level) + this.props.name}{" "}<a>View</a></p>
                </pre>
            </div>
        )
    }
}

export default FileComponent