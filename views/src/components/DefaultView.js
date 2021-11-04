import React from "react";
import Hierarchy from "./Hierachy";
import {defaultUserId, user2url} from "../constants/constants";

class DefaultView extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            dirs: [],
            root: "",
            fetched: false
        }
    }

    componentDidMount() {
        /* fetch data from the backend */
        const url = user2url["get"][defaultUserId]
        console.log(url)
        fetch(url)
            .then(response => response.json())
            .catch(error => console.error("Error: ", error))
            .then(response => this.setState({dirs: response["dirs"], root: response["top"], fetched: true}))
    }

    render() {
        switch (this.state.fetched) {
            case true:
                return (
                    <div>
                        <Hierarchy dirs={this.state.dirs} root={this.state.root} userId={defaultUserId}/>
                    </div>
                );
            default:
                return (
                    <div>
                        <p>loading</p>
                    </div>
                )
        }
    }

}

export default DefaultView