import React, {Component} from 'react'
import {Route, Routes} from "react-router-dom";

import ViewAll from "./ViewAll";
import ViewVersions from "./ViewVersions";


export default class View extends Component {

    constructor(props) {
        super(props);
        this.state = {pages: [], selectedPage: {}}
    }

    render() {
        return (
            <Routes>

                <Route
                    path="/:title"
                    exact
                    element={<ViewVersions />}
                />
                <Route
                    path="/"
                    element={<ViewAll />}
                />
                <Route
                    element={<h2>Not found</h2>}
                />
            </Routes>
        )
    }
}
