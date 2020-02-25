import React, { Component } from "react";
import { Redirect } from "react-router-dom" 
import axios from "axios";
import {Container, Form, Button, Row, Col} from "react-bootstrap";
import {Home} from './Home';
import {Login} from './Login';
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
   
  
    handleSubmit = event => {
            // axios({
            //     method: 'post',
            //     url: 'https://oddgarden.net/api/SignUp',
            //     data: this.state
            //   })
            //  .then((response) =>{
            //      console.log(response.data)
            //     this.state.NewUser = true
            //     console.log(this.state)
            // })
            // .catch(function (error) {
            //     console.log(error);
            // })
            event.preventDefault()
            axios.post('https://localhost/api/SignUp', this.state)
              .then((response) => {
                console.log(response.data.Status);
                this.setState({NewUser: response.data.Status})
                console.log(this.state);
              })
              .catch(function (error) {
                console.log(error);
              });
            
    }
  

    render(){
        if (this.state.NewUser === "true") {
            return <Redirect to='/login' exact component={Login}/>
        }
        return(
            <Container className="SignUP">
            <Form onSubmit={this.handleSubmit}>
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
                <Button block variant="primary" type="submit">Sign Up</Button>
                {/*  onSubmit={(e)=>{this.handleSubmit(e)}}*/}
            </Form>
            </Container>
        )
    }
}