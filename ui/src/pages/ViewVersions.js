import React, {Component, useEffect, useState} from 'react'
import {Header, List, Message} from 'semantic-ui-react'
import {NavLink, useLocation, useParams} from "react-router-dom";

const ViewVersions = () => {
    let { pathname } = useLocation();
    let { title } = useParams();

    const [versions, setVersions] = useState([]);
    useEffect(() => {
        console.log("/api/versions?title="+title)
        fetch("/api/versions?title="+title)
            .then(res => res.json())
            .then(
                (result) => {
                    console.log(result)
                    setVersions(result.versions)
                },
                (error) => {
                    console.log(error)
                }
            )
    }, [title])

    return (
        <div>
            <Header>{title}</Header>
            <List divided relaxed>
                {versions.map((version, i) => {
                    return (
                        <List.Item key={i} href={`${window.location.protocol}//${window.location.host}/ipfs/${version.cid}`} target="_blank">
                            <List.Icon name='clock' size='large' verticalAlign='middle' />
                            <List.Content>
                                <List.Header>{version.timestamp}</List.Header>
                                <List.Description as='p'>{version.cid}</List.Description>
                            </List.Content>
                        </List.Item>
                    )
                })}
            </List>
        </div>

    )
}

export default ViewVersions
