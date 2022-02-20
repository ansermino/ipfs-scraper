import React, {Component, useEffect, useState} from 'react'
import {List, Message} from 'semantic-ui-react'
import {NavLink, useLocation} from "react-router-dom";


const ViewAll = (props) => {
    let { pathname } = useLocation();
    const [pages, setPages] = useState([]);
    useEffect(() => {
        fetch("/pages")
            .then(res => res.json())
            .then(
                (result) => {
                    console.log(result)
                    setPages(result.pages)
                },
                (error) => {
                    console.log(error)
                }
            )
    }, [])

    return (
        <List divided relaxed>
            {pages.map((item, i) => {
                return (
                    <List.Item key={i} as={NavLink} to={`${pathname}/${item.title}`} >
                        <List.Icon name='folder' size='large' verticalAlign='middle' />
                        <List.Content>
                            <List.Header as='a'>{item.title}</List.Header>
                            <List.Description as='a'>{item.url}</List.Description>
                        </List.Content>
                    </List.Item>
                )
            })}
        </List>
    )
}

export default ViewAll
