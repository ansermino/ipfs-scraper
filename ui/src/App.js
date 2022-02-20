import React, { Component } from 'react'
import './App.css';
import './Sidebar'
import Sidebar from "./Sidebar";
import {Header, Segment, Grid, Image, Container} from "semantic-ui-react";
import {Route, BrowserRouter as Router, Routes} from "react-router-dom";
import Add from "./pages/Add";
import View from "./pages/View";
import Home from "./pages/Home";

class App extends React.Component{
  render () {
      return (
          <div>
              <Segment clearing style={{backgroundColor: "#bde1f1"}}>
                  <Header as='h2' floated='left' >
                      <Image src='./logo.png'/> IPFS That
                  </Header>
              </Segment>
              <Router>
                  <Grid streched>
                  <Sidebar />
                  <Grid.Column streched width={12} style={{height: "100%"}}>
                      <Segment>
                      <Routes>
                          <Route
                              path="/"
                              exact
                              element={<Home />}
                          />
                          <Route
                              path="/add"
                              exact
                              element={<Add />}
                          />
                          <Route
                              path="/view/*"
                              exact
                              element={<View />}
                          />
                      </Routes>
                      </Segment>
                  </Grid.Column>
                  </Grid>
              </Router>
              <Container textAlign='center'><br/>Made @ EthDenver 2022</Container>
          </div>
      );
  }
}

export default App;
