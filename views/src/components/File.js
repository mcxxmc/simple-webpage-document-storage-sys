import React from "react";
import "./css/FD.css";
import {tab} from "../constants/constants";

class File extends React.Component {

    render() {
        return (
            <div className={"div_file"}>
                <pre>
                    <p>{tab.repeat(this.props.level) + this.props.name}{" "}
                        {/* eslint-disable-next-line no-script-url */}
                        <a href={"javascript:void(0)"} className={"a_div_file_view"}
                           onClick={() => this.props.onClick()}>View</a>
                    </p>
                </pre>
            </div>
        )
    }
}

export default File