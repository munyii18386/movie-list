import React, {Component}  from "react";
import axios from "axios";
import {Container, Form, Button, Row, Col} from "react-bootstrap";
import {MyWall} from './MyWall';
import "./SignUp.css";
import { useSelector } from "react-redux";
import { useDispatch } from "react-redux";



export const SignUp = () => {
    // const {state , dispatch } = React.useContext(AppContext);
    const state  = useSelector(state => state);
    const dispatch = useDispatch();
    const initialState={
            FirstName:"",
            LastName:"",
            Email:"",
            Password:"",
            PasswordConf:""

        }
    const [data, setData] = React.useState(initialState);

    const handleInputChange = event => {
        // console.log("current user is authenticated: " + state.isAuthenticated);
        setData({
          ...data,
          [event.target.name]: event.target.value
        });
    };

      

  
        const handleSubmit = event => {
            event.preventDefault()
            axios.post('https://localhost/api/SignUp', data)
              .then((response) => {
                console.log(response);
                console.log("current user is authenticated: " + state.isAuthenticated);
                let user = response.data.firstName
                let token = response.headers.authorization
                let ok = true
                console.log("user: " + user + "auth: " + token);
                dispatch({
                    type: "SIGNUP",
                    payload: {user, token}  
                })
              })
              .catch(function (error) {
                console.log(error);
              });
            
        }
    
    
        
    
    
        return(
            
            <Container className="SignUP">
            {state.isAuthenticated ? <MyWall /> :
            <Form onSubmit={handleSubmit}>
                <Form.Row>
                    <Form.Group as={Col} controlId="formGridEmail">
                        <Form.Control 
                             value={data.firstname}
                             onChange={handleInputChange}
                             name="firstname"
                             id="firstname"
                             size="lg" placeholder="First Name" />
                    </Form.Group>

                    <Form.Group as={Col} controlId="formGridEmail">
                        <Form.Control  
                             value={data.lastname}
                             onChange={handleInputChange}
                             name="lastname"
                             id="lastname"
                             size="lg" placeholder="Last Name" />
                    </Form.Group>
                </Form.Row>

                <Form.Group controlId="formGridAddress1">
                    <Form.Control  value={data.email}
                        onChange={handleInputChange}
                        name="email"
                        id="email" 
                        placeholder="Enter Email Address" />
                </Form.Group>     

                <Form.Row>
                    <Form.Group as={Col} controlId="formGridPassword">
                        <Form.Control 
                            value={data.password}
                            onChange={handleInputChange}
                            name="password"
                            id="password"
                            size="lg" type="password" placeholder="Password" />
                    </Form.Group>
                    <Form.Group as={Col} controlId="formGridPassword">
                        <Form.Control 
                            value={data.PasswordConf}
                            onChange={handleInputChange}
                            name="PasswordConf"
                            id="PasswordConf"
                            size="lg" type="password" placeholder="Re-enter Password" />
                    </Form.Group>
                </Form.Row>
                <Button  block variant="primary" type="submit">Sign Up</Button>
            </Form>}
            
           </Container>
        )
    
};

export default SignUp;