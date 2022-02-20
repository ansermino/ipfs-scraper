import React, { Component } from 'react'
import { Grid, Menu, Segment } from 'semantic-ui-react'
import {BrowserRouter as Router, NavLink, Route, Routes, withRouter} from 'react-router-dom'

export default class Sidebar extends Component {
    state = { activeItem: 'home' }

    handleItemClick = (e, { name }) => this.setState({ activeItem: name })

    render() {
        const { activeItem } = this.state

        return (

                <Grid.Column streched width={4}>
                    <Menu fluid vertical tabular>
                        <Menu.Item
                            as={NavLink}
                            to={"/"}
                            name='home'
                            active={activeItem === 'home'}
                            onClick={this.handleItemClick}
                        />
                        <Menu.Item
                            as={NavLink}
                            to={"/add"}
                            name='add'
                            active={activeItem === 'add'}
                            onClick={this.handleItemClick}
                        />
                        <Menu.Item
                            as={NavLink}
                            to={"/view"}
                            name='view'
                            active={activeItem === 'view'}
                            onClick={this.handleItemClick}
                        />
                    </Menu>
                </Grid.Column>

        )
    }
}
