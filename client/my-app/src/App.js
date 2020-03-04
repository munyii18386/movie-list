import React from 'react';
import { Link } from "react-router-dom";
import { Container, Navbar, Nav, NavItem } from "react-bootstrap";
import {Routes} from "./Routes"
import './App.css';

export const AppContext = React.createContext();
const initialState = {
  isAuthenticated: false,
  disable: false,
  user: null,
  token: null,
};

const reducer = (state, action)=> {
  switch (action.type){
    case "LOGIN":
      localStorage.setItem("user", JSON.stringify(action.payload.user));
      localStorage.setItem("token", JSON.stringify(action.payload.token));
      return{
        ...state,
        isAuthenticated: true,
        user: action.payload.user,
        token: action.payload.token
      };
    case "SIGNUP":
      localStorage.setItem("user", JSON.stringify(action.payload.user));
      localStorage.setItem("token", JSON.stringify(action.payload.token));
      return{
        ...state,
        isAuthenticated: true,
        disable: true,
        user: action.payload.user,
        token: action.payload.token
      };
      case "LOGOUT":
        localStorage.clear();
        return {
          ...state,
          isAuthenticated: false,
          user: null
        };
      default:
        return state;
  }
}



function App(){
  const msg = "Wall Of Movies"
  const [state, dispatch] = React.useReducer(reducer, initialState);
  console.log ("user is auth: " + state.isAuthenticated)
  
      return (
        <AppContext.Provider
          value={{
            state,
            dispatch
          }}
        >
          <Container className="App container">
            <Navbar bg="light" expand="lg" fluid collapseOnSelect>
              <Navbar.Brand>
                <Link to="/">{!state.isAuthenticated ? "Wall Of Movies" : "Welcome, " + state.user}</Link>
              </Navbar.Brand>
              <Navbar.Toggle/>
              <Navbar.Collapse className="justify-content-end">
                <Nav>
                  <Nav.Link href="/SignUp" disabled={state.disable}>SignUp</Nav.Link>
                  <Nav.Link href="/Login" disabled={state.disable}>Login</Nav.Link>
                  <Nav.Link href="/Logout" disabled={!state.disable}>Logout</Nav.Link>
                </Nav>
              </Navbar.Collapse>
            </Navbar>
            <Routes />
        </Container>
      </AppContext.Provider>
      )

}

export default App;
