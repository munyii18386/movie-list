import React from "react";
import {Col, Form, Button} from "react-bootstrap";
import "./MyWall.css";

export default function MyWall(props){
    return(
        <Form id="search-bar">
            <Form.Row>
                <Col>
                    <Form.Control placeholder="Titles, People, Genres.." />
                </Col>
                <Button variant="primary" type="submit">Search</Button>
            </Form.Row>
</Form>
    )
}