import React from "react";
import {tab} from "../constants/constants";

class File extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            index: this.props.index,
            id: this.props.id,
        }
    }

    move() {
        this.props.callbackMove({"objId": this.state.id, "isDir": false, "index": this.state.index})
    }



    render() {
        return (
            <div className={"div-flex"}>
                <input readOnly={true} className={"input-readOnly"}
                       value={tab.repeat(this.props.level) + this.props.name}/>
                {/* eslint-disable-next-line no-script-url */}
                <a href={"javascript:void(0)"} className={"a-onClick"}
                   onClick={() => this.props.onClick()}>View</a>
                <button className={"small-basic-btn button-move"}
                        onClick={() => this.move()}>Move</button>
            </div>
        )
    }
}

export default File