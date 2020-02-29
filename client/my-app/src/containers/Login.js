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

    handleEmail(e){
        this.setState({email: e.target.value})
    }

    handlePassword(e){
        this.setState({password: e.target.value})
    }

    handleInput(e){
        console.log(e.target.id)
        console.log(this.state.email)
        console.log(this.state.password)
    }

    render(){
        return(
            
            <div className="Login">
                <Form onSubmit={(event) =>{this.handleSubmit(event)}}>
                    <Form.Group controlId="formBasicEmail">
                        <Form.Label>Email address</Form.Label>
                        <Form.Control onChange={(e)=>{this.handleEmail(e)}} size="lg"  type="email" placeholder="Enter email" />
                    </Form.Group>
                    <Form.Group controlId="formBasicPassword">
                        <Form.Label>Password</Form.Label>
                        <Form.Control id="pass"   onChange={(e)=>{this.handlePassword(e)}}size="lg" type="password" placeholder="Password" />
                    </Form.Group>
                    <Button id="login-button"  onClick={(e)=>{this.handleInput(e)}}variant="primary" block type="submit">Login</Button>
                </Form>
            </div>
        )
    }
}
