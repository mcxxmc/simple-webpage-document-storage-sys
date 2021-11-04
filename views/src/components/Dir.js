import React from "react";
import {tab} from "../constants/constants";

class Dir extends React.Component {

    render() {
        return (
            <div className={"div_dir"}>
                <pre>
                    <p>{tab.repeat(this.props.level) + this.props.name}</p>
                </pre>
            </div>
        )
    }

}

export default Dir