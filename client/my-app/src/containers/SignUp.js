import React, { Component } from "react";
import axios from "axios";
import {Container, Form, Button, Row, Col} from "react-bootstrap";
import "./SignUp.css";



export class SignUp extends Component{
    constructor(){
        super()
        this.handleChange =  this.handleChange.bind(this)
        this.handleSubmit = this.handleSubmit.bind(this)
        this.state={
            FirstName:"",
            LastName:"",
            Email:"",
            Password:"",
            PasswordConf:"",
            NewUser: false,
        }
    }

    handleChange(e){
        e.preventDefault();
        switch(e.target.id){
            case "firstname": this.state.FirstName = e.target.value;
            case "lastname": this.state.LastName = e.target.value;
            case "email": this.state.Email = e.target.value;
            case "pass1": this.state.Password = e.target.value;
            case "pass2": this.state.PasswordConf = e.target.value

        }
        console.log(this.state)
    }

    handleSubmit(e){
        e.preventDefault();
        axios.post('http://localhost/SignUp', this.state)
            .then((r) =>{
                console.log(r)
                this.state.NewUser = true
            })
            .catch(function (error) {
                console.log(error);
            });
        
    }

    render(){
        return(
            <Container className="SignUP">
            <Form>
                <Form.Row>
                    <Form.Group as={Col} controlId="formGridEmail">
                        <Form.Control id="firstname" onChange={(e)=>{this.handleChange(e)}} size="lg" placeholder="First Name" />
                    </Form.Group>

                    <Form.Group as={Col} controlId="formGridEmail">
                        <Form.Control id="lastname" onChange={(e)=>{this.handleChange(e)}} size="lg" placeholder="Last Name" />
                    </Form.Group>
                </Form.Row>

                <Form.Group controlId="formGridAddress1">
                    <Form.Control id="email" onChange={(e)=>{this.handleChange(e)}} placeholder="Enter Email Address" />
                </Form.Group>     

                <Form.Row>
                    <Form.Group as={Col} controlId="formGridPassword">
                        <Form.Control id="pass1" onChange={(e)=>{this.handleChange(e)}} size="lg" type="password" placeholder="Password" />
                    </Form.Group>
                    <Form.Group as={Col} controlId="formGridPassword">
                        <Form.Control id="pass2" onChange={(e)=>{this.handleChange(e)}} size="lg" type="password" placeholder="Re-enter Password" />
                    </Form.Group>
                </Form.Row>
                <Button  onSubmit={(e)=>{this.handleSubmit(e)}} block variant="primary" type="submit">Sign Up</Button>
                {/*  onSubmit={(e)=>{this.handleSubmit(e)}}*/}
            </Form>
            </Container>
        )
    }
}