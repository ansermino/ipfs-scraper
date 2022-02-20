import React, {useEffect, useState} from 'react'
import {Header, List, Message} from 'semantic-ui-react'
import {Link} from "react-router-dom";

const Home = () => {
    return (
        <div>
            <Header as='h1'>Welcome to IPFS That!</Header>
            <p><b>IPFS That</b> is a utility that bring the web to IPFS! It allows users to specify a web page and make it available via IPFS.</p>

            <Header as='h2'>What Is IPFS?</Header>
            <p>IPFS is a decentralized peer-to-peer file storage protocol. It allows anyone to store files that can be retrieved by anyone else on the network.</p>

            <Header as='h2'>Why Would I Want a Webpage on IPFS?</Header>
            <p>Because IPFS is a peer-to-peer technology it is censorship resistant. As a user when you request a file there are multiple peers that can serve it to you. So if one or some of the peers choose to censor the content it is still available from the remaining peers.</p>

            <Header as='h2'>How Can I Learn More?</Header>
            <List bulleted>
                <List.Item>
                    <Link to={"https://docs.ipfs.io/"}>IPFS Docs</Link>
                </List.Item>
                <List.Item>
                    <Link to={"https://github.com/ipfs"}>IPFS Github</Link>
                </List.Item>
                <List.Item>
                    <Link to={"https://github.com/ansermino/ipfs-that"}>IPFS That Github</Link>
                </List.Item>
            </List>
        </div>
    )
}

export default Home
