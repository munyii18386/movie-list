import React, { Component } from "react";
import { Button, Form } from "react-bootstrap";
import "./Login.css";

export class  Login extends Component {

    constructor(){
        super()
        this.state = {
            email: "",
            password: ""
        }
    }


    handleSubmit(event) {
        event.preventDefault();
        
    }

    handleInput(e){
        console.log(this.state.email)
    }

    render(){
        return(
            
            <div className="Login">
                <Form onSubmit={(event) =>{this.handleSubmit(event)}}>
                    <Form.Group controlId="formBasicEmail">
                        <Form.Label>Email address</Form.Label>
                        <Form.Control onChange={(e)=>{this.setState({email: e.target.value})}} size="lg"  type="email" placeholder="Enter email" />
                    </Form.Group>
                    <Form.Group controlId="formBasicPassword">
                        <Form.Label>Password</Form.Label>
                        <Form.Control size="lg" type="password" placeholder="Password" />
                    </Form.Group>
                    <Button onClick={(e)=>{this.handleInput(e)}}variant="primary" block type="submit">Login</Button>
                </Form>
            </div>
        )
    }
}
