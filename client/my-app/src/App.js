import React from 'react';
import { BrowserRouter as Router, Link } from "react-router-dom";
import { Container, Navbar, Nav } from "react-bootstrap";
import {Routes} from "./Routes"
import { useSelector } from "react-redux";
import './App.css';






function App(){
  // const msg = "Wall Of Movies"
  // const [state, dispatch] = React.useReducer(reducer, initialState);
  const state  = useSelector(state => state);
  // const dispatch = useDispatch();
  console.log ("Auth: " + state.isAuthenticated, " disabled: " + state.disable + " user: " + state.user + " token: " + state.token)
  
      return (
          <Router>
            <Container className="App container">
              <Navbar bg="light" expand="lg" fluid collapseOnSelect>
                <Navbar.Brand>
                  <Link to="/"> {state.isAuthenticated ? "Welcome " + state.user : "Wall Of Movies"} </Link>
                </Navbar.Brand>
                <Navbar.Toggle/>
                <Navbar.Collapse className="justify-content-end">
                  <Nav>
                    <Nav.Link href="/SignUp">SignUp</Nav.Link>
                    <Nav.Link href="/Login">Login</Nav.Link>
                    <Nav.Link href="/Logout">Logout</Nav.Link>
                  </Nav>
                </Navbar.Collapse>
              </Navbar>
              <Routes />
          </Container>
        </Router>
      )

}



export default App
