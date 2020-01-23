import React from 'react';
import { Link } from "react-router-dom";
import { Container, Navbar, Nav, NavItem } from "react-bootstrap";
import Routes from "./Routes"


import './App.css';

function App(props) {
  
    return (
      <Container className="App container">
        <Navbar bg="light" expand="lg" fluid collapseOnSelect>
          <Navbar.Brand>
            <Link to="/">Wall Of Movies</Link>
          </Navbar.Brand>
          <Navbar.Toggle/>
          <Navbar.Collapse className="justify-content-end">
            <Nav>
              <Nav.Link href="/Login">Login</Nav.Link>
              <Nav.Link href="/SignUp">SignUp</Nav.Link>
            </Nav>
          </Navbar.Collapse>
        </Navbar>
        <Routes />
    </Container>
    )
}

export default App;
