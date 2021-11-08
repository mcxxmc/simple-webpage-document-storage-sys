import React from "react";
import Hierarchy from "./Hierachy";
import {defaultUserId} from "../constants/constants";

class DefaultView extends React.Component {

    render() {
        return (
            <div>
                <Hierarchy userId={defaultUserId}/>
            </div>
        )
    }
}

export default DefaultView