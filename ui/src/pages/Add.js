import React, {Component, useEffect, useState} from 'react'
import {Button, Card, Form, Header, Input, Message} from 'semantic-ui-react'

export default class Add extends Component {
    constructor(props) {
        super(props);
        this.state = {loading: false, err: null, url: ""}
        this.handleUrlChange = this.handleUrlChange.bind(this)
        this.submitUrl = this.submitUrl.bind(this)
    }
    handleUrlChange(input) {
        this.setState({url: input.target.value})
    }
    submitUrl(input) {
        this.setState({loading: true})
        console.log("/add?url="+encodeURIComponent(this.state.url))
        fetch("/add?url="+encodeURIComponent(this.state.url))
            .then(
                (result) => {
                    console.log(result)
                    this.setState({loading: false})
                },
                (error) => {
                    console.log(error)
                    this.setState({err: error, loading: false})
                }
            )
    }
    render() {
        return (
            <div>
                <Header>
                    Enter a url to start scraping.
                </Header>
                {/*<Input fluid icon='search' placeholder='Enter a url...' loading={this.state.loading} onClick={this.submitUrl}/>*/}
                <Form onSubmit={this.submitUrl} loading={this.state.loading}>
                    <Form.Input type="url" name="url" onChange={this.handleUrlChange}/>
                    <Form.Button>Submit</Form.Button>
                </Form>
                {/*{this.state.err? <Card>{this.state.err}</Card> : null}*/}
            </div>
        )
    }
}
