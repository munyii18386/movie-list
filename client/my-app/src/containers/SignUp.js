import React, { Component } from "react";
import {Container, Form, Button, Row, Col} from "react-bootstrap";
import "./SignUp.css";

export class SignUp extends Component{
    constructor(){
        super()
        this.state={
            fname:"",
            lname:"",
            email:"",
            pass1:"",
            pass2:"",
        }
    }
    render(){
        return(
            <Container className="SignUP">
            <Form>
                <Form.Row>
                    <Form.Group as={Col} controlId="formGridEmail">
                        <Form.Control  size="lg" placeholder="First Name" />
                    </Form.Group>

                    <Form.Group as={Col} controlId="formGridEmail">
                        <Form.Control size="lg" placeholder="Last Name" />
                    </Form.Group>
                </Form.Row>

                <Form.Group controlId="formGridAddress1">
                    <Form.Control placeholder="Enter Email Address" />
                </Form.Group>     

                <Form.Row>
                    <Form.Group as={Col} controlId="formGridPassword">
                        <Form.Control size="lg" type="password" placeholder="Password" />
                    </Form.Group>
                    <Form.Group as={Col} controlId="formGridPassword">
                        <Form.Control size="lg" type="password" placeholder="Re-enter Password" />
                    </Form.Group>
                </Form.Row>
                <Button block variant="primary" type="submit">Sign Up</Button>
            </Form>
            </Container>
        )
    }
}