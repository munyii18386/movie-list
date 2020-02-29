import React, { Component } from "react";
import {Col, Form, Button} from "react-bootstrap";
import "./MyWall.css";

export class  MyWall extends Component{
    constructor(){
        super()
        this.state ={
            SearchFor: ""
        }
    }

    handleChange(e){
        this.state.SearchFor = e.target.value
    }


    render(){
        return(
            <Form id="search-bar" onSubmit={this.handleSubmit}>
                <Form.Row>
                    <Col>
                        <Form.Control onChange={(e)=>{this.handleChange(e)}} placeholder="Search Titles, People, Genres.." />
                    </Col>
                    <Button variant="primary" type="submit">Search</Button>
                </Form.Row>
            </Form>
        )
    }
}