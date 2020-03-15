import React from "react";
import {Container, Form, Button} from "react-bootstrap";
import axios from "axios";
import { Redirect } from 'react-router-dom'
import {MyWall} from './MyWall';
import "./Login.css";
import { useSelector } from "react-redux";
import { useDispatch } from "react-redux";

export const Login = () => {
    const state  = useSelector(state => state);
    const dispatch = useDispatch();
    const initialState = {
        Email: "",
        Password: ""
    };
    const [data, setData] = React.useState(initialState);
    const handleInputChange = event => {
        // console.log("current user is authenticated: " + state.isAuthenticated);
        let re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
        if(event.target.name == "email"){
            let email = event.target.value;
            console.log("this is the email input")
            if(re.test(email)){
                console.log("valid email")
            }
        }
        setData({
          ...data,
          [event.target.name]: event.target.value
        });
        
    };
    console.log(data)

    const handleSubmit = event => {
        event.preventDefault()
        setData({
            ...data
        });

        axios.post('https://localhost/api/Login', data)
        .then((response)=>{
            console.log(response)
            let user = response.data.firstName
            let token = response.headers.authorization
            console.log("user: " + user + "auth: " + token);
            dispatch({
                type: "LOGIN",
                payload: {user, token}  
            })
        }) 
        .catch((e) =>{
            console.log(e)
        })   

    }

   

        return(
            <Container className="Login">
            {state.isAuthenticated ?  <Redirect to='/mywall' /> :
                <Form onSubmit={handleSubmit}>
                    <Form.Group controlId="formBasicEmail">
                        <Form.Label>Email address</Form.Label>
                        <Form.Control 
                            value={data.email}
                            onChange={handleInputChange} 
                            name="email"
                            size="lg"  
                            type="email" 
                            placeholder="Enter email" 
                        />
                    </Form.Group>
                    <Form.Group controlId="formBasicPassword">
                        <Form.Label>Password</Form.Label>
                        <Form.Control
                             value={data.password}
                             onChange={handleInputChange}
                             name="password"
                             id="password"  
                             size="lg" 
                             type="password" 
                             placeholder="Password" 
                        />
                    </Form.Group>
                    <Button id="login-button" variant="primary" block type="submit">Login</Button>
                </Form>
            }
            </Container>
        )

}
