import React from 'react';
import {Container, Form, Button, Row, Col} from "react-bootstrap";
import './Home.css';

export const Home = () => {
   
        return(
            <Container className="Home">
                <div className="landing">
                    <h1>Wall Of Movies</h1>
                    <p>A simple movie list app</p>
                </div>
            </Container>
        )
}