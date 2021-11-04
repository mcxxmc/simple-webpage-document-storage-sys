import React from "react";
import "./css/FD.css";
import {defaultUserId, user2url, tab} from "../constants/constants";
import ReactDOM from "react-dom";
import FileVis from "./FileVis";

class File extends React.Component {

    render() {
        return (
            <div className={"div_file"}>
                <pre>
                    <p>{tab.repeat(this.props.level) + this.props.name}{" "}
                        {/* eslint-disable-next-line no-script-url */}
                        <a href={"javascript:void(0)"} className={"a_div_file_view"}
                           onClick={() => fetchFile(this.props.userId, this.props.id)}>View</a>
                    </p>
                </pre>
            </div>
        )
    }
}

/**
 * fetch the file content from the backend
 * @param {string} userId the user id
 * @param {string} fileId the file id
 */
function fetchFile(userId, fileId) {
    fetch(user2url["post"]["readFile"][defaultUserId], {
        body: JSON.stringify({"user": userId, "fid": fileId}),
        method: 'POST'
    })
        .then(response => response.json())
        .catch(error => console.error("Error: ", error))
        .then(response => visFile(response))
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

export default File