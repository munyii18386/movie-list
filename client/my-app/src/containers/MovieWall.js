import React  from "react";
import {Container, Nav, Row, Col, Form, Button, Card, CardGroup} from "react-bootstrap";
import axios from "axios";
import "./MovieWall.css";
import { useSelector } from "react-redux";
import { useDispatch } from "react-redux";

export const MovieWall = () =>{
    const state  = useSelector(state => state);
    const dispatch = useDispatch();

    const initialState = {
       update: false
    };
    const [data, setData] = React.useState(initialState);

    const load = () => {
        let config = {
            headers: {
                Authorization: state.token
            }
        }
        axios.get('https://localhost/api/GetWall', config)
        .then((response)=>{
            console.log(response)
        }) 
        .catch((e) =>{
            console.log(e)
        })
    }

    return(
       <Container>
            <div className="search-margin"></div>
                <Nav fill variant="tabs" defaultActiveKey="/moviewall">
                    <Nav.Item>
                        <Nav.Link href="/search">Search</Nav.Link>
                    </Nav.Item>
                    <Nav.Item>
                        <Nav.Link href="/moviewall">My Wall</Nav.Link>
                    </Nav.Item>
                </Nav>
                <div className="search-margin"></div>
                {data.update ? load : <div></div>}
                
       </Container>
    )
}