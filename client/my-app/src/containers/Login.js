import React from "react";
import { Button, Form } from "react-bootstrap";
import "./Login.css";

export default function Login(props) {

    function handleSubmit(event) {
        event.preventDefault();
    }

    return(
        
        <div className="Login">
            <Form onSubmit={handleSubmit}>
                <Form.Group controlId="formBasicEmail">
                    <Form.Label>Email address</Form.Label>
                    <Form.Control size="lg"  type="email" placeholder="Enter email" />
                </Form.Group>
                <Form.Group controlId="formBasicPassword">
                    <Form.Label>Password</Form.Label>
                    <Form.Control size="lg" type="password" placeholder="Password" />
                </Form.Group>
                <Button variant="primary" block type="submit">Login</Button>
            </Form>
        </div>
    )
}
